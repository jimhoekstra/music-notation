package parser

import (
	"testing"
)

func TestParseTimeSignature(t *testing.T) {
	tokens := []Token{
		{Type: TokenTime, Value: "time"},
		{Type: TokenWhitespace, Value: " "},
		{Type: TokenNumber, Value: "3"},
		{Type: TokenForwardSlash, Value: "/"},
		{Type: TokenNumber, Value: "4"},
	}
	parseContext := ParseContext{}
	time, remainingTokens, _, err := ParseTimeSignature(tokens, &parseContext)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if time.Beats != 3 {
		t.Errorf("expected beats 3, got '%d'", time.Beats)
	}
	if time.BeatType != 4 {
		t.Errorf("expected beat-type 4, got '%d'", time.BeatType)
	}
	if len(remainingTokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(remainingTokens))
	}
}
