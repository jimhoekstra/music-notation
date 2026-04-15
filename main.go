package main

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/jimhoekstra/music-notation/io"
	"github.com/jimhoekstra/music-notation/parser"
	"github.com/jimhoekstra/music-notation/svg"
)

func main() {
	input, err := io.LoadInputFile("input.txt")
	if err != nil {
		fmt.Printf("Error loading input file: %v\n", err)
		return
	}

	measure, err := parser.ParseUserInput(input)
	fmt.Printf("Parsed measure:\n%s\n", measure.Name())

	if err != nil {
		fmt.Printf("Error parsing measure: %v\n", err)
		return
	}

	font, err := svg.LoadFont("svg/fonts/leland/Leland.otf")
	if err != nil {
		panic(err)
	}

	measureGroup, err := measure.Render(font)
	if err != nil {
		panic(err)
	}

	measureWidth, err := measureGroup.Width(font)
	if err != nil {
		panic(err)
	}

	barlines := svg.Group{
		Elements: []svg.SVGElement{
			{Line: &svg.Line{X1: 0, Y1: 0, X2: int(measureWidth), Y2: 0, Stroke: "#4c4f69", Width: 15}},
			{Line: &svg.Line{X1: 0, Y1: 250, X2: int(measureWidth), Y2: 250, Stroke: "#4c4f69", Width: 15}},
			{Line: &svg.Line{X1: 0, Y1: 500, X2: int(measureWidth), Y2: 500, Stroke: "#4c4f69", Width: 15}},
			{Line: &svg.Line{X1: 0, Y1: 750, X2: int(measureWidth), Y2: 750, Stroke: "#4c4f69", Width: 15}},
			{Line: &svg.Line{X1: 0, Y1: 1000, X2: int(measureWidth), Y2: 1000, Stroke: "#4c4f69", Width: 15}},
		},
		XOffset: 0,
		YOffset: 0,
		Scale:   1,
	}

	fullMeasure := svg.Group{
		Elements: []svg.SVGElement{
			{Group: &barlines},
			{Group: &measureGroup},
		},
		XOffset: 500,
		YOffset: 1000,
		Scale:   1,
	}

	image := svg.SVG{
		Width:  1000,
		Height: 300,
		Elements: []svg.SVGElement{
			{Group: &fullMeasure},
		},
		Scale: 0.1,
	}

	output, err := xml.MarshalIndent(image, "", "  ")
	if err != nil {
		panic(err)
	}

	os.WriteFile("result.xml", output, 0644)
}
