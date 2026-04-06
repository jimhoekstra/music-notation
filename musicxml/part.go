package musicxml

import "encoding/xml"

type Part struct {
	XMLName  xml.Name  `xml:"part"`
	ID       string    `xml:"id,attr"`
	Measures []Measure `xml:"measure"`
}

func (p Part) isMusicXMLElement() {}

func (p Part) Name() string {
	name := "Part: " + p.ID
	for _, measure := range p.Measures {
		name += "\n" + measure.Name()
	}

	return name
}
