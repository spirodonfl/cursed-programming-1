package interpreter

import (
	"fmt"

	"github.com/shreyassanthu77/cisp/ast"
)

func isNilValue(v ast.Value) bool {
	_, ok := v.(ast.NilValue)
	return ok
}

func evalVarDeclaration(decl ast.Declaration, env *Environment) (ast.Value, error) {
	name := decl.Property.Name[2:] // remove the -- from the name

	if len(decl.Parameters) != 1 {
		return ast.NilValue{}, fmt.Errorf("variable declaration should have exactly one value")
	}
	value, err := evalValue(decl.Parameters[0], env)

	val, err := evalValue(value, env)
	if err != nil {
		return ast.NilValue{}, err
	}
	env.setVar(name, val)
	return val, nil
}

func evalStmt(stmt ast.Statement, ifState *IfState, env *Environment) (ast.Value, error) {
	switch stmt := stmt.(type) {
	case ast.Rule:
		err := env.setFn(stmt)
		if err != nil {
			return ast.NilValue{}, err
		}
		ifState.reset()
		return ast.NilValue{}, nil
	case ast.AtRule:
		return evalAtRule(env, ifState, stmt)
	case ast.Declaration:
		if len(stmt.Property.Name) > 2 && stmt.Property.Name[:2] == "--" {
			return evalVarDeclaration(stmt, env)
		}
		fnCall := ast.FunctionCall{
			Fn:         stmt.Property,
			Parameters: stmt.Parameters,
		}
		ifState.reset()
		return evalFnCall(fnCall, env)
	case NativeFnCall:
		ifState.reset()
		return stmt.Handler(env)
	}

	return ast.NilValue{}, nil
}

func verifyAndAddParamsToEnv(attributes []ast.Attreibute, params []ast.Value, env *Environment) error {
	if len(attributes) != len(params) {
		return fmt.Errorf("expected %d parameters, got %d", len(attributes), len(params))
	}

	for i, attr := range attributes {
		param := params[i]
		if isNilValue(param) {
			if attr.Default != nil {
				param = attr.Default
			} else {
				return fmt.Errorf("parameter %s is required", attr.Name.Name)
			}
		}
		env.setVar(attr.Name.Name, param)
	}

	return nil
}

func evalStatementList(stmts []ast.Statement, env *Environment) (ast.Value, error) {
	var res ast.Value = ast.NilValue{}
	var err error
	ifState := IfState{}
	for _, stmt := range stmts {
		res, err = evalStmt(stmt, &ifState, env)
		if err != nil {
			return ast.NilValue{}, err
		}
		if isReturnValue(res) {
			break
		}
	}
	return res, err
}

func evalRule(rule ast.Rule, params []ast.Value, parent *Environment) (ast.Value, error) {
	env := parent.fork()
	err := verifyAndAddParamsToEnv(rule.Selector.Atrributes, params, env)
	if err != nil {
		return ast.NilValue{}, err
	}

	res, err := evalStatementList(rule.Body, env)
	if err != nil {
		return ReturnValue{
			Value: ast.NilValue{},
		}, err
	}

	if isReturnValue(res) {
		return res.(ReturnValue).Value, nil
	}

	return ast.NilValue{}, nil
}
