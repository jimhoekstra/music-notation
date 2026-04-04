package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func MatchesTimeSignature(tokens []lexer.Token) bool {
	return matchTypes(tokens, lexer.TokenTime, lexer.TokenWhitespace, lexer.TokenNumber, lexer.TokenForwardSlash, lexer.TokenNumber)
}

func ParseTimeSignature(tokens []lexer.Token, ctx *ParseContext) (musicxml.Time, []lexer.Token, ParseContext, error) {
	if MatchesTimeSignature(tokens) {
		beats, err := tokenInt(tokens[2])
		if err != nil {
			return musicxml.Time{}, tokens, *ctx, err
		}
		beatType, err := tokenInt(tokens[4])
		if err != nil {
			return musicxml.Time{}, tokens, *ctx, err
		}
		timeSignature := musicxml.Time{
			Beats:    beats,
			BeatType: beatType,
		}
		return timeSignature, tokens[5:], *ctx, nil
	}

	return musicxml.Time{}, tokens, *ctx, errors.New("expected a time signature")
}
