package parser

import (
	"testing"

	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func TestParseTimeSignature(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.TokenTime, Value: "time"},
		{Type: lexer.TokenOpenParen, Value: "("},
		{Type: lexer.TokenNumber, Value: "3"},
		{Type: lexer.TokenForwardSlash, Value: "/"},
		{Type: lexer.TokenNumber, Value: "4"},
		{Type: lexer.TokenCloseParen, Value: ")"},
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
