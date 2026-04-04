package parser

import (
	"testing"

	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func TestParseWhiteSpace(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.TokenWhitespace, Value: " "},
		{Type: lexer.TokenNote, Value: "c"},
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
