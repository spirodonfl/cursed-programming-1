package parser

import (
	"fmt"

	. "github.com/shreyassanthu77/cisp/ast"
	"github.com/shreyassanthu77/cisp/lexer"
)

type Parser struct {
	lex    *lexer.Lexer
	tok    lexer.Token
	hasTok bool
}

func New(lex *lexer.Lexer) *Parser {
	return &Parser{
		lex:    lex,
		hasTok: false,
	}
}

func (p *Parser) peek() (lexer.Token, error) {
	if p.hasTok {
		return p.tok, nil
	}

	tok, err := p.lex.Next()
	if err != nil {
		return lexer.Token{}, err
	}

	p.tok = tok
	p.hasTok = true
	return tok, nil
}

func (p *Parser) next() (lexer.Token, error) {
	tok, err := p.peek()
	if err != nil {
		return lexer.Token{}, err
	}
	p.hasTok = false
	return tok, nil
}

func (p *Parser) expect(typ string) (lexer.Token, error) {
	tok, err := p.next()
	if err != nil {
		return lexer.Token{}, err
	}
	if tok.Typ != typ {
		return lexer.Token{}, fmt.Errorf("%d:%d Expected %s but got %s", tok.Span.Start.Line, tok.Span.Start.Col, typ, tok.Typ)
	}
	return tok, nil
}

func (p *Parser) Parse() (Program, error) {
	rules := []IRule{}
	var lastErr error
outer:
	for {
		next, err := p.peek()

		// Skip semicolons
		for {
			if err != nil {
				return Program{}, err
			}

			if next.Typ == lexer.EOF {
				break outer
			}

			if next.Typ == lexer.TOK_SEMICOLON {
				p.next()
				next, err = p.peek()
				continue
			}
			break
		}

		if next.Typ == lexer.TOK_AT {
			atRule, err := p.parseAtRule()
			if err != nil {
				lastErr = err
				break
			}
			rules = append(rules, atRule)
		} else {
			rule, err := p.parseRule()
			if err != nil {
				lastErr = err
				break
			}
			rules = append(rules, rule)
		}
	}

	if lastErr != nil {
		return Program{}, lastErr
	}
	return Program{
		Rules: rules,
	}, nil
}
