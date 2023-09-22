package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	INT   = "INT"
	IDENT = "IDENT"

	// Operators
	PLUS     = "+"
	ASSIGN   = "="
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	EQ       = "=="
	NOT_EQ   = "!="

	LT = "<"
	GT = ">"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
	IF       = "IF"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
