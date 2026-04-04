package lexer

import "regexp"

type TokenType int

const (
	TokenNote          TokenType = iota
	TokenNumber        TokenType = iota
	TokenWhitespace    TokenType = iota
	TokenClef          TokenType = iota
	TokenClefSpecifier TokenType = iota
	TokenKey           TokenType = iota
	TokenTime          TokenType = iota
	TokenForwardSlash  TokenType = iota
	TokenVerticalBar   TokenType = iota
	TokenOpenParen     TokenType = iota
	TokenCloseParen    TokenType = iota
)

func (t TokenType) String() string {
	switch t {
	case TokenNote:
		return "TokenNote"
	case TokenNumber:
		return "TokenNumber"
	case TokenWhitespace:
		return "TokenWhitespace"
	case TokenClef:
		return "TokenClef"
	case TokenClefSpecifier:
		return "TokenClefSpecifier"
	case TokenKey:
		return "TokenKey"
	case TokenTime:
		return "TokenTime"
	case TokenForwardSlash:
		return "TokenForwardSlash"
	case TokenVerticalBar:
		return "TokenVerticalBar"
	case TokenOpenParen:
		return "TokenOpenParen"
	case TokenCloseParen:
		return "TokenCloseParen"
	default:
		return "TokenUnknown"
	}
}

type Pattern struct {
	Type  TokenType
	Regex regexp.Regexp
}

var Patterns = []Pattern{
	{Type: TokenClef, Regex: *regexp.MustCompile(`^clef\b`)},
	{Type: TokenClefSpecifier, Regex: *regexp.MustCompile(`^(treble|bass)\b`)},
	{Type: TokenKey, Regex: *regexp.MustCompile(`^key\b`)},
	{Type: TokenTime, Regex: *regexp.MustCompile(`^time\b`)},
	{Type: TokenForwardSlash, Regex: *regexp.MustCompile(`^/`)},
	{Type: TokenNote, Regex: *regexp.MustCompile(`^[a-gA-G][#b]?`)},
	{Type: TokenNumber, Regex: *regexp.MustCompile(`^\d+`)},
	{Type: TokenWhitespace, Regex: *regexp.MustCompile(`^\s+`)},
	{Type: TokenVerticalBar, Regex: *regexp.MustCompile(`^\|`)},
	{Type: TokenOpenParen, Regex: *regexp.MustCompile(`^\(`)},
	{Type: TokenCloseParen, Regex: *regexp.MustCompile(`^\)`)},
}

type Token struct {
	Type  TokenType
	Value string
}
