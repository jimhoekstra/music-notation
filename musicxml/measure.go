package musicxml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/jimhoekstra/music-notation/svg"
	"golang.org/x/image/font/sfnt"
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

// Render returns an svg.Group containing the rendered Attributes and Notes of
// the measure, laid out horizontally using a running cursor. Barlines are
// ignored. Each element is positioned by calling Transform on the returned
// group before appending it.
// TODO: add proper spacing between elements, and handle barlines.
func (m Measure) Render(font *sfnt.Font) (svg.Group, error) {
	var elements []svg.SVGElement
	cursor := 0.0

	for _, el := range m.Elements {
		switch {
		case el.Attributes != nil:
			attrsGroup, err := el.Attributes.Render(font)
			if err != nil {
				return svg.Group{}, fmt.Errorf("cannot render attributes in measure %d: %w", m.Number, err)
			}
			attrsGroup.Transform(cursor, 0, 1)
			w, err := attrsGroup.Width(font)
			if err != nil {
				return svg.Group{}, fmt.Errorf("cannot get attributes width in measure %d: %w", m.Number, err)
			}
			cursor += w
			elements = append(elements, svg.SVGElement{Group: &attrsGroup})

		case el.Note != nil:
			noteGroup, err := el.Note.Render(font)
			if err != nil {
				return svg.Group{}, fmt.Errorf("cannot render note in measure %d: %w", m.Number, err)
			}
			noteGroup.Transform(cursor, 0, 1)
			w, err := noteGroup.Width(font)
			if err != nil {
				return svg.Group{}, fmt.Errorf("cannot get note width in measure %d: %w", m.Number, err)
			}
			cursor += w
			elements = append(elements, svg.SVGElement{Group: &noteGroup})
		}
	}

	return svg.Group{
		Elements: elements,
		XOffset:  0,
		YOffset:  0,
		Scale:    1,
	}, nil
}
