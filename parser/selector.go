package parser

import (
	. "github.com/shreyassanthu77/cisp/ast"
	"github.com/shreyassanthu77/cisp/lexer"
)

func (p *Parser) parseSelector(ident Identifier) (Selector, error) {
	next, err := p.peek()
	if err != nil {
		return Selector{}, err
	}

	if next.Typ == lexer.TOK_LBRACKET {
		attr, err := p.parseAttributes()
		if err != nil {
			return Selector{}, err
		}

		span := ident.Span
		span.End = attr[len(attr)-1].Span.End

		return Selector{
			Identifier: ident,
			Atrributes: attr,
			Span:       span,
		}, nil
	}

	next, err = p.peek()
	if err != nil {
		return Selector{}, err
	}

	if next.Typ == lexer.TOK_IDENTIFIER {
		panic("Complex selectors not implemented")
	}

	return Selector{
		Identifier: ident,
		Atrributes: nil,
		Span:       ident.Span,
	}, nil
}

func (p *Parser) parseAttributes() ([]Attreibute, error) {
	attrs := []Attreibute{}

	for {
		next, err := p.peek()
		if err != nil {
			return nil, err
		}

		if next.Typ != lexer.TOK_LBRACKET {
			break
		}

		p.next() // Consume '['
		id, err := p.expect(lexer.TOK_IDENTIFIER)
		if err != nil {
			return nil, err
		}

		next, err = p.peek()
		if err != nil {
			return nil, err
		}

		attrId := Identifier{
			Name: id.Value,
			Span: id.Span,
		}

		attr := Attreibute{
			Name:    attrId,
			Default: NilValue{},
			Span:    id.Span,
		}

		if next.Typ == lexer.TOK_EQUAL {
			p.next() // Consume '='
			val, err := p.parseLiteralVal()
			if err != nil {
				return nil, err
			}

			attr.Default = val
			attr.Span.End = val.GetSpan().End
		}

		attrs = append(attrs, attr)
		_, err = p.expect(lexer.TOK_RBRACKET)
		if err != nil {
			return nil, err
		}
	}

	return attrs, nil
}
