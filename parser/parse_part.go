package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
)

func ParsePart(tokens []Token, ctx *ParseContext) (
	musicxml.Part, []Token, ParseContext, error) {
	parsers := []ParseFunction{
		adapt(ParseMeasure),
	}

	elements, remainingTokens, newCtx, err := ParseElements(tokens, ctx, parsers, adapt(NeverMatch))
	if err != nil {
		return musicxml.Part{}, tokens, *ctx, err
	}
	if len(elements) == 0 {
		return musicxml.Part{}, tokens, *ctx, errors.New("no measure elements found")
	}

	var measures []musicxml.Measure
	for _, element := range elements {
		switch e := element.(type) {
		case musicxml.Measure:
			measures = append(measures, e)
		}
	}

	part := musicxml.Part{
		ID:       "P1",
		Measures: measures,
	}

	return part, remainingTokens, newCtx, nil
}
