package lexer

import (
	"interpreter/token"
) // Import your own token package

type Lexer struct {
	source       string
	position     int  // Points to the current character
	nextPosition int  // Points to the next character in the input
	char         byte // Current character under examination
}

// New function to create a new Lexer instance
func New(source string) *Lexer {
	lex := &Lexer{
		source: source,
	}
	lex.ReadCharVal() // Initialize the lexer by reading the first character
	return lex
}

// ReadCharVal reads the next character in the input
func (l *Lexer) ReadCharVal() {
	if l.nextPosition >= len(l.source) {
		l.char = 0 // Null character to indicate end of input
	} else {
		l.char = l.source[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

func Isletter(c rune) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func IsDigit(c rune) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.ReadCharVal()
	}
}

func (lex *Lexer) ReadIdentifier() string {
	start_position := lex.position
	for Isletter(rune(lex.char)) {
		lex.ReadCharVal()
	}
	return lex.source[start_position:lex.position]
}

func (lex *Lexer) ReadDigit() string {
	start_position := lex.position
	for IsDigit(rune(lex.char)) {
		lex.ReadCharVal()
	}
	return lex.source[start_position:lex.position]
}

// newToken generates a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,  // This should be the type (e.g., token.PLUS)
		Literal: string(ch), // This should be the actual character (e.g., "+")
	}
}

func (lex *Lexer) NextToken() token.Token {
	var tokn token.Token
	lex.skipWhitespace()
	switch lex.char {
	case '=':
		tokn = newToken(token.ASSIGN, lex.char)
	case ';':
		tokn = newToken(token.SEMICOLON, lex.char)
	case '(':
		tokn = newToken(token.LPAREN, lex.char)
	case ')':
		tokn = newToken(token.RPAREN, lex.char)
	case ',':
		tokn = newToken(token.COMMA, lex.char)
	case '+':
		tokn = newToken(token.PLUS, lex.char)
	case '{':
		tokn = newToken(token.LBRACE, lex.char)
	case '}':
		tokn = newToken(token.RBRACE, lex.char)
	case 0:
		tokn.Literal = ""
		tokn.Type = token.EOF
	default:
		if Isletter(rune(lex.char)) {
			tokn.Literal = lex.ReadIdentifier()
			tokn.Type = token.LookIdentity(tokn.Literal)
			return tokn

		} else if IsDigit(rune(lex.char)) {
			tokn.Type = token.INT
			tokn.Literal = lex.ReadDigit()
			return tokn

		} else {
			tokn = newToken(token.ILLEGAL, lex.char)
		}
	}
	lex.ReadCharVal()
	return tokn
}
