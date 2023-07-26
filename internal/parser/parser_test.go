package parser

import "testing"

func TestParser(t *testing.T) {
	cases := []struct {
		input           []Token
		expected        Node
		isexpectedError bool
	}{
		{
			[]Token{
				{"left_paren", "("},
				{"operator", "+"},
				{"number", "1"},
				{"number", "2"},
				{"right_paren", ")"},
			},
			Operator{Op: "+", Left: Number{Value: 1}, Right: Number{Value: 2}},
			false,
		},
		// TODO: 以下はハックケースであるため、テストを通すためにはparser.goを修正する必要がある
		{
			[]Token{
				{"left_paren", "("},
				{"operator", "+"},
				{"number", "1"},
				{"number", "2"},
				{"right_paren", ")"},
				{"right_paren", ")"},
			},
			Operator{Op: "+", Left: Number{Value: 1}, Right: Number{Value: 2}},
			true,
		},
	}

	for _, tt := range cases {
		tt := tt
		actual, _, err := Parse(tt.input)
		if err != nil {
			t.Errorf("Parse(%v) returned error: %v", tt.input, err)
		}
		if actual != tt.expected {
			t.Errorf("Parse(%v) = %v, expected %v", tt.input, actual, tt.expected)
		}
		if tt.isexpectedError && err == nil {
			t.Errorf("expected error, but error is nil")
		}
	}
}
