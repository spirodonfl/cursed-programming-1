package interpreter

import (
	"fmt"

	"github.com/shreyassanthu77/cisp/ast"
	"github.com/shreyassanthu77/cisp/lexer"
)

type ReturnValue struct {
	Value ast.Value
}

func (r ReturnValue) IsValue() {}

func (r ReturnValue) GetSpan() lexer.Span {
	return r.Value.GetSpan()
}

func isReturnValue(v ast.Value) bool {
	_, ok := v.(ReturnValue)
	return ok
}

type IfState struct {
	IsIf         bool
	ShouldBranch bool
}

func (i *IfState) reset() {
	i.IsIf = false
	i.ShouldBranch = false
}

func evalIfRule(env *Environment, ifState *IfState, rule ast.AtRule) (ast.Value, error) {
	if len(rule.Parameters) != 1 {
		return ast.NilValue{}, fmt.Errorf("if rules should have exactly one parameter")
	}

	condition, err := evalValue(rule.Parameters[0], env)
	if err != nil {
		return ast.NilValue{}, err
	}

	conditionResult, ok := condition.(ast.Boolean)
	if !ok {
		return ast.NilValue{}, fmt.Errorf("if rule condition must evaluate to a boolean")
	}

	if conditionResult.Value {
		ifState.ShouldBranch = false
		return evalStatementList(rule.Body, env)
	}
	ifState.ShouldBranch = true

	return ast.NilValue{}, nil
}

func evalReturnRule(env *Environment, rule ast.AtRule) (ast.Value, error) {
	if len(rule.Parameters) != 1 {
		return ast.NilValue{}, fmt.Errorf("return rules should have exactly one parameter")
	}

	value, err := evalValue(rule.Parameters[0], env)
	if err != nil {
		return ast.NilValue{}, err
	}

	return ReturnValue{Value: value}, nil
}

func evalAtRule(env *Environment, ifState *IfState, rule ast.AtRule) (ast.Value, error) {
	switch rule.Name {
	case "if":
		ifState.IsIf = true
		return evalIfRule(env, ifState, rule)
	case "elif":
		if !ifState.IsIf {
			return ast.NilValue{}, fmt.Errorf("elif rule must be preceded by an if rule")
		}
		if ifState.ShouldBranch {
			return evalIfRule(env, ifState, rule)
		}
		return ast.NilValue{}, nil
	case "else":
		if !ifState.IsIf {
			return ast.NilValue{}, fmt.Errorf("else rule must be preceded by an if rule")
		}
		if ifState.ShouldBranch {
			ifState.reset()
			return evalStatementList(rule.Body, env)
		}
		ifState.reset()
		return ast.NilValue{}, nil
	case "return":
		return evalReturnRule(env, rule)
	}

	return ast.NilValue{}, fmt.Errorf("at rules are not supported yet")
}
