package parser

import (
	"errors"

	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func ParseElements(tokens []lexer.Token, ctx *ParseContext, parsers []ParseFunction, stopParser ParseFunction) (
	[]musicxml.Element, []lexer.Token, ParseContext, error) {
	var elements []musicxml.Element

	for len(tokens) > 0 {
		matched := false

		if stopParser != nil {
			element, remainingTokens, newCtx, err := stopParser(tokens, ctx)
			if err == nil {
				elements = append(elements, element)
				return elements, remainingTokens, newCtx, nil
			}
		}

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

func ParseAttributes(tokens []lexer.Token, ctx *ParseContext) (
	musicxml.Attributes, []lexer.Token, ParseContext, error) {
	parsers := []ParseFunction{
		adapt(ParseClef),
		adapt(ParseKeySignature),
		adapt(ParseTimeSignature),
		adapt(ParseWhiteSpace),
	}

	elements, remainingTokens, newCtx, err := ParseElements(tokens, ctx, parsers, adapt(NeverMatch))
	if err != nil {
		return musicxml.Attributes{}, tokens, *ctx, err
	}
	if len(elements) == 0 {
		return musicxml.Attributes{}, tokens, *ctx, errors.New("no attributes found")
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

func NeverMatch(tokens []lexer.Token, ctx *ParseContext) (
	musicxml.EmptyElement, []lexer.Token, ParseContext, error) {
	return musicxml.EmptyElement{}, tokens, *ctx, errors.New("this parser never matches")
}

func MatchesBarline(tokens []lexer.Token) bool {
	return matchTypes(tokens, lexer.TokenForwardSlash)
}

func ParseBarline(tokens []lexer.Token, ctx *ParseContext) (
	musicxml.Barline, []lexer.Token, ParseContext, error) {

	if MatchesBarline(tokens) {
		return musicxml.Barline{
			Location: "right",
			BarStyle: "regular",
		}, tokens[1:], *ctx, nil
	} else {
		return musicxml.Barline{}, tokens, *ctx, errors.New("expected a barline token")
	}
}

func ParseMeasure(tokens []lexer.Token, ctx *ParseContext) (
	musicxml.Measure, []lexer.Token, ParseContext, error) {
	parsers := []ParseFunction{
		adapt(ParseAttributes),
		adapt(ParseNote),
		adapt(ParseWhiteSpace),
	}

	elements, remainingTokens, newCtx, err := ParseElements(tokens, ctx, parsers, adapt(ParseBarline))
	if err != nil {
		return musicxml.Measure{}, tokens, *ctx, err
	}
	if len(elements) == 0 {
		return musicxml.Measure{}, tokens, *ctx, errors.New("no measure elements found")
	}

	var measureElements []musicxml.MeasureElement
	for _, element := range elements {
		switch e := element.(type) {
		case musicxml.Attributes:
			measureElements = append(measureElements, musicxml.MeasureElement{Attributes: &e})
		case musicxml.Note:
			measureElements = append(measureElements, musicxml.MeasureElement{Note: &e})
		case musicxml.Barline:
			measureElements = append(measureElements, musicxml.MeasureElement{Barline: &e})
		}
	}

	newCtx.MeasureNumber = ctx.MeasureNumber + 1

	return musicxml.Measure{
		Number:   newCtx.MeasureNumber,
		Elements: measureElements,
	}, remainingTokens, newCtx, nil

}
