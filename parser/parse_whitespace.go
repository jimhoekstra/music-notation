package parser

import (
	"errors"
)

func MatchesWhiteSpace(tokens []Token) bool {
	return matchTypes(tokens, TokenWhitespace)
}

func ParseWhiteSpace(tokens []Token) ([]Token, error) {
	if matchTypes(tokens, TokenWhitespace) {
		return tokens[1:], nil
	}

	return tokens, errors.New("expected a whitespace token")
}
