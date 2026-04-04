package parser

import (
	"fmt"
	"strconv"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

type ParseContext struct {
	CurrentDuration int
	CurrentOctave   int
	MeasureNumber   int
}

func matchTypes(tokens []lexer.Token, types ...lexer.TokenType) bool {
	if len(tokens) < len(types) {
		return false
	}
	for i, t := range types {
		if tokens[i].Type != t {
			return false
		}
	}
	return true
}

func tokenInt(t lexer.Token) (int, error) {
	v, err := strconv.Atoi(t.Value)
	if err != nil {
		return 0, fmt.Errorf("invalid number token %q: %w", t.Value, err)
	}
	return v, nil
}

func buildNote(step string, duration int, octave int) musicxml.Note {
	return musicxml.Note{
		Pitch:    musicxml.Pitch{Step: step, Octave: octave},
		Duration: duration,
		Type:     "quarter", // TODO: Determine the note type based on the duration and divisions
	}
}

type ParseFunction func(tokens []lexer.Token, ctx *ParseContext) (musicxml.Element, []lexer.Token, ParseContext, error)

func adapt[T musicxml.Element](f func([]lexer.Token, *ParseContext) (T, []lexer.Token, ParseContext, error)) ParseFunction {
	return func(tokens []lexer.Token, ctx *ParseContext) (musicxml.Element, []lexer.Token, ParseContext, error) {
		return f(tokens, ctx)
	}
}
