package musicxml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"github.com/jimhoekstra/music-notation/svg"
	"golang.org/x/image/font/sfnt"
)

// ElementType identifies the type of a rendered measure element.
type ElementType int

const (
	AttributesElement ElementType = iota
	ClefElement
	KeyElement
	TimeElement
	WholeNoteElement
	HalfNoteElement
	QuarterNoteElement
	EighthNoteElement
	SixteenthNoteElement
	BarlineElement
	MeasureStartElement
	WildcardElement // matches any element type in the second position of a SpacingTable key
)

// noteElementType maps a MusicXML note type string to the corresponding ElementType.
func noteElementType(noteType string) ElementType {
	switch noteType {
	case "whole":
		return WholeNoteElement
	case "half":
		return HalfNoteElement
	case "quarter":
		return QuarterNoteElement
	case "eighth":
		return EighthNoteElement
	case "16th":
		return SixteenthNoteElement
	default:
		return QuarterNoteElement
	}
}

// SpacingTable maps pairs of consecutive element types to the minimum
// required horizontal spacing between them.
type SpacingTable map[[2]ElementType]float64

// Lookup returns the spacing for the pair (from, to), falling back to
// (from, WildcardElement) if no exact match exists.
func (t SpacingTable) Lookup(from, to ElementType) float64 {
	if spacing, ok := t[[2]ElementType{from, to}]; ok {
		return spacing * 250
	}
	if spacing, ok := t[[2]ElementType{from, WildcardElement}]; ok {
		return spacing * 250
	}
	if spacing, ok := t[[2]ElementType{WildcardElement, to}]; ok {
		return spacing * 250
	}
	return t[[2]ElementType{WildcardElement, WildcardElement}] * 250
}

// DefaultSpacingTable defines the default minimum spacing between element pairs.
var DefaultSpacingTable = SpacingTable{
	{MeasureStartElement, ClefElement}:      2.0 / 3.0,
	{ClefElement, KeyElement}:               2.0 / 3.0,
	{ClefElement, TimeElement}:              2.0 / 3.0,
	{KeyElement, TimeElement}:               1.0 / 3.0,
	{TimeElement, WildcardElement}:          1.5,
	{WholeNoteElement, WildcardElement}:     5.0,
	{HalfNoteElement, WildcardElement}:      4.0,
	{QuarterNoteElement, WildcardElement}:   3.0,
	{EighthNoteElement, WildcardElement}:    2.5,
	{SixteenthNoteElement, WildcardElement}: 2.0,
	{WildcardElement, BarlineElement}:       1.0,
	{WildcardElement, WildcardElement}:      1.0,
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
	prevType := MeasureStartElement
	cursor := 0.0

	for _, el := range m.Elements {
		switch {
		case el.Attributes != nil:
			attrsGroup, newPrevType, err := el.Attributes.Render(font, DefaultSpacingTable, prevType)
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
			prevType = newPrevType

		case el.Note != nil:
			currentType := noteElementType(el.Note.Type)
			cursor += DefaultSpacingTable.Lookup(prevType, currentType)

			noteGroup, err := el.Note.Render(font)
			if err != nil {
				return svg.Group{}, fmt.Errorf("cannot render note in measure %d: %w", m.Number, err)
			}
			noteGroup.Transform(cursor, 0, 1)
			// w, err := noteGroup.Width(font)
			// if err != nil {
			// 	return svg.Group{}, fmt.Errorf("cannot get note width in measure %d: %w", m.Number, err)
			// }
			// cursor += w
			elements = append(elements, svg.SVGElement{Group: &noteGroup})
			prevType = currentType

		case el.Barline != nil:
			currentType := BarlineElement
			cursor += DefaultSpacingTable.Lookup(prevType, currentType)

			barlineGroup := el.Barline.Render()
			barlineGroup.Transform(cursor, 0, 1)

			elements = append(elements, svg.SVGElement{Group: &barlineGroup})
			prevType = currentType
		}

	}

	return svg.Group{
		Elements: elements,
		XOffset:  0,
		YOffset:  0,
		Scale:    1,
	}, nil
}
