package parser

import (
	"errors"
	"strings"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func MatchesNote(tokens []lexer.Token) bool {
	return matchTypes(tokens, lexer.TokenNumber, lexer.TokenNote, lexer.TokenNumber) ||
		matchTypes(tokens, lexer.TokenNote, lexer.TokenNumber) ||
		matchTypes(tokens, lexer.TokenNumber, lexer.TokenNote) ||
		matchTypes(tokens, lexer.TokenNote)
}

func ParseNote(tokens []lexer.Token, ctx *ParseContext) (
	musicxml.Note, []lexer.Token, ParseContext, error) {
	// Check for a note with an explicit duration and octave, e.g. "4c4"
	if matchTypes(tokens, lexer.TokenNumber, lexer.TokenNote, lexer.TokenNumber) {
		duration, err := tokenInt(tokens[0])
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		octave, err := tokenInt(tokens[2])
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return buildNote(strings.ToUpper(tokens[1].Value[0:1]), duration, octave),
			tokens[3:], ParseContext{CurrentDuration: duration, CurrentOctave: octave}, nil
	}

	// Check for a note with an explicit octave, e.g. "c4"
	if matchTypes(tokens, lexer.TokenNote, lexer.TokenNumber) {
		octave, err := tokenInt(tokens[1])
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return buildNote(strings.ToUpper(tokens[0].Value[0:1]), ctx.CurrentDuration, octave),
			tokens[2:], ParseContext{CurrentDuration: ctx.CurrentDuration, CurrentOctave: octave}, nil
	}

	// Check for a note with an explicit duration, e.g. "4c"
	if matchTypes(tokens, lexer.TokenNumber, lexer.TokenNote) {
		duration, err := tokenInt(tokens[0])
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return buildNote(strings.ToUpper(tokens[1].Value[0:1]), duration, ctx.CurrentOctave),
			tokens[2:], ParseContext{CurrentDuration: duration, CurrentOctave: ctx.CurrentOctave}, nil
	}

	// Check for a note with only a pitch, e.g. "c"
	if matchTypes(tokens, lexer.TokenNote) {
		return buildNote(strings.ToUpper(tokens[0].Value[0:1]),
			ctx.CurrentDuration, ctx.CurrentOctave), tokens[1:], *ctx, nil
	}

	return musicxml.Note{}, tokens, *ctx, errors.New("expected a note token")
}
