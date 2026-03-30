package main

import (
	"github.com/jimhoekstra/music-notation/io"
	"github.com/jimhoekstra/music-notation/musicxml"
)

func main() {
	// input, err := io.LoadInputFile("input.txt")
	// if err != nil {
	// 	fmt.Println("Error loading input file:", err)
	// 	return
	// }

	// tokens := parser.Tokenize(input)
	// for _, token := range tokens {
	// 	fmt.Println(token)
	// }

	hello_world := musicxml.BuildHelloWorld()
	io.SaveScoreAsMusicXml(hello_world, "result.xml")
}
