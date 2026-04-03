package musicxml

import (
	"encoding/xml"
	"fmt"
	"strconv"
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

type MeasureElement struct {
	Attributes *Attributes
	Note       *Note
	Barline    *Barline
}

type Measure struct {
	XMLName  xml.Name `xml:"measure"`
	Number   int      `xml:"number,attr"`
	Elements []MeasureElement
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

type Barline struct {
	XMLName  xml.Name `xml:"barline"`
	Location string   `xml:"location,attr"`
	BarStyle string   `xml:"bar-style"`
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
func (b Barline) isMusicXMLElement()       {}
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
func (m Measure) Name() string {
	text := fmt.Sprintf("Measure %d:", m.Number)
	for _, el := range m.Elements {
		if el.Attributes != nil {
			text += "\n  " + el.Attributes.Name()
		}
		if el.Note != nil {
			text += "\n  " + el.Note.Name()
		}
		if el.Barline != nil {
			text += "\n  " + el.Barline.Name()
		}
	}

	return text
}
func (p Part) Name() string {
	name := "Part: " + p.ID
	for _, measure := range p.Measures {
		name += "\n" + measure.Name()
	}

	return name
}
func (s ScorePartWise) Name() string { return "ScorePartWise" }
func (e EmptyElement) Name() string  { return "EmptyElement" }

func (m Measure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "measure"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "number"}, Value: strconv.Itoa(m.Number)},
	}
	e.EncodeToken(start)

	for _, el := range m.Elements {
		switch {
		case el.Attributes != nil:
			if err := e.EncodeElement(el.Attributes, xml.StartElement{Name: xml.Name{Local: "attributes"}}); err != nil {
				return err
			}
		case el.Note != nil:
			if err := e.EncodeElement(el.Note, xml.StartElement{Name: xml.Name{Local: "note"}}); err != nil {
				return err
			}
		case el.Barline != nil:
			if err := e.EncodeElement(el.Barline, xml.StartElement{Name: xml.Name{Local: "barline"}}); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (b Barline) Name() string {
	return "Barline: " + b.Location + " " + b.BarStyle
}
