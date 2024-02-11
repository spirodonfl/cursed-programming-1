package parser

import (
	"fmt"
	"strconv"

	. "github.com/shreyassanthu77/cisp/ast"
	"github.com/shreyassanthu77/cisp/lexer"
)

func (p *Parser) parseLiteralVal() (Value, error) {
	tok, err := p.next()
	if err != nil {
		return nil, err
	}

	switch tok.Typ {
	case lexer.TOK_IDENTIFIER:
		next, err := p.peek()
		if err != nil {
			return nil, err
		}
		if next.Typ == lexer.TOK_LPAREN {
			return p.parseFunctionCall(tok)
		}
		return Identifier{Name: tok.Value, Span: tok.Span}, nil
	case lexer.TOK_STRING:
		return String{Value: tok.Value, Span: tok.Span}, nil
	case lexer.TOK_INT:
		f, err := strconv.ParseInt(tok.Value, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse number: %s", err)
		}
		return Int{Value: f, Span: tok.Span}, nil
	case lexer.TOK_FLOAT:
		f, err := strconv.ParseFloat(tok.Value, 64)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse number: %s", err)
		}
		return Float{Value: f, Span: tok.Span}, nil
	case lexer.TOK_TRUE:
		return Boolean{Value: true, Span: tok.Span}, nil
	case lexer.TOK_FALSE:
		return Boolean{Value: false, Span: tok.Span}, nil
	case lexer.TOK_EMPTY:
		return NilValue{Span: tok.Span}, nil
	case lexer.TOK_LPAREN:
		val, err := p.parseValue()
		if err != nil {
			return nil, err
		}

		_, err = p.expect(lexer.TOK_RPAREN)
		if err != nil {
			return nil, err
		}

		return val, nil
	}

	return nil, fmt.Errorf("Unexpected token: %s", tok.Typ)
}

func (p *Parser) parseUnaryExpression() (Value, error) {
	next, err := p.peek()
	if err != nil {
		return nil, err
	}

	if next.Typ == lexer.TOK_DOLLAR {
		p.next() // Consume the operator
		ident, err := p.expect(lexer.TOK_IDENTIFIER)
		if err != nil {
			return nil, err
		}
		variable := Identifier{
			Name: ident.Value,
			Span: ident.Span,
		}
		return VarianleDerefValue{
			Variable: variable,
			Span:     ident.Span,
		}, nil
	}

	if next.Typ == lexer.TOK_BANG ||
		next.Typ == lexer.TOK_MINUS ||
		next.Typ == lexer.TOK_PLUS ||
		next.Typ == lexer.TOK_TILDE {
		p.next() // Consume the operator
		val, err := p.parseLiteralVal()
		if err != nil {
			return nil, err
		}
		span := next.Span
		span.End = val.GetSpan().End
		return UnaryOp{
			Op:    next.Value,
			Value: val,
			Span:  span,
		}, nil
	}

	return p.parseLiteralVal()
}

func (p *Parser) parseMultiplicativeExpr() (Value, error) {
	left, err := p.parseUnaryExpression()
	if err != nil {
		return BinaryOp{}, err
	}

	for {
		next, err := p.peek()
		if err != nil {
			return BinaryOp{}, err
		}

		if next.Typ != lexer.TOK_ASTERISK && next.Typ != lexer.TOK_SLASH {
			break
		}

		p.next() // Consume the operator

		right, err := p.parseUnaryExpression()
		if err != nil {
			return BinaryOp{}, err
		}

		span := left.GetSpan()
		span.End = right.GetSpan().End

		left = BinaryOp{
			Left:  left,
			Op:    next.Value,
			Right: right,
			Span:  span,
		}
	}

	return left, nil
}

func (p *Parser) parseAdditiveExpr() (Value, error) {
	left, err := p.parseMultiplicativeExpr()
	if err != nil {
		return BinaryOp{}, err
	}

	for {
		next, err := p.peek()
		if err != nil {
			return BinaryOp{}, err
		}

		if next.Typ != lexer.TOK_PLUS && next.Typ != lexer.TOK_MINUS {
			break
		}

		p.next() // Consume the operator

		right, err := p.parseMultiplicativeExpr()
		if err != nil {
			return BinaryOp{}, err
		}

		span := left.GetSpan()
		span.End = right.GetSpan().End

		left = BinaryOp{
			Left:  left,
			Op:    next.Value,
			Right: right,
			Span:  span,
		}
	}

	return left, nil
}

func (p *Parser) parseRelationalExpr() (Value, error) {
	left, err := p.parseAdditiveExpr()
	if err != nil {
		return BinaryOp{}, err
	}

	for {
		next, err := p.peek()
		if err != nil {
			return BinaryOp{}, err
		}

		if next.Typ != lexer.TOK_LESS_THAN && next.Typ != lexer.TOK_LESS_THAN_EQUAL &&
			next.Typ != lexer.TOK_GREATER_THAN && next.Typ != lexer.TOK_GREATER_THAN_EQUAL {
			break
		}

		p.next() // Consume the operator

		right, err := p.parseAdditiveExpr()
		if err != nil {
			return BinaryOp{}, err
		}

		span := left.GetSpan()
		span.End = right.GetSpan().End

		left = BinaryOp{
			Left:  left,
			Op:    next.Value,
			Right: right,
			Span:  span,
		}
	}

	return left, nil
}

func (p *Parser) parseEqualityExpr() (Value, error) {
	left, err := p.parseRelationalExpr()
	if err != nil {
		return BinaryOp{}, err
	}

	for {
		next, err := p.peek()
		if err != nil {
			return BinaryOp{}, err
		}

		if next.Typ != lexer.TOK_DOUBLE_EQUAL && next.Typ != lexer.TOK_NOT_EQUAL {
			break
		}

		p.next() // Consume the operator

		right, err := p.parseRelationalExpr()
		if err != nil {
			return BinaryOp{}, err
		}

		span := left.GetSpan()
		span.End = right.GetSpan().End

		left = BinaryOp{
			Left:  left,
			Op:    next.Value,
			Right: right,
			Span:  span,
		}
	}

	return left, nil
}

func (p *Parser) parseLogicalAndExpr() (Value, error) {
	left, err := p.parseEqualityExpr()
	if err != nil {
		return BinaryOp{}, err
	}

	for {
		next, err := p.peek()
		if err != nil {
			return BinaryOp{}, err
		}

		if next.Typ != lexer.TOK_AND {
			break
		}

		p.next() // Consume the operator

		right, err := p.parseEqualityExpr()
		if err != nil {
			return BinaryOp{}, err
		}

		span := left.GetSpan()
		span.End = right.GetSpan().End

		left = BinaryOp{
			Left:  left,
			Op:    next.Value,
			Right: right,
			Span:  span,
		}
	}

	return left, nil
}

func (p *Parser) parseLogicalOrExpr() (Value, error) {
	left, err := p.parseLogicalAndExpr()
	if err != nil {
		return BinaryOp{}, err
	}

	for {
		next, err := p.peek()
		if err != nil {
			return BinaryOp{}, err
		}

		if next.Typ != lexer.TOK_OR {
			break
		}

		p.next() // Consume the operator

		right, err := p.parseLogicalAndExpr()
		if err != nil {
			return BinaryOp{}, err
		}

		span := left.GetSpan()
		span.End = right.GetSpan().End

		left = BinaryOp{
			Left:  left,
			Op:    next.Value,
			Right: right,
			Span:  span,
		}
	}

	return left, nil
}

func (p *Parser) parseValue() (Value, error) {
	return p.parseLogicalOrExpr()
}
