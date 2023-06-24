package parser

import (
	"fmt"
	"unicode"
)

type Token struct {
	Type  string
	Value string
}

func isOperator(char string) bool {
	return char == "+" || char == "-" || char == "*" || char == "/"
}

func Lexer(input string) ([]Token, error) {
	var tokens []Token
	var err error
	i := 0

	for i < len(input) {
		char := string(input[i])

		switch {
		case char == "(":
			tokens = append(tokens, Token{
				"left_paren",
				char,
			})
			i++
		case char == ")":
			tokens = append(tokens, Token{
				"right_paren",
				char,
			})
			i++
		case isOperator(char):
			j := i
			for ; j < len(input) && isOperator(string(input[j])); j++ {
			}
			tokens = append(tokens, Token{
				"operator",
				input[i:j],
			})
			i = j
		case unicode.IsDigit(rune(char[0])):
			j := i
			for ; j < len(input) && unicode.IsDigit(rune(input[j])); j++ {
			}
			tokens = append(tokens, Token{
				"number",
				input[i:j],
			})
			i = j
		case unicode.IsLetter(rune(char[0])):
			j := i
			for ; j < len(input) && unicode.IsLetter(rune(input[j])); j++ {
			}
			tokens = append(tokens, Token{
				"identifier",
				input[i:j],
			})
			i = j
		case unicode.IsSpace(rune(char[0])):
			i++
		default:
			err = fmt.Errorf("unknown token: %s", char)
			return tokens, err
		}
	}
	return tokens, err
}
