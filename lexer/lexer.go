package lexer

import (
	"interpreter/token"
) 

// Lexer struct defines the lexer with necessary fields
// - `source`: the input source code as a string
// - `position`: current position in the input (points to current character)
// - `nextPosition`: next position in the input (points to the next character)
// - `char`: the current character being examined
type Lexer struct {
	source       string
	position     int  // Points to the current character
	nextPosition int  // Points to the next character in the input
	char         byte // Current character under examination
}

// New function to create a new Lexer instance
// - Initializes the lexer with the input source code
// - Calls ReadCharVal() to set the first character
func New(source string) *Lexer {
	lex := &Lexer{
		source: source,
	}
	lex.ReadCharVal() // Initialize the lexer by reading the first character
	return lex
}

// ReadCharVal reads the next character from the input source
// - If the lexer reaches the end of the input, it sets the character to 0 (null)
// - Otherwise, it advances `position` and `nextPosition` to the next character
func (l *Lexer) ReadCharVal() {
	if l.nextPosition >= len(l.source) {
		l.char = 0 // Null character to indicate end of input
	} else {
		l.char = l.source[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition += 1
}

// Isletter checks if the character is a letter (a-z, A-Z, or '_')
// - Returns true if the character is a valid letter, otherwise false
func IsApha(c rune) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

// IsDigit checks if the character is a digit (0-9)
// - Returns true if the character is a digit, otherwise false
func IsInteger(c rune) bool {
	return '0' <= c && c <= '9'
}

// skipWhitespace skips over whitespace characters (spaces, tabs, newlines, etc.)
// - Advances the lexer until a non-whitespace character is found
func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.ReadCharVal()
	}
}

// ReadIdentifier reads a complete identifier (e.g., variable names or keywords)
// - Starts at the current position and reads until a non-letter character is encountered
// - Returns the identifier as a string
func (lex *Lexer) ReadIdentifier() string {
	start_position := lex.position
	for IsApha(rune(lex.char)) {
		lex.ReadCharVal()
	}
	return lex.source[start_position:lex.position]
}

// ReadDigit reads a complete integer (e.g., 12345)
// - Starts at the current position and reads until a non-digit character is encountered
// - Returns the integer as a string
func (lex *Lexer) ReadInt() string {
	start_position := lex.position
	for IsInteger(rune(lex.char)) {
		lex.ReadCharVal()
	}
	return lex.source[start_position:lex.position]
}

// newToken generates a new token with the specified token type and literal character
// - `tokenType`: the type of the token (e.g., token.PLUS, token.INT)
// - `ch`: the character literal (e.g., '+', '5')
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,  // This should be the type (e.g., token.PLUS)
		Literal: string(ch), // This should be the actual character (e.g., "+")
	}
}

// NextToken generates the next token from the input source
// - Skips whitespace, then determines the type of token based on the current character
// - Handles operators, delimiters, and identifiers or numbers
func (lex *Lexer) NextToken() token.Token {
	var tokn token.Token
	lex.skipWhitespace() // Skip any whitespace before identifying the next token

	// Switch statement to handle single-character tokens
	switch lex.char {
	case '=':
		tokn = newToken(token.ASSIGN, lex.char)   // Assignment operator
	case ';':
		tokn = newToken(token.SEMICOLON, lex.char) // Semicolon
	case '(':
		tokn = newToken(token.LPAREN, lex.char)    // Left parenthesis
	case ')':
		tokn = newToken(token.RPAREN, lex.char)    // Right parenthesis
	case ',':
		tokn = newToken(token.COMMA, lex.char)     // Comma
	case '+':
		tokn = newToken(token.PLUS, lex.char)      // Plus operator
	case '{':
		tokn = newToken(token.LBRACE, lex.char)    // Left brace
	case '}':
		tokn = newToken(token.RBRACE, lex.char)    // Right brace
	case 0: // End of file (EOF)
		tokn.Literal = ""
		tokn.Type = token.EOF
	default: // Handle identifiers (e.g., variables) and numbers
		if IsApha(rune(lex.char)) {
			// If it's a letter, read an identifier or keyword
			tokn.Literal = lex.ReadIdentifier()
			tokn.Type = token.LookIdentity(tokn.Literal) // Check if it's a keyword (e.g., "let")
			return tokn
		} else if IsInteger(rune(lex.char)) {
			// If it's a digit, read a full number
			tokn.Type = token.INT
			tokn.Literal = lex.ReadInt()
			return tokn
		} else {
			// If it's an unrecognized character, mark it as ILLEGAL
			tokn = newToken(token.ILLEGAL, lex.char)
		}
	}

	// Move to the next character after tokenizing the current one
	lex.ReadCharVal()
	return tokn
}
