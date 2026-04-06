package musicxml

import (
	"encoding/xml"
	"fmt"
)

type Note struct {
	XMLName  xml.Name `xml:"note"`
	Chord    *Chord   `xml:"chord,omitempty"`
	Pitch    Pitch    `xml:"pitch"`
	Duration int      `xml:"duration"`
	Type     string   `xml:"type"`
}

func (n Note) isMusicXMLElement() {}

func (n Note) Name() string {
	return "Note: " + n.Pitch.Step + fmt.Sprintf("%d", n.Pitch.Octave) + " (" + n.Type + ")"
}
