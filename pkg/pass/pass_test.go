package pass_test

import (
	"os"
	"testing"

	"github.com/adirelle/pass-dbus/pkg/pass"
)

func newTestPass(t *testing.T) *pass.Pass {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf(err.Error())
	}
	p := pass.NewPass("/usr/bin/pass")
	p.Add("default", wd+"/testdata/store")
	return p
}

func TestPass(t *testing.T) {
	p := newTestPass(t)
	if p == nil {
		t.Fail()
	}
}
