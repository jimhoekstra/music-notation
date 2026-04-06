package musicxml

import "encoding/xml"

type ScorePart struct {
	XMLName  xml.Name `xml:"score-part"`
	ID       string   `xml:"id,attr"`
	PartName string   `xml:"part-name"`
}

func (s ScorePart) isMusicXMLElement() {}

func (s ScorePart) Name() string { return "ScorePart" }
