package parser

import (
	"testing"

	"github.com/jimhoekstra/music-notation/musicxml"
)

func TestMatchTypesSingle(t *testing.T) {
	tokens := []Token{
		{Type: TokenNote, Value: "c"},
	}
	match := matchTypes(tokens, TokenNote)
	if !match {
		t.Errorf("Expected match to be true, got false")
	}
}

func TestMatchTypesSingleNoMatch(t *testing.T) {
	tokens := []Token{
		{Type: TokenNote, Value: "c"},
	}
	match := matchTypes(tokens, TokenWhitespace)
	if match {
		t.Errorf("Expected match to be false, got true")
	}
}

func TestMatchTypesMultiple(t *testing.T) {
	tokens := []Token{
		{Type: TokenNote, Value: "c"},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenWhitespace, Value: "  "},
	}
	match := matchTypes(tokens, TokenNote, TokenNumber, TokenWhitespace)
	if !match {
		t.Errorf("Expected match to be true, got false")
	}
}

func TestMatchTypesMultipleNoMatch(t *testing.T) {
	tokens := []Token{
		{Type: TokenNote, Value: "c"},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenWhitespace, Value: "  "},
	}
	match := matchTypes(tokens, TokenNote, TokenWhitespace, TokenNumber)
	if match {
		t.Errorf("Expected match to be false, got true")
	}
}

func TestMatchTypesTooShort(t *testing.T) {
	tokens := []Token{
		{Type: TokenNote, Value: "c"},
	}
	match := matchTypes(tokens, TokenNote, TokenNumber)
	if match {
		t.Errorf("Expected match to be false, got true")
	}
}

func TestMatchTypesEmpty(t *testing.T) {
	tokens := []Token{}
	match := matchTypes(tokens, TokenNote)
	if match {
		t.Errorf("Expected match to be false, got true")
	}
}

func TestTokenIntValid(t *testing.T) {
	token := Token{Type: TokenNumber, Value: "42"}
	v, err := tokenInt(token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 42 {
		t.Errorf("expected 42, got %d", v)
	}
}

func TestTokenIntInvalid(t *testing.T) {
	token := Token{Type: TokenNumber, Value: "abc"}
	_, err := tokenInt(token)
	if err == nil {
		t.Error("expected error for non-numeric token, got nil")
	}
}

func TestTokenIntZero(t *testing.T) {
	token := Token{Type: TokenNumber, Value: "0"}
	v, err := tokenInt(token)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v != 0 {
		t.Errorf("expected 0, got %d", v)
	}
}

func TestBuildNote(t *testing.T) {
	note := buildNote("C", 4, 4)
	expected := musicxml.Note{
		Pitch:    musicxml.Pitch{Step: "C", Octave: 4},
		Duration: 4,
		Type:     "quarter",
	}
	if note != expected {
		t.Errorf("expected %+v, got %+v", expected, note)
	}
}
