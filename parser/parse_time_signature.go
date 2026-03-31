package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
)

func MatchesTimeSignature(tokens []Token) bool {
	return matchTypes(tokens, TokenTime, TokenWhitespace, TokenNumber, TokenForwardSlash, TokenNumber)
}

func ParseTimeSignature(tokens []Token, ctx *ParseContext) (musicxml.Time, []Token, ParseContext, error) {
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
