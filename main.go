package main

import (
	"fmt"

	"github.com/jimhoekstra/music-notation/parser"
)

func main() {
	input := "time 4/4 clef treble c d clef bass"
	tokens := parser.Tokenize(input)

	ctx := parser.ParseContext{
		CurrentDuration: 4,
		CurrentOctave:   4,
	}

	attr, _, _, err := parser.ParseAttributes(tokens, &ctx)
	if err != nil {
		fmt.Printf("Error parsing note: %v\n", err)
		return
	}

	fmt.Println(attr.Name())
}
