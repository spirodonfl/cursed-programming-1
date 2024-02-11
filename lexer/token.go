package lexer

import "fmt"

const (
	EOF            = "EOF"
	TOK_IDENTIFIER = "IDENTIFIER"
	TOK_STRING     = "STRING"
	TOK_INT        = "INT"
	TOK_FLOAT      = "FLOAT"
	TOK_TRUE       = "TRUE"
	TOK_FALSE      = "FALSE"

	// Unary operators
	TOK_BANG  = "BANG"
	TOK_TILDE = "TILDE"

	// Binary operators
	TOK_PLUS               = "PLUS"
	TOK_MINUS              = "MINUS"
	TOK_ASTERISK           = "ASTERISK"
	TOK_SLASH              = "SLASH"
	TOK_PERCENT            = "PERCENT"
	TOK_CARET              = "CARET"
	TOK_EQUAL              = "EQUAL"
	TOK_DOUBLE_EQUAL       = "DOUBLE_EQUAL"
	TOK_NOT_EQUAL          = "NOT_EQUAL"
	TOK_GREATER_THAN       = "GREATER_THAN"
	TOK_GREATER_THAN_EQUAL = "GREATER_THAN_EQUAL"
	TOK_LESS_THAN          = "LESS_THAN"
	TOK_LESS_THAN_EQUAL    = "LESS_THAN_EQUAL"
	TOK_AND                = "AND"
	TOK_OR                 = "OR"
	TOK_DOLLAR             = "DOLLAR"

	// Expressions
	TOK_LPAREN = "LPAREN"
	TOK_RPAREN = "RPAREN"
	TOK_EMPTY  = "EMPTY" // empty tuple `()`
	TOK_COMMA  = "COMMA"

	// Selectors
	TOK_DOT      = "DOT"
	TOK_HASH     = "HASH"
	TOK_LBRACKET = "LBRACKET"
	TOK_RBRACKET = "RBRACKET"

	// rules
	TOK_AT        = "AT"
	TOK_LSQUIRLY  = "LSQUIRLY"
	TOK_RSQUIRLY  = "RSQUIRLY"
	TOK_COLON     = "COLON"
	TOK_SEMICOLON = "SEMICOLON"
	TOK_IMPORTANT = "IMPORTANT"
)

type Loc struct {
	Pos  int
	Line int
	Col  int
}

type Span struct {
	Start Loc
	End   Loc
}

type Token struct {
	Typ   string
	Value string
	Span  Span
}

func (t Token) String() string {
	line := t.Span.Start.Line
	col := t.Span.Start.Col
	return fmt.Sprintf("{%d:%d} TOK<%s>(%s)", line, col, t.Typ, t.Value)
}
