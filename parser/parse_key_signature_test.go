package parser

import (
	"testing"

	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func TestParseKeySignature(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.TokenKey, Value: "key"},
		{Type: lexer.TokenWhitespace, Value: " "},
		{Type: lexer.TokenNote, Value: "c"},
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
