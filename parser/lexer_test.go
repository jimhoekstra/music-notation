package parser

import (
	"testing"
)

func assertTokensEqual(t *testing.T, input string, expected []Token) {
	t.Helper()
	actual := Tokenize(input)

	if len(actual) != len(expected) {
		t.Errorf("Expected %d tokens, got %d", len(expected), len(actual))
		return
	}

	for i, token := range actual {
		if token != expected[i] {
			t.Errorf("Expected token %d to be %v, got %v", i, expected[i], token)
		}
	}
}

func TestTokenizeNote(t *testing.T) {
	assertTokensEqual(t, "a", []Token{
		{Type: TokenNote, Value: "a"},
	})
}

func TestTokenizeWhitespace(t *testing.T) {
	assertTokensEqual(t, "  \n", []Token{
		{Type: TokenWhitespace, Value: "  \n"},
	})
}

func TestTokenizeNumber(t *testing.T) {
	assertTokensEqual(t, "42", []Token{
		{Type: TokenNumber, Value: "42"},
	})
}

func TestTokenizeClef(t *testing.T) {
	assertTokensEqual(t, "clef", []Token{
		{Type: TokenClef, Value: "clef"},
	})
}

func TestTokenizeClefSpecifier(t *testing.T) {
	assertTokensEqual(t, "treble", []Token{
		{Type: TokenClefSpecifier, Value: "treble"},
	})
	assertTokensEqual(t, "bass", []Token{
		{Type: TokenClefSpecifier, Value: "bass"},
	})
}

func TestTokenize(t *testing.T) {
	assertTokensEqual(t, "a b# 4c", []Token{
		{Type: TokenNote, Value: "a"},
		{Type: TokenWhitespace, Value: " "},
		{Type: TokenNote, Value: "b#"},
		{Type: TokenWhitespace, Value: " "},
		{Type: TokenNumber, Value: "4"},
		{Type: TokenNote, Value: "c"},
	})
}
