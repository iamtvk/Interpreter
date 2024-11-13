package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var KeywordMap = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIden(ident string) TokenType {
	if tok, ok := KeywordMap[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LSQRBRACE = "["
	RSQRBRACE = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
