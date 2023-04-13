package goargs

import (
	"testing"
)

func TestNoArgs(t *testing.T) {
	p := New()
	p.Optional("--not-used", Zero)

	val, err := p.parse([]string{"path/to/binary"})
	if err != nil {
		t.Error(err)
	}

	if len(val) > 0 {
		t.Error("no args expected")
	}
}

func TestFlagArgs(t *testing.T) {
	p := New()
	p.Required("-v", Zero)

	val, err := p.parse([]string{"-v"})
	if err != nil {
		t.Error(err)
	}

	if _, ok := val["-v"]; !ok {
		t.Error("flag not in result")
	}
}

func TestVariaticArgs(t *testing.T) {
	t.Error("Not implemented")
}

func TestMixedArgs(t *testing.T) {
	t.Error("Not implemented")
}
