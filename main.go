package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/jimhoekstra/music-notation/parser"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func main() {
	input := "time 4/4 clef treble key c c d e f / g a b c5 / b4 a g f / e d clef bass 2c /"
	tokens := lexer.Tokenize(input)

	ctx := parser.ParseContext{
		CurrentDuration: 4,
		CurrentOctave:   4,
		MeasureNumber:   0,
	}

	part, _, _, err := parser.ParsePart(tokens, &ctx)
	if err != nil {
		fmt.Printf("Error parsing part: %v\n", err)
		return
	}

	fmt.Println(part.Name())
	output, err := xml.MarshalIndent(part, "", "  ")
	if err != nil {
		panic(err)
	}

	os.WriteFile("result.xml", output, 0644)
}
