package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/jimhoekstra/music-notation/io"
	"github.com/jimhoekstra/music-notation/parser"
	"github.com/jimhoekstra/music-notation/parser/lexer"
)

func main() {
	input, err := io.LoadInputFile("input.txt")
	if err != nil {
		fmt.Printf("Error loading input file: %v\n", err)
		return
	}
	tokens := lexer.Tokenize(input)

	ctx := parser.ParseContext{
		CurrentDuration: 4,
		CurrentOctave:   4,
		Division:        4,
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
