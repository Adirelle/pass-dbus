package pass

import (
	"bufio"
	"io"
	"strings"
)

type (
	Collection struct {
		Name  string
		Path  string
		Items map[ItemPath]*Item

		pass *Pass
	}
)

func (c *Collection) exec(args ...string) (io.Reader, error) {
	return c.pass.exec(c.Path, args)
}

func (c *Collection) List() ([]string, error) {
	output, err := c.exec("ls")
	if err != nil {
		return nil, err
	}

	return c.parseList(output)
}

const (
	newChild  = "|-- "
	lastChild = "`-- "
	dive      = "|   "
	diveLast  = "    "
)

func (c *Collection) parseList(output io.Reader) (items []string, err error) {
	scanner := bufio.NewScanner(output)
	scanner.Split(bufio.ScanLines)

	// Skip first line (header)
	_ = scanner.Scan()

	path := make([]string, 100)
	leaves := make(map[string]bool, 100)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 4 {
			continue
		}
		level := 0
	levels:
		for i := 4; i <= len(line); i += 4 {
			token := line[i-4 : i]
			switch token {
			case newChild, lastChild:
				j := strings.IndexByte(line[i:], ' ')
				if j == -1 {
					path[level] = line[i:]
				} else {
					path[level] = line[i : i+j]
				}
				if level > 0 {
					parent := strings.Join(path[:level], "/")
					leaves[parent] = false
				}
				level += 1
				leaf := strings.Join(path[:level], "/")
				leaves[leaf] = true
				items = append(items, leaf)
				break levels
			case dive, diveLast:
				level += 1
			}
		}
	}

	j := 0
	for i := 0; i < len(items); i++ {
		if leaves[items[i]] {
			items[j] = items[i]
			j++
		}
	}
	items = items[:j]

	err = scanner.Err()
	return
}
