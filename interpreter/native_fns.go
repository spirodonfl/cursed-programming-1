package interpreter

import (
	"fmt"

	"github.com/shreyassanthu77/cisp/ast"
	"github.com/shreyassanthu77/cisp/lexer"
)

type NativeFnCall struct {
	Fn         ast.Identifier
	Parameters []ast.Attreibute
	Handler    func(env *Environment) (ast.Value, error)
}

func (n NativeFnCall) IsStatement() {}

func (n NativeFnCall) GetSpan() lexer.Span {
	return lexer.Span{}
}

func ruleFromNativeFnCall(n NativeFnCall) ast.Rule {
	return ast.Rule{
		Selector: ast.Selector{
			Identifier: n.Fn,
			Atrributes: n.Parameters,
		},
		Body: []ast.Statement{n},
	}
}

var printFn = ruleFromNativeFnCall(NativeFnCall{
	Fn: ast.Identifier{Name: "print"},
	Parameters: []ast.Attreibute{
		{
			Name:    ast.Identifier{Name: "value"},
			Default: ast.NilValue{},
		},
	},
	Handler: func(env *Environment) (ast.Value, error) {
		val, _ := env.getVar("value")
		switch val := val.(type) {
		case ast.String:
			fmt.Println(val.Value)
		case ast.Int:
			fmt.Println(val.Value)
		case ast.Float:
			fmt.Println(val.Value)
		case ast.Boolean:
			fmt.Println(val.Value)
		case ast.NilValue:
			fmt.Println("nil")
		default:
			fmt.Println(val)
		}
		return ast.NilValue{}, nil
	},
})
