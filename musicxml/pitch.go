package musicxml

import "encoding/xml"

type Pitch struct {
	XMLName xml.Name `xml:"pitch"`
	Step    string   `xml:"step"`
	Octave  int      `xml:"octave"`
}

func (p Pitch) isMusicXMLElement() {}

func (p Pitch) Name() string { return "Pitch" }
