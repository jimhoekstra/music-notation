package musicxml

import (
	"encoding/xml"

	"github.com/jimhoekstra/music-notation/svg"
)

type Barline struct {
	XMLName  xml.Name `xml:"barline"`
	Location string   `xml:"location,attr"`
	BarStyle string   `xml:"bar-style"`
}

func (b Barline) isMusicXMLElement() {}

func (b Barline) Name() string {
	return "Barline: " + b.Location + " " + b.BarStyle
}

// Render returns an svg.Group containing a vertical line spanning the full
// staff height (line 1 at Y=1000 to line 5 at Y=0).
func (b Barline) Render() svg.Group {
	line := svg.Line{
		X1:     0,
		Y1:     int(staveLineToY(1)),
		X2:     0,
		Y2:     int(staveLineToY(5)),
		Stroke: "black",
		Width:  25,
	}
	return svg.Group{
		Elements: []svg.SVGElement{{Line: &line}},
		XOffset:  0,
		YOffset:  0,
		Scale:    1,
	}
}
