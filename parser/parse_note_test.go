package parser

import (
	"testing"

	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func TestParseNoteDurationAndOctave(t *testing.T) {
	// "4c4" — explicit duration and octave
	tokens := []lexer.Token{
		{Type: lexer.TokenNumber, Value: "4"},
		{Type: lexer.TokenNote, Value: "c"},
		{Type: lexer.TokenNumber, Value: "4"},
	}
	ctx := &ParseContext{CurrentDuration: 1, CurrentOctave: 5}
	note, tokens, newCtx, err := ParseNote(tokens, ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if note.Pitch.Step != "C" || note.Pitch.Octave != 4 || note.Duration != 4 {
		t.Errorf("unexpected note: %+v", note)
	}
	if newCtx.CurrentDuration != 4 || newCtx.CurrentOctave != 4 {
		t.Errorf("expected duration=4 octave=4, got duration=%d octave=%d", newCtx.CurrentDuration, newCtx.CurrentOctave)
	}
	if len(tokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(tokens))
	}
}

func TestParseNoteOctaveOnly(t *testing.T) {
	// "c4" — explicit octave, use current duration
	tokens := []lexer.Token{
		{Type: lexer.TokenNote, Value: "c"},
		{Type: lexer.TokenNumber, Value: "4"},
	}
	ctx := &ParseContext{CurrentDuration: 2, CurrentOctave: 5}
	note, tokens, newCtx, err := ParseNote(tokens, ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if note.Pitch.Step != "C" || note.Duration != 2 || note.Pitch.Octave != 4 {
		t.Errorf("unexpected note: %+v", note)
	}
	if newCtx.CurrentDuration != 2 || newCtx.CurrentOctave != 4 {
		t.Errorf("expected duration=2 octave=4, got duration=%d octave=%d", newCtx.CurrentDuration, newCtx.CurrentOctave)
	}
	if len(tokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(tokens))
	}
}

func TestParseNoteDurationOnly(t *testing.T) {
	// "4c" — explicit duration, use current octave
	tokens := []lexer.Token{
		{Type: lexer.TokenNumber, Value: "4"},
		{Type: lexer.TokenNote, Value: "c"},
	}
	ctx := &ParseContext{CurrentDuration: 1, CurrentOctave: 5}
	note, tokens, newCtx, err := ParseNote(tokens, ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if note.Pitch.Step != "C" || note.Duration != 4 || note.Pitch.Octave != 5 {
		t.Errorf("unexpected note: %+v", note)
	}
	if newCtx.CurrentDuration != 4 || newCtx.CurrentOctave != 5 {
		t.Errorf("expected duration=4 octave=5, got duration=%d octave=%d", newCtx.CurrentDuration, newCtx.CurrentOctave)
	}
	if len(tokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(tokens))
	}
}

func TestParseNotePitchOnly(t *testing.T) {
	// "c" — use current duration and octave
	tokens := []lexer.Token{
		{Type: lexer.TokenNote, Value: "c"},
	}
	ctx := &ParseContext{CurrentDuration: 2, CurrentOctave: 3}
	note, tokens, newCtx, err := ParseNote(tokens, ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if note.Pitch.Step != "C" || note.Duration != 2 || note.Pitch.Octave != 3 {
		t.Errorf("unexpected note: %+v", note)
	}
	if newCtx.CurrentDuration != 2 || newCtx.CurrentOctave != 3 {
		t.Errorf("expected duration=2 octave=3, got duration=%d octave=%d", newCtx.CurrentDuration, newCtx.CurrentOctave)
	}
	if len(tokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(tokens))
	}
}

func TestParseNoteNoMatch(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.TokenWhitespace, Value: " "},
	}
	_, _, _, err := ParseNote(tokens, &ParseContext{CurrentDuration: 1, CurrentOctave: 4})
	if err == nil {
		t.Error("expected error for non-note token, got nil")
	}
}

func TestParseRemainingTokens(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.TokenNumber, Value: "4"},
		{Type: lexer.TokenNote, Value: "c"},
		{Type: lexer.TokenNumber, Value: "4"},
		{Type: lexer.TokenWhitespace, Value: " "},
	}
	note, tokens, _, err := ParseNote(tokens, &ParseContext{CurrentDuration: 1, CurrentOctave: 5})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if note.Pitch.Step != "C" || note.Pitch.Octave != 4 || note.Duration != 4 {
		t.Errorf("unexpected note: %+v", note)
	}
	if len(tokens) != 1 || tokens[0].Type != lexer.TokenWhitespace || tokens[0].Value != " " {
		t.Errorf("expected remaining token to be a single whitespace token, got %v", tokens)
	}
}
