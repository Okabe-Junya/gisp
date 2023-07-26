package parser

import (
	"testing"
)

func TestLexer(t *testing.T) {
	cases := []struct {
		input           string
		expected        []Token
		isexpectedError bool
	}{
		{
			"(+ 1 2)",
			[]Token{
				{"left_paren", "("},
				{"operator", "+"},
				{"number", "1"},
				{"number", "2"},
				{"right_paren", ")"},
			},
			false,
		},
		{
			"(+ 1 2))",
			[]Token{
				{"left_paren", "("},
				{"operator", "+"},
				{"number", "1"},
				{"number", "2"},
				{"right_paren", ")"},
				{"right_paren", ")"},
			},
			false,
		},
		{
			"(* 2 (/ 1 2))",
			[]Token{
				{"left_paren", "("},
				{"operator", "*"},
				{"number", "2"},
				{"left_paren", "("},
				{"operator", "/"},
				{"number", "1"},
				{"number", "2"},
				{"right_paren", ")"},
				{"right_paren", ")"},
			},
			false,
		},
		{
			"(+ 1 2) (+ 3 4)",
			[]Token{
				{"left_paren", "("},
				{"operator", "+"},
				{"number", "1"},
				{"number", "2"},
				{"right_paren", ")"},
				{"left_paren", "("},
				{"operator", "+"},
				{"number", "3"},
				{"number", "4"},
				{"right_paren", ")"},
			},
			false,
		},
		{
			"(((+ 1 2)))",
			[]Token{
				{"left_paren", "("},
				{"left_paren", "("},
				{"left_paren", "("},
				{"operator", "+"},
				{"number", "1"},
				{"number", "2"},
				{"right_paren", ")"},
				{"right_paren", ")"},
				{"right_paren", ")"},
			},
			false,
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			tokens, err := Lexer(tt.input)
			if tt.isexpectedError && err == nil {
				t.Errorf("expected error, but error is nil")
			} else if !tt.isexpectedError && err != nil {
				t.Errorf("expected no error, but error got: %s", err)
			}

			if len(tokens) != len(tt.expected) {
				t.Errorf("expected %d tokens, got %d", len(tt.expected), len(tokens))
			}

			for i, token := range tokens {
				if token.Type != tt.expected[i].Type {
					t.Errorf("expected %s, got %s", tt.expected[i].Type, token.Type)
				}
				if token.Value != tt.expected[i].Value {
					t.Errorf("expected %s, got %s", tt.expected[i].Value, token.Value)
				}
			}
		})
	}
}
