package musicxml

import (
	"encoding/xml"
	"fmt"
)

type Attributes struct {
	XMLName   xml.Name `xml:"attributes"`
	Clef      *Clef    `xml:"clef"`
	Key       *Key     `xml:"key"`
	Divisions *int     `xml:"divisions"`
	Time      *Time    `xml:"time"`
}

func (a Attributes) isMusicXMLElement() {}

func (a Attributes) Name() string {
	name := "Attributes:"

	if a.Key != nil {
		name += " " + a.Key.Name()
	}

	if a.Clef != nil {
		name += " " + a.Clef.Name()
	}

	if a.Divisions != nil {
		name += fmt.Sprintf(" Divisions: %d", *a.Divisions)
	}

	if a.Time != nil {
		name += " " + a.Time.Name()
	}

	return name
}
