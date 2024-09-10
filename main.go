package main

import (
	"fmt"
	"os"

	"interpreter/lexer"
)

func main() {
	byte, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	input := string(byte)
	lex := lexer.New(input)

	// Print out each token produced by the lexer until we hit the end of the input (EOF)
	for {
		tok := lex.NextToken() // Get the next token from the lexer
		fmt.Printf("Token: Type=  (%s)   , Literal=  (%s) \n", tok.Type, tok.Literal)

		// Stop printing tokens when we hit EOF
		if tok.Type == "EOF" {
			break
		}
	}
}
