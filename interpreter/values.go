package interpreter

import (
	"fmt"

	"github.com/shreyassanthu77/cisp/ast"
)

func evalUnaryOp(op ast.UnaryOp, env *Environment) (ast.Value, error) {
	val, err := evalValue(op.Value, env)
	if err != nil {
		return ast.NilValue{}, err
	}

	switch op.Op {
	case "+":
		return val, nil
	case "-":
		switch val := val.(type) {
		case ast.Int:
			return ast.Int{Value: -val.Value}, nil
		case ast.Float:
			return ast.Float{Value: -val.Value}, nil
		}
	case "!":
		switch val := val.(type) {
		case ast.Boolean:
			return ast.Boolean{Value: !val.Value}, nil
		default:
			return ast.NilValue{}, fmt.Errorf("invalid type for unary operator !: %T", val)
		}
	}

	return ast.NilValue{}, fmt.Errorf("invalid unary operator %s", op.Op)
}

func evalValue(value ast.Value, env *Environment) (ast.Value, error) {
	switch value := value.(type) {
	case ast.FunctionCall:
		return evalFnCall(value, env)
	case ast.Identifier:
		_, err := env.getVar(value.Name)
		if err != nil {
			_, err := env.genFn(value.Name)
			if err == nil {
				return ast.NilValue{}, fmt.Errorf("You cannot use a function as a value use %s() instead of %s if you want to call it", value.Name, value.Name)
			}
			return ast.NilValue{}, fmt.Errorf("Literal Identifiers are not allowed use $variable if you want to use a variable")
		}
		return ast.NilValue{}, fmt.Errorf("Literal Identifiers are not allowed use $%s instead of %s", value.Name, value.Name)
	case ast.Int, ast.Float, ast.String, ast.Boolean, ast.NilValue:
		return value, nil
	case ast.UnaryOp:
		return evalUnaryOp(value, env)
	case ast.BinaryOp:
		return evalBinaryOp(value, env)
	case ast.VarianleDerefValue:
		val, err := env.getVar(value.Variable.Name)
		if err != nil {
			return ast.NilValue{}, err
		}
		return val, nil
	}

	return ast.NilValue{}, fmt.Errorf("invalid value type: %T", value)
}

func evalFnCall(fnCall ast.FunctionCall, env *Environment) (ast.Value, error) {
	fn, err := env.genFn(fnCall.Fn.Name)
	if err != nil {
		return ast.NilValue{}, err
	}

	params := make([]ast.Value, len(fnCall.Parameters))
	for i, param := range fnCall.Parameters {
		params[i], err = evalValue(param, env)
		if err != nil {
			return ast.NilValue{}, err
		}
	}

	return evalRule(fn, params, env)
}
