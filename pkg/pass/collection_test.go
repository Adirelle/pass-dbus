package pass_test

import "testing"

func TestCollection(t *testing.T) {

	p := newTestPass(t)

	c, err := p.Get("default")
	if err != nil {
		t.Fatalf(err.Error())
	}

	list, err := c.List()
	if err != nil {
		t.Fatalf(err.Error())
	}

	if list == nil {
		t.FailNow()
	}
}
