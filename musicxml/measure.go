package musicxml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

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

func (m Measure) isMusicXMLElement() {}

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
