package ast

import (
	"fmt"

	"github.com/shreyassanthu77/cisp/lexer"
)

type Value interface {
	IsValue()
	GetSpan() lexer.Span
}

type NilValue struct {
	Span lexer.Span
}

func (n NilValue) IsValue() {}

func (n NilValue) GetSpan() lexer.Span {
	return n.Span
}

func (n NilValue) String() string {
	return "Nil"
}

type Identifier struct {
	Name string
	Span lexer.Span
}

func (i Identifier) IsValue() {}

func (i Identifier) GetSpan() lexer.Span {
	return i.Span
}

func (i Identifier) String() string {
	return fmt.Sprintf("Id(%s)", i.Name)
}

type String struct {
	Value string
	Span  lexer.Span
}

func (s String) IsValue() {}

func (s String) GetSpan() lexer.Span {
	return s.Span
}

func (s String) String() string {
	return fmt.Sprintf("'%s'", s.Value)
}

type Int struct {
	Value int64
	Span  lexer.Span
}

func (n Int) IsValue() {}

func (n Int) GetSpan() lexer.Span {
	return n.Span
}

func (n Int) String() string {
	return fmt.Sprintf("%d", n.Value)
}

type Float struct {
	Value float64
	Span  lexer.Span
}

func (n Float) IsValue() {}

func (n Float) GetSpan() lexer.Span {
	return n.Span
}

func (n Float) String() string {
	return fmt.Sprintf("%f", n.Value)
}

type Boolean struct {
	Value bool
	Span  lexer.Span
}

func (b Boolean) IsValue() {}

func (b Boolean) GetSpan() lexer.Span {
	return b.Span
}

func (b Boolean) String() string {
	return fmt.Sprintf("Bool(%t)", b.Value)
}

type VarianleDerefValue struct {
	Variable Identifier
	Span     lexer.Span
}

func (v VarianleDerefValue) IsValue() {}

func (v VarianleDerefValue) GetSpan() lexer.Span {
	return v.Span
}

func (v VarianleDerefValue) String() string {
	return fmt.Sprintf("var(%s)", v.Variable.Name)
}

type Statement interface {
	IsStatement()
	GetSpan() lexer.Span
}

type Declaration struct {
	Property   Identifier
	Parameters []Value
	Span       lexer.Span
}

func (d Declaration) IsStatement() {}

func (d Declaration) GetSpan() lexer.Span {
	return d.Span
}

type FunctionCall struct {
	Fn         Identifier
	Parameters []Value
	Span       lexer.Span
}

func (f FunctionCall) IsValue() {}

func (f FunctionCall) GetSpan() lexer.Span {
	return f.Span
}

func (f FunctionCall) String() string {
	return fmt.Sprintf("Call(%s, %v)", f.Fn, f.Parameters)
}

type UnaryOp struct {
	Op    string
	Value Value
	Span  lexer.Span
}

func (u UnaryOp) IsValue() {}

func (u UnaryOp) GetSpan() lexer.Span {
	return u.Span
}

type BinaryOp struct {
	Left  Value
	Op    string
	Right Value
	Span  lexer.Span
}

func (b BinaryOp) IsValue() {}

func (b BinaryOp) GetSpan() lexer.Span {
	return b.Span
}

type AtRule struct {
	Name       string
	Parameters []Value
	Body       []Statement
	Span       lexer.Span
}

func (r AtRule) isRule() {}

func (r AtRule) IsStatement() {}

func (r AtRule) GetSpan() lexer.Span {
	return r.Span
}

type Attreibute struct {
	Name    Identifier
	Default Value
	Span    lexer.Span
}

type Selector struct {
	Identifier Identifier
	Atrributes []Attreibute
	Span       lexer.Span
}

type IRule interface {
	isRule()
}

type Rule struct {
	Selector Selector
	Body     []Statement
	Span     lexer.Span
}

func (r Rule) isRule() {}

func (r Rule) IsStatement() {}

func (r Rule) GetSpan() lexer.Span {
	return r.Span
}

type Program struct {
	Rules []IRule
}
