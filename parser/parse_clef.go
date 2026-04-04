package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func MatchesClef(tokens []lexer.Token) bool {
	return matchTypes(tokens, lexer.TokenClef, lexer.TokenWhitespace, lexer.TokenClefSpecifier)
}

func ParseClef(tokens []lexer.Token, ctx *ParseContext) (musicxml.Clef, []lexer.Token, ParseContext, error) {
	if MatchesClef(tokens) {
		clef := musicxml.Clef{
			Sign: musicxml.ClefSign(tokens[2].Value),
			Line: 2, // Default line for treble clef; this could be made more flexible if needed
		}
		return clef, tokens[3:], *ctx, nil
	}

	return musicxml.Clef{}, tokens, *ctx, errors.New("expected a clef token followed by a clef specifier")
}
