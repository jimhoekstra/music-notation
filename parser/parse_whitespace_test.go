package parser

import (
	"testing"
)

func TestParseWhiteSpace(t *testing.T) {
	tokens := []Token{
		{Type: TokenWhitespace, Value: " "},
		{Type: TokenNote, Value: "c"},
	}
	parseContext := ParseContext{}
	_, remainingTokens, _, err := ParseWhiteSpace(tokens, &parseContext)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(remainingTokens) != 1 {
		t.Errorf("expected 1 remaining token, got %d", len(remainingTokens))
	}
}
