package parser

import (
	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

type ParseFunction func(tokens []lexer.Token, ctx *ParseContext) (musicxml.Element, []lexer.Token, ParseContext, error)

func adapt[T musicxml.Element](f func([]lexer.Token, *ParseContext) (T, []lexer.Token, ParseContext, error)) ParseFunction {
	return func(tokens []lexer.Token, ctx *ParseContext) (musicxml.Element, []lexer.Token, ParseContext, error) {
		return f(tokens, ctx)
	}
}
