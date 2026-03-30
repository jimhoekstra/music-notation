package parser

import (
	"errors"
	"testing"
)

func findPatternForToken(token TokenType, t *testing.T) (Pattern, error) {
	t.Helper()

	for _, pattern := range Patterns {
		if pattern.Type == token {
			return pattern, nil
		}
	}
	return Pattern{}, errors.New("no pattern found for token type")
}

func testExpectedMatches(t *testing.T, pattern Pattern, expectedMatches []string) {
	t.Helper()

	for _, input := range expectedMatches {
		if !pattern.Regex.MatchString(input) {
			t.Errorf("expected pattern to match '%s', but it did not", input)
		}
	}
}

func testExpectedNonMatches(t *testing.T, pattern Pattern, expectedNonMatches []string) {
	t.Helper()

	for _, input := range expectedNonMatches {
		if pattern.Regex.MatchString(input) {
			t.Errorf("expected pattern to not match '%s', but it did", input)
		}
	}
}

func testPattern(t *testing.T, tokenType TokenType, expectedMatches []string, expectedNonMatches []string) {
	pattern, err := findPatternForToken(tokenType, t)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	testExpectedMatches(t, pattern, expectedMatches)
	testExpectedNonMatches(t, pattern, expectedNonMatches)
}

func TestClefPattern(t *testing.T) {
	testPattern(t, TokenClef, []string{"clef"}, []string{"clefX", "cle", "Xclef"})
}

func TestClefSpecifierPattern(t *testing.T) {
	testPattern(t, TokenClefSpecifier, []string{"treble", "bass"}, []string{"trebleX", "bassX", "Xtreble"})
}

func TestKeyPattern(t *testing.T) {
	testPattern(t, TokenKey, []string{"key"}, []string{"keyX", "ke", "Xkey"})
}

func TestTimePattern(t *testing.T) {
	testPattern(t, TokenTime, []string{"time"}, []string{"timeX", "tim", "Xtime"})
}

func TestForwardSlashPattern(t *testing.T) {
	testPattern(t, TokenForwardSlash, []string{"/"}, []string{"X/"})
}

func TestNotePattern(t *testing.T) {
	testPattern(t, TokenNote, []string{"c", "C", "d#", "Eb"}, []string{"h", "i", "j", "k"})
}

func TestNumberPattern(t *testing.T) {
	testPattern(t, TokenNumber, []string{"1", "4", "16"}, []string{"a", "b", "c"})
}

func TestWhitespacePattern(t *testing.T) {
	testPattern(t, TokenWhitespace, []string{" ", "\t", "\n"}, []string{"a", "1", "/"})
}
