package goargs

import (
	"fmt"
	"os"
)

const (
	Zero int = iota
	One
	Many
)

type Narg int

type Arg struct {
	name     string
	nvar     Narg
	required bool
}

type Parser struct {
	args []Arg
}

func (p *Parser) addArg(arg Arg) {
	p.args = append(p.args, arg)
}

func (p *Parser) Required(name string, nvar Narg) {
	p.addArg(Arg{
		required: true,
		name:     name,
		nvar:     nvar,
	})
}

func (p *Parser) Optional(name string, nvar Narg) {
	p.addArg(Arg{
		required: false,
		name:     name,
		nvar:     nvar,
	})
}

func (p *Parser) Parse() (Vals, error) {
	cmd := os.Args
	fmt.Println(cmd)

	return nil, fmt.Errorf("not implemented")
}
