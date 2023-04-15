package goargs

import (
	"fmt"
	"strings"
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
	p := New()
	p.Required("--myList", Many)
	p.Required("-f", Zero)
	p.Required("--flag", Zero)

	vals, err := p.parse(strings.Split("-f --myList one two three --flag", " "))
	if err != nil {
		t.Error(err)
	}

	if len(vals["--myList"]) != 3 {
		t.Error(fmt.Errorf("expected three args"))
	}
}

func TestMixedArgs(t *testing.T) {
	p := New()
	p.Required("--files", Many)
	p.Required("--name", One)

	vals, err := p.parse(strings.Split("--files one.txt two.txt index.html --name Connor", " "))
	if err != nil {
		t.Error(err)
	}

	if len(vals["--files"]) != 3 {
		t.Error(fmt.Errorf("incorrect number of 'files' args"))
	}
	if len(vals["--name"]) != 1 {
		t.Error(fmt.Errorf("name expecting exactly one value"))
	}
}
