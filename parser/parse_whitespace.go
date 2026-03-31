package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
)

func MatchesWhiteSpace(tokens []Token) bool {
	return matchTypes(tokens, TokenWhitespace)
}

func ParseWhiteSpace(tokens []Token, ctx *ParseContext) (musicxml.Element, []Token, ParseContext, error) {
	if MatchesWhiteSpace(tokens) {
		return musicxml.EmptyElement{}, tokens[1:], *ctx, nil
	}

	return musicxml.EmptyElement{}, tokens, *ctx, errors.New("expected a whitespace token")
}
