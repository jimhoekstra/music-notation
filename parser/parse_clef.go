package parser

import (
	"errors"
)

func MatchesClef(tokens []Token) bool {
	return matchTypes(tokens, TokenClef, TokenWhitespace, TokenClefSpecifier)
}

func ParseClef(tokens []Token) (string, []Token, error) {
	if MatchesClef(tokens) {
		clefSpecifier := tokens[2].Value
		return clefSpecifier, tokens[3:], nil
	}

	return "", tokens, errors.New("expected a clef token followed by a clef specifier")
}
