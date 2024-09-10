package token

// allows us to differenciate between various types of tokens used
type TokenType string

// Here we call the string we have declared above and a string for defining the various token
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF" // End Of File

	IDENT   = "IDENT" // identifier (e.g., function names like "add")
	INT     = "INT"   // integer literals
	STRING  = "STRING"
	BOOLEAN = "BOOLEAN"

	// Operators
	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	SUB      = "SUB"
	ASTERICK = "ASTERICK"
	SLASH    = "SLASH"
	BANG     = "BANG"

	// Delimiters
	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	COLON     = "COLON"
	LBRACKET  = "LBRACKET"
	RBRACKET  = "RBRACKET"
	LBRACE    = "LBRACE"
	RBRACE    = "RBRACE"
	LPAREN    = "LPAREN"
	RPAREN    = "RPAREN"
	GT        = "GT"
	LT        = "LT"

	// keywords- words that have a preserved meaning around them and can't be used as variables
	FUNCTION  = "FUNCTION"
	LET       = "LET"
	STRUCTURE = "STRUCTURE"
	IF        = "IF"
	ELSE      = "ELSE"
	FOR       = "FOR"
	WHILE     = "WHILE"
	RETURN    = "RETURN"
	DO        = "DO"
	RANGE     = "RANGE"
)

/*
This variable maps the various token we have declared above to what we would like to call them,
essetially you can see the difference as per the strings we've choosen to use.
Example 'else' is 'el', 'return' is 'output', 'func' is 'funct'
We are free to make as much changes as we would like so long as we can remember while using them. ;)
*/
var keywords = map[string]TokenType{
	"funct":   FUNCTION,
	"let":     LET,
	"struct":  STRUCTURE,
	"if":      IF,
	"el":      ELSE,
	"for":     FOR,
	"while":   WHILE,
	"output":  RETURN,
	"do":      DO,
	"through": RANGE,
}

func LookIdentity(id string) TokenType {
	if tokentype, ok := keywords[id]; ok {
		return tokentype
	}
	return IDENT
}

