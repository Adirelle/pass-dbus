package pass

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

type (
	Pass struct {
		PassPath    string
		Collections map[string]*Collection
	}
)

func NewPass(path string) *Pass {
	p := &Pass{
		PassPath:    path,
		Collections: make(map[string]*Collection),
	}
	return p
}

func (p *Pass) Add(name, path string) *Collection {
	c := &Collection{
		Name:  name,
		Path:  path,
		Items: make(map[ItemPath]*Item),
		pass:  p,
	}
	p.Collections[name] = c
	return c
}

func (p *Pass) Get(name string) (*Collection, error) {
	if c, found := p.Collections[name]; found {
		return c, nil
	}
	return nil, fmt.Errorf("unknown collection: %s", name)
}

func (p *Pass) exec(storePath string, args []string) (io.Reader, error) {

	buffer := &bytes.Buffer{}

	cmd := exec.Command(p.PassPath, args...)
	cmd.Env = append(
		cmd.Env,
		fmt.Sprintf("PASSWORD_STORE_DIR=%s", storePath),
		"LANG=C",
	)
	cmd.Stdin = nil
	cmd.Stdout = buffer

	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return buffer, nil
}
