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
	parseContext := ParseContext{}
	clef, remainingTokens, _, err := ParseClef(tokens, &parseContext)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if clef.Sign != "treble" {
		t.Errorf("expected clef specifier 'treble', got '%s'", clef.Sign)
	}
	if len(remainingTokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(remainingTokens))
	}
}
