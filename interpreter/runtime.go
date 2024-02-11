package interpreter

import (
	"fmt"

	"github.com/shreyassanthu77/cisp/ast"
)

func Eval(program ast.Program) (ast.Value, error) {
	rootEnv := NewRootEnv()
	for _, rule := range program.Rules {
		switch rule := rule.(type) {
		case ast.AtRule:
			panic("global at-rules not supported yet")
		case ast.Rule:
			rootEnv.setFn(rule)
		}
	}

	main, err := rootEnv.genFn("main")
	if err != nil {
		return ast.NilValue{}, fmt.Errorf("no main rule found")
	}

	return evalRule(main, []ast.Value{}, rootEnv)
}
