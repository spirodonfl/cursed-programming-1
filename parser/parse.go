package parser

import (
	. "github.com/shreyassanthu77/cisp/ast"
	"github.com/shreyassanthu77/cisp/lexer"
)

func (p *Parser) parseFunctionCall(tok lexer.Token) (FunctionCall, error) {
	name := tok.Value
	_, err := p.expect(lexer.TOK_LPAREN)
	if err != nil {
		return FunctionCall{}, err
	}

	fn := Identifier{
		Name: name,
		Span: tok.Span,
	}

	params := []Value{}
	next, err := p.peek()
	if err != nil {
		return FunctionCall{}, err
	}

	if next.Typ == lexer.TOK_RPAREN {
		p.next()
		return FunctionCall{
			Fn:         fn,
			Parameters: params,
			Span:       tok.Span,
		}, nil
	}

	for {
		param, err := p.parseValue()
		if err != nil {
			return FunctionCall{}, err
		}
		params = append(params, param)

		next, err = p.peek()
		if err != nil {
			return FunctionCall{}, err
		}

		if next.Typ != lexer.TOK_COMMA {
			break
		}
		p.next()
		next, err = p.peek()
	}

	_, err = p.expect(lexer.TOK_RPAREN)
	if err != nil {
		return FunctionCall{}, err
	}

	span := tok.Span
	span.End = params[len(params)-1].GetSpan().End

	return FunctionCall{
		Fn:         fn,
		Parameters: params,
		Span:       span,
	}, nil
}

func (p *Parser) parseDeclarationStmt(id Identifier) (Declaration, error) {
	_, err := p.expect(lexer.TOK_COLON)
	if err != nil {
		return Declaration{
			Property: id,
		}, err
	}

	span := id.Span

	values := []Value{}
	for {
		next, err := p.peek()
		if err != nil {
			return Declaration{}, err
		}

		if next.Typ == lexer.TOK_SEMICOLON {
			break
		}

		value, err := p.parseValue()
		if err != nil {
			return Declaration{}, err
		}

		values = append(values, value)
	}

	semi, err := p.expect(lexer.TOK_SEMICOLON)
	if err != nil {
		return Declaration{}, err
	}

	span.End = semi.Span.End

	return Declaration{
		Property:   id,
		Parameters: values,
		Span:       span,
	}, nil
}

func (p *Parser) parseStatement() (Statement, error) {
	next, err := p.peek()
	if err != nil {
		return nil, err
	}

	if next.Typ == lexer.TOK_AT {
		return p.parseAtRule()
	}

	id, err := p.expect(lexer.TOK_IDENTIFIER)
	if err != nil {
		return nil, err
	}

	next, err = p.peek()
	if err != nil {
		return nil, err
	}

	stmtId := Identifier{
		Name: id.Value,
		Span: id.Span,
	}

	if next.Typ == lexer.TOK_COLON {
		return p.parseDeclarationStmt(stmtId)
	}

	return p.parseNestedRule(stmtId)
}

func (p *Parser) parseDeclarationBlock() ([]Statement, lexer.Span, error) {
	lsq, err := p.expect(lexer.TOK_LSQUIRLY)
	if err != nil {
		return nil, lexer.Span{}, err
	}

	span := lsq.Span

	stmts := []Statement{}
	var lastErr error
	for {
		next, err := p.peek()
		if err != nil {
			return nil, lexer.Span{}, err
		}

		if next.Typ == lexer.TOK_RSQUIRLY {
			break
		}

		stmt, err := p.parseStatement()
		if err != nil {
			lastErr = err
			break
		}

		stmts = append(stmts, stmt)
	}

	if lastErr != nil {
		return nil, lexer.Span{}, lastErr
	}

	rsq, err := p.expect(lexer.TOK_RSQUIRLY)
	if err != nil {
		return nil, lexer.Span{}, err
	}

	span.End = rsq.Span.End

	return stmts, span, nil
}

func (p *Parser) parseNestedRule(id Identifier) (Rule, error) {
	selector, err := p.parseSelector(id)
	if err != nil {
		return Rule{}, err
	}

	decls, declSpan, err := p.parseDeclarationBlock()
	if err != nil {
		return Rule{}, err
	}

	span := selector.Span
	span.End = declSpan.End

	if len(decls) > 0 {
		span.End = decls[len(decls)-1].GetSpan().End
	}

	return Rule{
		Selector: selector,
		Body:     decls,
		Span:     span,
	}, nil
}

func (p *Parser) parseRule() (Rule, error) {
	id, err := p.expect(lexer.TOK_IDENTIFIER)
	if err != nil {
		return Rule{}, err
	}

	ruleId := Identifier{
		Name: id.Value,
		Span: id.Span,
	}

	res, err := p.parseNestedRule(ruleId)
	if err != nil {
		return Rule{}, err
	}

	return res, nil
}

func (p *Parser) parseAtRule() (AtRule, error) {
	at, _ := p.next() // Consume '@'
	span := at.Span

	name, err := p.expect(lexer.TOK_IDENTIFIER)
	if err != nil {
		return AtRule{}, err
	}

	span.End = name.Span.End

	params := []Value{}
	for {
		next, err := p.peek()
		if err != nil {
			return AtRule{}, err
		}

		if next.Typ == lexer.TOK_LSQUIRLY || next.Typ == lexer.TOK_SEMICOLON {
			break
		}

		param, err := p.parseValue()
		if err != nil {
			return AtRule{}, err
		}

		params = append(params, param)
	}

	if len(params) > 0 {
		span.End = params[len(params)-1].GetSpan().End
	}

	next, err := p.peek()
	if err != nil {
		return AtRule{}, err
	}

	if next.Typ == lexer.TOK_LSQUIRLY {
		decls, declSpan, err := p.parseDeclarationBlock()
		if err != nil {
			return AtRule{}, err
		}

		span.End = declSpan.End

		return AtRule{
			Name:       name.Value,
			Parameters: params,
			Body:       decls,
			Span:       span,
		}, nil
	}

	_, err = p.expect(lexer.TOK_SEMICOLON)
	if err != nil {
		return AtRule{}, err
	}

	return AtRule{
		Name:       name.Value,
		Parameters: params,
		Span:       span,
	}, nil
}
