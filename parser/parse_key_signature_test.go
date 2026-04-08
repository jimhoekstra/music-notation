package parser

import (
	"testing"

	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func TestParseKeySignature(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.TokenKey, Value: "key"},
		{Type: lexer.TokenOpenParen, Value: "("},
		{Type: lexer.TokenNote, Value: "c"},
		{Type: lexer.TokenCloseParen, Value: ")"},
	}
	parseContext := ParseContext{}
	key, remainingTokens, _, err := ParseKeySignature(tokens, &parseContext)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key.Fifths != 0 {
		t.Errorf("expected fifths 0, got '%d'", key.Fifths)
	}
	if len(remainingTokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(remainingTokens))
	}
}

func TestNoteToFifthsUnknown(t *testing.T) {
	_, err := noteToFifths("x")
	if err == nil {
		t.Error("expected an error for unknown note, got nil")
	}
}

func TestParseKeySignatureG(t *testing.T) {
	testKeySignature(t, "g", 1)
}

func TestParseKeySignatureD(t *testing.T) {
	testKeySignature(t, "d", 2)
}

func TestParseKeySignatureFSharp(t *testing.T) {
	testKeySignature(t, "f#", 6)
}

func TestParseKeySignatureF(t *testing.T) {
	testKeySignature(t, "f", -1)
}

func TestParseKeySignatureBB(t *testing.T) {
	testKeySignature(t, "bb", -2)
}

func TestParseKeySignatureEB(t *testing.T) {
	testKeySignature(t, "eb", -3)
}

func testKeySignature(t *testing.T, note string, expectedFifths int) {
	t.Helper()

	tokens := []lexer.Token{
		{Type: lexer.TokenKey, Value: "key"},
		{Type: lexer.TokenOpenParen, Value: "("},
		{Type: lexer.TokenNote, Value: note},
		{Type: lexer.TokenCloseParen, Value: ")"},
	}
	parseContext := ParseContext{}
	key, _, _, err := ParseKeySignature(tokens, &parseContext)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if key.Fifths != expectedFifths {
		t.Errorf("ParseKeySignature(%q): expected fifths %d, got %d", note, expectedFifths, key.Fifths)
	}
}
