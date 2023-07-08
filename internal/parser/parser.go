package parser

import (
	"fmt"
	"strconv"
)

type Node interface{}

type Number struct {
	Value float64
}

type Operator struct {
	Op    string
	Left  Node
	Right Node
}

func Parse(tokens []Token) (Node, []Token, error) {
	token, restTokens := tokens[0], tokens[1:]

	switch token.Type {
	case "number":
		value, _ := strconv.ParseFloat(token.Value, 64)
		return Number{Value: value}, restTokens, nil
	case "operator":
		return Operator{Op: token.Value}, restTokens, nil
	case "left_paren":
		operator, restTokens := restTokens[0], restTokens[1:]
		left, restTokens, err := Parse(restTokens)
		if err != nil {
			return nil, nil, err
		}
		right, restTokens, err := Parse(restTokens)
		if err != nil {
			return nil, nil, err
		}
		if restTokens[0].Type != "right_paren" {
			return nil, nil, fmt.Errorf("expected right_paren, got %s", restTokens[0].Type)
		}
		return Operator{Op: operator.Value, Left: left, Right: right}, restTokens[1:], nil
	case "right_paren":
		return nil, nil, fmt.Errorf("unexpected token %s", token.Type)
	default:
		return nil, nil, fmt.Errorf("unexpected token %s", token.Type)
	}
}
