package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func MatchesKeySignature(tokens []lexer.Token) bool {
	return matchTypes(tokens, lexer.TokenKey, lexer.TokenWhitespace, lexer.TokenNote)
}

func ParseKeySignature(tokens []lexer.Token, ctx *ParseContext) (musicxml.Key, []lexer.Token, ParseContext, error) {
	if MatchesKeySignature(tokens) {
		key := musicxml.Key{
			Fifths: 0, // TODO: Set actual key signature based on the note token
		}
		return key, tokens[3:], *ctx, nil
	}

	return musicxml.Key{}, tokens, *ctx, errors.New("expected a key token followed by a note token")
}
