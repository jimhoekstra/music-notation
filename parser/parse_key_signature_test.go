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
