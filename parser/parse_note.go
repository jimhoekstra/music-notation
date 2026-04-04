package parser

import (
	"errors"
	"fmt"
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
		note, err := buildNote(strings.ToUpper(tokens[1].Value[0:1]), duration, octave, ctx.Division)
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return note, tokens[3:], ParseContext{CurrentDuration: duration, CurrentOctave: octave, Division: ctx.Division}, nil
	}

	// Check for a note with an explicit octave, e.g. "c4"
	if matchTypes(tokens, lexer.TokenNote, lexer.TokenNumber) {
		octave, err := tokenInt(tokens[1])
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		note, err := buildNote(strings.ToUpper(tokens[0].Value[0:1]), ctx.CurrentDuration, octave, ctx.Division)
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return note, tokens[2:], ParseContext{CurrentDuration: ctx.CurrentDuration, CurrentOctave: octave, Division: ctx.Division}, nil
	}

	// Check for a note with an explicit duration, e.g. "4c"
	if matchTypes(tokens, lexer.TokenNumber, lexer.TokenNote) {
		duration, err := tokenInt(tokens[0])
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		note, err := buildNote(strings.ToUpper(tokens[1].Value[0:1]), duration, ctx.CurrentOctave, ctx.Division)
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return note, tokens[2:], ParseContext{CurrentDuration: duration, CurrentOctave: ctx.CurrentOctave, Division: ctx.Division}, nil
	}

	// Check for a note with only a pitch, e.g. "c"
	if matchTypes(tokens, lexer.TokenNote) {
		note, err := buildNote(strings.ToUpper(tokens[0].Value[0:1]), ctx.CurrentDuration, ctx.CurrentOctave, ctx.Division)
		if err != nil {
			return musicxml.Note{}, tokens, *ctx, err
		}
		return note, tokens[1:], *ctx, nil
	}

	return musicxml.Note{}, tokens, *ctx, errors.New("expected a note token")
}

func noteType(duration int, division int) (string, error) {
	if division <= 0 {
		return "", errors.New("division must be greater than zero")
	}
	switch {
	case duration == division*4:
		return "whole", nil
	case duration == division*2:
		return "half", nil
	case duration == division:
		return "quarter", nil
	case duration*2 == division:
		return "eighth", nil
	case duration*4 == division:
		return "16th", nil
	case duration*8 == division:
		return "32nd", nil
	default:
		return "", fmt.Errorf("no matching note type for duration %d with division %d", duration, division)
	}
}

func buildNote(step string, duration int, octave int, division int) (musicxml.Note, error) {
	xmlDuration := division * 4 / duration
	t, err := noteType(xmlDuration, division)
	if err != nil {
		return musicxml.Note{}, err
	}
	return musicxml.Note{
		Chord:    nil,
		Pitch:    musicxml.Pitch{Step: step, Octave: octave},
		Duration: xmlDuration,
		Type:     t,
	}, nil
}
