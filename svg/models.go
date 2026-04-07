package svg

import (
	"encoding/xml"
)

type Rect struct {
	XMLName xml.Name `xml:"rect"`
	X       int      `xml:"x,attr"`
	Y       int      `xml:"y,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Fill    string   `xml:"fill,attr"`
}

type Line struct {
	XMLName xml.Name `xml:"line"`
	X1      int      `xml:"x1,attr"`
	Y1      int      `xml:"y1,attr"`
	X2      int      `xml:"x2,attr"`
	Y2      int      `xml:"y2,attr"`
	Stroke  string   `xml:"stroke,attr"`
	Width   int      `xml:"stroke-width,attr"`
}

type Path struct {
	XMLName   xml.Name `xml:"path"`
	D         string   `xml:"d,attr"`
	Fill      string   `xml:"fill,attr"`
	Transform string   `xml:"transform,attr"`
}

type SVGElement struct {
	Rect  *Rect
	Line  *Line
	Path  *Path
	Group *Group
	// Character is a special utility element that gets converted to a Path
	// during XML marshaling.
	Character *Character
}
