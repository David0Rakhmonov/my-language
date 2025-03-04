package lexer

import (
	"strings"
	"unicode"
)

type TokenType string

const (
	Number  TokenType = "Number"
	String  TokenType = "String"
	Keyword TokenType = "Keyword"
	Symbol  TokenType = "Symbol"
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(code string) []Token {
	var tokens []Token
	var buffer strings.Builder

	for _, char := range code {
		if unicode.IsSpace(char) {
			if buffer.Len() > 0 {
				tokens = append(tokens, Token{Type: Keyword, Value: buffer.String()})
				buffer.Reset()
			}
			continue
		}

		if unicode.IsDigit(char) || unicode.IsLetter(char) {
			buffer.WriteRune(char)
		} else {
			if buffer.Len() > 0 {
				tokens = append(tokens, Token{Type: Keyword, Value: buffer.String()})
				buffer.Reset()
			}
			tokens = append(tokens, Token{Type: Symbol, Value: string(char)})
		}
	}

	if buffer.Len() > 0 {
		tokens = append(tokens, Token{Type: Keyword, Value: buffer.String()})
	}

	return tokens
}
