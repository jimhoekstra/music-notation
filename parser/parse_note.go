package parser

import (
	"errors"
	"strings"

	"github.com/jimhoekstra/music-notation/musicxml"
)

func MatchesNote(tokens []Token) bool {
	return matchTypes(tokens, TokenNumber, TokenNote, TokenNumber) ||
		matchTypes(tokens, TokenNote, TokenNumber) ||
		matchTypes(tokens, TokenNumber, TokenNote) ||
		matchTypes(tokens, TokenNote)
}

func ParseNote(tokens []Token, ctx *ParseContext) (
	musicxml.Note, []Token, ParseContext, error) {
	// Check for a note with an explicit duration and octave, e.g. "4c4"
	if matchTypes(tokens, TokenNumber, TokenNote, TokenNumber) {
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
	if matchTypes(tokens, TokenNote, TokenNumber) {
		octave, err := tokenInt(tokens[1])
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return buildNote(strings.ToUpper(tokens[0].Value[0:1]), ctx.CurrentDuration, octave),
			tokens[2:], ParseContext{CurrentDuration: ctx.CurrentDuration, CurrentOctave: octave}, nil
	}

	// Check for a note with an explicit duration, e.g. "4c"
	if matchTypes(tokens, TokenNumber, TokenNote) {
		duration, err := tokenInt(tokens[0])
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return buildNote(strings.ToUpper(tokens[1].Value[0:1]), duration, ctx.CurrentOctave),
			tokens[2:], ParseContext{CurrentDuration: duration, CurrentOctave: ctx.CurrentOctave}, nil
	}

	// Check for a note with only a pitch, e.g. "c"
	if matchTypes(tokens, TokenNote) {
		return buildNote(strings.ToUpper(tokens[0].Value[0:1]),
			ctx.CurrentDuration, ctx.CurrentOctave), tokens[1:], *ctx, nil
	}

	return musicxml.Note{}, tokens, *ctx, errors.New("expected a note token")
}
