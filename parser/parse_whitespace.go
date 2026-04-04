package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func MatchesWhiteSpace(tokens []lexer.Token) bool {
	return matchTypes(tokens, lexer.TokenWhitespace)
}

func ParseWhiteSpace(tokens []lexer.Token, ctx *ParseContext) (musicxml.Element, []lexer.Token, ParseContext, error) {
	if MatchesWhiteSpace(tokens) {
		return musicxml.EmptyElement{}, tokens[1:], *ctx, nil
	}

	return musicxml.EmptyElement{}, tokens, *ctx, errors.New("expected a whitespace token")
}
