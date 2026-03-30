package parser

import (
	"testing"
)

func TestParseClef(t *testing.T) {
	tokens := []Token{
		{Type: TokenClef, Value: "clef"},
		{Type: TokenWhitespace, Value: " "},
		{Type: TokenClefSpecifier, Value: "treble"},
	}
	clefSpecifier, remainingTokens, err := ParseClef(tokens)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if clefSpecifier != "treble" {
		t.Errorf("expected clef specifier 'treble', got '%s'", clefSpecifier)
	}
	if len(remainingTokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(remainingTokens))
	}
}
