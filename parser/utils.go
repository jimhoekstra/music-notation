package parser

import (
	"fmt"
	"strconv"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

type ParseContext struct {
	CurrentDuration int
	CurrentOctave   int
	Division        int
	MeasureNumber   int
}

func matchTypes(tokens []lexer.Token, types ...lexer.TokenType) bool {
	if len(tokens) < len(types) {
		return false
	}
	for i, t := range types {
		if tokens[i].Type != t {
			return false
		}
	}
	return true
}

func tokenInt(t lexer.Token) (int, error) {
	v, err := strconv.Atoi(t.Value)
	if err != nil {
		return 0, fmt.Errorf("invalid number token %q: %w", t.Value, err)
	}
	return v, nil
}

type ParseFunction func(tokens []lexer.Token, ctx *ParseContext) (musicxml.Element, []lexer.Token, ParseContext, error)

func adapt[T musicxml.Element](f func([]lexer.Token, *ParseContext) (T, []lexer.Token, ParseContext, error)) ParseFunction {
	return func(tokens []lexer.Token, ctx *ParseContext) (musicxml.Element, []lexer.Token, ParseContext, error) {
		return f(tokens, ctx)
	}
}
