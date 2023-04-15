package goargs

import (
	"fmt"
	"os"
)

const (
	Zero Narg = iota
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
	args map[string]Arg
}

func (p *Parser) Contains(key string) bool {
	_, ok := p.args[key]
	return ok
}

func New() *Parser {
	return &Parser{
		args: make(map[string]Arg),
	}
}

func (p *Parser) addArg(arg Arg) {
	p.args[arg.name] = arg
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

func (p *Parser) parse(args []string) (Vals, error) {
	vals := make(Vals)

	for i, arg := range args {
		if p.Contains(arg) {
			vs := []string{}
			for _, val := range args[i+1:] {
				if p.Contains(val) {
					break
				}
				vs = append(vs, val)
			}
			vals[arg] = vs
		}
	}

	// check for no missing required args
	for key := range p.args {
		if !vals.Contains(key) && p.args[key].required {
			return nil, fmt.Errorf("missing required arg %v", key)
		}
	}

	for key, val := range vals {
		switch p.args[key].nvar {
		case Zero:
			if len(val) > 0 {
				return nil, fmt.Errorf("%v has values when it shouldn't", key)
			}
		case One:
			if len(val) != 1 {
				return nil, fmt.Errorf("%v does not have exactly one value", key)
			}
		}
	}

	return vals, nil
}

func (p *Parser) Parse() (Vals, error) {
	return p.parse(os.Args[1:])
}
