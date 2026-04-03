package musicxml

import (
	"encoding/xml"
	"fmt"
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
	Clef      *Clef    `xml:"clef"`
	Key       *Key     `xml:"key"`
	Divisions *int     `xml:"divisions"`
	Time      *Time    `xml:"time"`
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
	Name() string
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

func (s ScorePart) Name() string { return "ScorePart" }
func (p PartList) Name() string  { return "PartList" }
func (k Key) Name() string       { return "Key: " + fmt.Sprintf("%d", k.Fifths) }
func (c Clef) Name() string      { return "Clef: " + string(c.Sign) }
func (t Time) Name() string      { return "Time: " + fmt.Sprintf("%d/%d", t.Beats, t.BeatType) }
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
func (p Pitch) Name() string { return "Pitch" }
func (n Note) Name() string {
	return "Note: " + n.Pitch.Step + fmt.Sprintf("%d", n.Pitch.Octave) + " (" + n.Type + ")"
}
func (m Measure) Name() string       { return "Measure" }
func (p Part) Name() string          { return "Part" }
func (s ScorePartWise) Name() string { return "ScorePartWise" }
func (e EmptyElement) Name() string  { return "EmptyElement" }
