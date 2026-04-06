package musicxml

import "encoding/xml"

type PartList struct {
	XMLName    xml.Name    `xml:"part-list"`
	ScoreParts []ScorePart `xml:"score-part"`
}

func (p PartList) isMusicXMLElement() {}

func (p PartList) Name() string { return "PartList" }
