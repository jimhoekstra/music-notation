package parser

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

// noteToFifths converts a tonic note name (e.g. "c", "g#", "bb") to its
// number of fifths in the key signature (negative = flats, positive = sharps).
func noteToFifths(note string) (int, error) {
	fifthsMap := map[string]int{
		"c":  0,
		"g":  1,
		"d":  2,
		"a":  3,
		"e":  4,
		"b":  5,
		"f#": 6,
		"c#": 7,
		"f":  -1,
		"bb": -2,
		"eb": -3,
		"ab": -4,
		"db": -5,
		"gb": -6,
		"cb": -7,
	}
	fifths, ok := fifthsMap[strings.ToLower(note)]
	if !ok {
		return 0, fmt.Errorf("unknown key signature tonic: %q", note)
	}
	return fifths, nil
}

func MatchesKeySignature(tokens []lexer.Token) bool {
	return matchTypes(tokens, lexer.TokenKey, lexer.TokenOpenParen, lexer.TokenNote, lexer.TokenCloseParen)
}

func ParseKeySignature(tokens []lexer.Token, ctx *ParseContext) (musicxml.Key, []lexer.Token, ParseContext, error) {
	if MatchesKeySignature(tokens) {
		fifths, err := noteToFifths(tokens[2].Value)
		if err != nil {
			return musicxml.Key{}, tokens, *ctx, err
		}
		key := musicxml.Key{
			Fifths: fifths,
		}
		return key, tokens[4:], *ctx, nil
	}

	return musicxml.Key{}, tokens, *ctx, errors.New("expected a key token followed by a note token")
}
