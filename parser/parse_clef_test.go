package parser

import (
	"testing"

	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func TestParseClef(t *testing.T) {
	tokens := []lexer.Token{
		{Type: lexer.TokenClef, Value: "clef"},
		{Type: lexer.TokenOpenParen, Value: "("},
		{Type: lexer.TokenClefSpecifier, Value: "treble"},
		{Type: lexer.TokenCloseParen, Value: ")"},
	}
	parseContext := ParseContext{}
	clef, remainingTokens, _, err := ParseClef(tokens, &parseContext)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if clef.Sign != "G" {
		t.Errorf("expected clef specifier 'G', got '%s'", clef.Sign)
	}
	if len(remainingTokens) != 0 {
		t.Errorf("expected no remaining tokens, got %d", len(remainingTokens))
	}
}
