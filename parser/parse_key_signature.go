package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
)

func MatchesKeySignature(tokens []Token) bool {
	return matchTypes(tokens, TokenKey, TokenWhitespace, TokenNote)
}

func ParseKeySignature(tokens []Token, ctx *ParseContext) (musicxml.Key, []Token, ParseContext, error) {
	if MatchesKeySignature(tokens) {
		key := musicxml.Key{
			Fifths: 0, // TODO: Set actual key signature based on the note token
		}
		return key, tokens[3:], *ctx, nil
	}

	return musicxml.Key{}, tokens, *ctx, errors.New("expected a key token followed by a note token")
}
