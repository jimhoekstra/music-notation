package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func MatchesClef(tokens []lexer.Token) bool {
	return matchTypes(tokens, lexer.TokenClef, lexer.TokenOpenParen, lexer.TokenClefSpecifier, lexer.TokenCloseParen)
}

func getClefSign(specifier string) (musicxml.ClefSign, error) {
	switch specifier {
	case "treble":
		return musicxml.TrebleClef, nil
	case "bass":
		return musicxml.BassClef, nil
	default:
		return "", errors.New("unknown clef specifier: " + specifier)
	}
}

func getClefLine(clefSign musicxml.ClefSign) (int, error) {
	switch clefSign {
	case musicxml.TrebleClef:
		return 2, nil
	case musicxml.BassClef:
		return 4, nil
	default:
		return 0, errors.New("unknown clef sign: " + string(clefSign))
	}
}

func ParseClef(tokens []lexer.Token, ctx *ParseContext) (musicxml.Clef, []lexer.Token, ParseContext, error) {

	if !MatchesClef(tokens) {
		return musicxml.Clef{}, tokens, *ctx, errors.New("expected a clef token followed by a clef specifier")
	}

	clefSign, err := getClefSign(tokens[2].Value)
	if err != nil {
		return musicxml.Clef{}, tokens, *ctx, err
	}

	clefLine, err := getClefLine(clefSign)
	if err != nil {
		return musicxml.Clef{}, tokens, *ctx, err
	}

	clef := musicxml.Clef{
		Sign: clefSign,
		Line: clefLine,
	}
	return clef, tokens[4:], *ctx, nil
}
