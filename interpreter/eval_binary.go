package interpreter

import (
	"fmt"

	"github.com/shreyassanthu77/cisp/ast"
)

type ValueType int

const (
	val_type_unknown ValueType = iota
	val_type_int
	val_type_float
	val_type_string
	val_type_boolean
	val_type_nil
)

func getValueType(v ast.Value) ValueType {
	switch v.(type) {
	case ast.Int:
		return val_type_int
	case ast.Float:
		return val_type_float
	case ast.String:
		return val_type_string
	case ast.Boolean:
		return val_type_boolean
	case ast.NilValue:
		return val_type_nil
	default:
		return val_type_unknown
	}
}

func evalAdd(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Int{Value: left.(ast.Int).Value + right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Float{Value: left.(ast.Float).Value + right.(ast.Float).Value}, nil
	case val_type_string:
		return ast.String{Value: left.(ast.String).Value + right.(ast.String).Value}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for addition: %T and %T", left, right)
	}
}

func evalSub(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Int{Value: left.(ast.Int).Value - right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Float{Value: left.(ast.Float).Value - right.(ast.Float).Value}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for subtraction: %T and %T", left, right)
	}
}

func evalMul(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Int{Value: left.(ast.Int).Value * right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Float{Value: left.(ast.Float).Value * right.(ast.Float).Value}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for multiplication: %T and %T", left, right)
	}
}

func evalDiv(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Int{Value: left.(ast.Int).Value / right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Float{Value: left.(ast.Float).Value / right.(ast.Float).Value}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for division: %T and %T", left, right)
	}
}

func evalEq(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Boolean{Value: left.(ast.Int).Value == right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Boolean{Value: left.(ast.Float).Value == right.(ast.Float).Value}, nil
	case val_type_string:
		return ast.Boolean{Value: left.(ast.String).Value == right.(ast.String).Value}, nil
	case val_type_boolean:
		return ast.Boolean{Value: left.(ast.Boolean).Value == right.(ast.Boolean).Value}, nil
	case val_type_nil:
		return ast.Boolean{Value: true}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for equality: %T and %T", left, right)
	}
}

func evalLt(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Boolean{Value: left.(ast.Int).Value < right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Boolean{Value: left.(ast.Float).Value < right.(ast.Float).Value}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for less than: %T and %T", left, right)
	}
}

func evalLe(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Boolean{Value: left.(ast.Int).Value <= right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Boolean{Value: left.(ast.Float).Value <= right.(ast.Float).Value}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for less than or equal: %T and %T", left, right)
	}
}

func evalGt(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Boolean{Value: left.(ast.Int).Value > right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Boolean{Value: left.(ast.Float).Value > right.(ast.Float).Value}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for greater than: %T and %T", left, right)
	}
}

func evalGe(left, right ast.Value, leftType ValueType) (ast.Value, error) {
	switch leftType {
	case val_type_int:
		return ast.Boolean{Value: left.(ast.Int).Value >= right.(ast.Int).Value}, nil
	case val_type_float:
		return ast.Boolean{Value: left.(ast.Float).Value >= right.(ast.Float).Value}, nil
	default:
		return ast.NilValue{}, fmt.Errorf("invalid types for greater than or equal: %T and %T", left, right)
	}
}

func evalAnd(left, right ast.Value) (ast.Value, error) {
	return ast.Boolean{Value: left.(ast.Boolean).Value && right.(ast.Boolean).Value}, nil
}

func evalOr(left, right ast.Value) (ast.Value, error) {
	return ast.Boolean{Value: left.(ast.Boolean).Value || right.(ast.Boolean).Value}, nil
}

func evalBinaryOp(op ast.BinaryOp, env *Environment) (ast.Value, error) {
	left, err := evalValue(op.Left, env)
	if err != nil {
		return ast.NilValue{}, err
	}

	right, err := evalValue(op.Right, env)
	if err != nil {
		return ast.NilValue{}, err
	}

	leftType := getValueType(left)
	rightType := getValueType(right)

	if leftType == val_type_unknown || rightType == val_type_unknown {
		return ast.NilValue{}, fmt.Errorf("invalid types for binary operation: %T and %T", left, right)
	}

	if leftType != rightType {
		if leftType == val_type_int && rightType == val_type_float {
			left = ast.Float{Value: float64(left.(ast.Int).Value)}
			leftType = val_type_float
		} else if leftType == val_type_float && rightType == val_type_int {
			right = ast.Float{Value: float64(right.(ast.Int).Value)}
			rightType = val_type_float
		} else {
			return ast.NilValue{}, fmt.Errorf("invalid types for binary operation: %T and %T", left, right)
		}
	}

	switch op.Op {
	case "+":
		return evalAdd(left, right, leftType)
	case "-":
		return evalSub(left, right, leftType)
	case "*":
		return evalMul(left, right, leftType)
	case "/":
		return evalDiv(left, right, leftType)
	case "<":
		return evalLt(left, right, leftType)
	case "<=":
		return evalLe(left, right, leftType)
	case ">":
		return evalGt(left, right, leftType)
	case ">=":
		return evalGe(left, right, leftType)
	case "==":
		return evalEq(left, right, leftType)
	case "!=":
		val, err := evalEq(left, right, leftType)
		if err != nil {
			return ast.NilValue{}, err
		}
		return ast.Boolean{Value: !val.(ast.Boolean).Value}, nil
	case "&&":
		if leftType != val_type_boolean || rightType != val_type_boolean {
			return ast.NilValue{}, fmt.Errorf("invalid types for &&: %T and %T", left, right)
		}
		return evalAnd(left, right)
	case "||":
		if leftType != val_type_boolean || rightType != val_type_boolean {
			return ast.NilValue{}, fmt.Errorf("invalid types for ||: %T and %T", left, right)
		}
		return evalOr(left, right)
	}

	panic(fmt.Sprintf("invalid binary operator %s", op.Op))
}
