package parser

import (
	"github.com/jimhoekstra/music-notation/musicxml"
)

func ParseElements(tokens []Token, ctx *ParseContext, parsers []ParseFunction) (
	[]musicxml.Element, []Token, ParseContext, error) {
	var elements []musicxml.Element

	for len(tokens) > 0 {
		matched := false

		for _, parser := range parsers {
			element, remainingTokens, newCtx, err := parser(tokens, ctx)
			if err == nil {
				if element.Name() != "EmptyElement" {
					elements = append(elements, element)
				}
				ctx = &newCtx
				tokens = remainingTokens
				matched = true
				break
			}
		}

		// If none of the parser functions matched, break the loop and return
		if !matched {
			break
		}
	}

	return elements, tokens, *ctx, nil
}

func ParseAttributes(tokens []Token, ctx *ParseContext) (
	musicxml.Attributes, []Token, ParseContext, error) {
	parsers := []ParseFunction{
		adapt(ParseClef),
		adapt(ParseKeySignature),
		adapt(ParseTimeSignature),
		adapt(ParseWhiteSpace),
	}

	elements, remainingTokens, newCtx, err := ParseElements(tokens, ctx, parsers)
	if err != nil {
		return musicxml.Attributes{}, tokens, *ctx, err
	}

	var clef *musicxml.Clef
	var key *musicxml.Key
	var time *musicxml.Time

	for _, element := range elements {
		switch e := element.(type) {
		case musicxml.Clef:
			clef = &e
		case musicxml.Key:
			key = &e
		case musicxml.Time:
			time = &e
		}
	}

	return musicxml.Attributes{
		Clef: clef,
		Key:  key,
		Time: time,
	}, remainingTokens, newCtx, nil
}
