package musicxml

import "encoding/xml"

type Chord struct {
	XMLName xml.Name `xml:"chord"`
}

func (c Chord) isMusicXMLElement() {}

func (c Chord) Name() string {
	return "Chord"
}
