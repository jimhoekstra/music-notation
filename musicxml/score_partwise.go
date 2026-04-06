package musicxml

import "encoding/xml"

type ScorePartWise struct {
	XMLName  xml.Name `xml:"score-partwise"`
	Version  string   `xml:"version,attr"`
	PartList PartList `xml:"part-list"`
	Parts    []Part   `xml:"part"`
}

func (s ScorePartWise) isMusicXMLElement() {}

func (s ScorePartWise) Name() string { return "ScorePartWise" }
