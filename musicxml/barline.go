package musicxml

import "encoding/xml"

type Barline struct {
	XMLName  xml.Name `xml:"barline"`
	Location string   `xml:"location,attr"`
	BarStyle string   `xml:"bar-style"`
}

func (b Barline) isMusicXMLElement() {}

func (b Barline) Name() string {
	return "Barline: " + b.Location + " " + b.BarStyle
}
