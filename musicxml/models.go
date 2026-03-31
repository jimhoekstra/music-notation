package musicxml

import (
	"encoding/xml"
)

type ClefSign string

const (
	TrebleClef ClefSign = "G"
	BassClef   ClefSign = "F"
)

type ScorePart struct {
	XMLName  xml.Name `xml:"score-part"`
	ID       string   `xml:"id,attr"`
	PartName string   `xml:"part-name"`
}

type PartList struct {
	XMLName    xml.Name    `xml:"part-list"`
	ScoreParts []ScorePart `xml:"score-part"`
}

type Key struct {
	XMLName xml.Name `xml:"key"`
	Fifths  int      `xml:"fifths"`
}

type Clef struct {
	XMLName xml.Name `xml:"clef"`
	Sign    ClefSign `xml:"sign"`
	Line    int      `xml:"line"`
}

type Time struct {
	XMLName  xml.Name `xml:"time"`
	Beats    int      `xml:"beats"`
	BeatType int      `xml:"beat-type"`
}

type Attributes struct {
	XMLName   xml.Name `xml:"attributes"`
	Divisions int      `xml:"divisions"`
	Key       Key      `xml:"key"`
	Time      Time     `xml:"time"`
	Clef      Clef     `xml:"clef"`
}

type Pitch struct {
	XMLName xml.Name `xml:"pitch"`
	Step    string   `xml:"step"`
	Octave  int      `xml:"octave"`
}

type Note struct {
	XMLName  xml.Name `xml:"note"`
	Pitch    Pitch    `xml:"pitch"`
	Duration int      `xml:"duration"`
	Type     string   `xml:"type"`
}

type Measure struct {
	XMLName    xml.Name   `xml:"measure"`
	Number     int        `xml:"number,attr"`
	Attributes Attributes `xml:"attributes"`
	Notes      []Note     `xml:"note"`
}

type Part struct {
	XMLName  xml.Name  `xml:"part"`
	ID       string    `xml:"id,attr"`
	Measures []Measure `xml:"measure"`
}

type ScorePartWise struct {
	XMLName  xml.Name `xml:"score-partwise"`
	Version  string   `xml:"version,attr"`
	PartList PartList `xml:"part-list"`
	Parts    []Part   `xml:"part"`
}

type EmptyElement struct {
	XMLName xml.Name
}

type Element interface {
	isMusicXMLElement()
}

func (s ScorePart) isMusicXMLElement()     {}
func (p PartList) isMusicXMLElement()      {}
func (k Key) isMusicXMLElement()           {}
func (c Clef) isMusicXMLElement()          {}
func (t Time) isMusicXMLElement()          {}
func (a Attributes) isMusicXMLElement()    {}
func (p Pitch) isMusicXMLElement()         {}
func (n Note) isMusicXMLElement()          {}
func (m Measure) isMusicXMLElement()       {}
func (p Part) isMusicXMLElement()          {}
func (s ScorePartWise) isMusicXMLElement() {}
func (e EmptyElement) isMusicXMLElement()  {}
