package parser

import (
	"github.com/jimhoekstra/music-notation/musicxml"
)

type ParseFunction func(tokens []Token, ctx *ParseContext) (musicxml.Element, []Token, ParseContext, error)

func adapt[T musicxml.Element](f func([]Token, *ParseContext) (T, []Token, ParseContext, error)) ParseFunction {
	return func(tokens []Token, ctx *ParseContext) (musicxml.Element, []Token, ParseContext, error) {
		return f(tokens, ctx)
	}
}

var parseFunctions = []ParseFunction{
	adapt(ParseKeySignature),
	adapt(ParseNote),
	adapt(ParseTimeSignature),
	adapt(ParseWhiteSpace),
}
