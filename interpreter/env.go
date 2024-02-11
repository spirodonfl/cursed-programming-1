package interpreter

import (
	"fmt"

	"github.com/shreyassanthu77/cisp/ast"
)

type Environment struct {
	Parent *Environment
	Funcs  map[string]ast.Rule
	Vars   map[string]ast.Value
}

func NewRootEnv() *Environment {
	return &Environment{
		Parent: nil,
		Funcs: map[string]ast.Rule{
			"print": printFn,
		},
	}
}

func (e *Environment) fork() *Environment {
	return &Environment{
		Parent: e,
		Funcs:  make(map[string]ast.Rule),
		Vars:   make(map[string]ast.Value),
	}
}

func (e *Environment) genFn(name string) (ast.Rule, error) {
	fn, ok := e.Funcs[name]
	if !ok && e.Parent != nil {
		return e.Parent.genFn(name)
	}

	if !ok {
		return ast.Rule{}, fmt.Errorf("function %s not found", name)
	}

	return fn, nil
}

func (e *Environment) setFn(fn ast.Rule) error {
	name := fn.Selector.Identifier.Name
	_, ok := e.Funcs[name]
	if ok {
		return fmt.Errorf("function %s already defined in this scope", name)
	}

	e.Funcs[name] = fn
	return nil
}

func (e *Environment) setVar(name string, value ast.Value) {
	e.Vars[name] = value
}

func (e *Environment) getVar(name string) (ast.Value, error) {
	val, ok := e.Vars[name]
	if ok {
		return val, nil
	}

	if e.Parent != nil {
		return e.Parent.getVar(name)
	}

	return ast.NilValue{}, fmt.Errorf("variable %s not found", name)
}
