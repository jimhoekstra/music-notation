package parser

import (
	"github.com/jimhoekstra/music-notation/musicxml"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func ParseUserInput(input string) (musicxml.Measure, error) {
	tokens := lexer.Tokenize(input)

	ctx := ParseContext{
		CurrentDuration: 4,
		CurrentOctave:   4,
		Division:        4,
		MeasureNumber:   0,
	}

	measure, _, _, err := ParseMeasure(tokens, &ctx)
	if err != nil {
		return musicxml.Measure{}, err
	}

	return measure, nil
}
