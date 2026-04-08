package musicxml

import (
	"encoding/xml"
	"fmt"

	"github.com/jimhoekstra/music-notation/svg"
	"golang.org/x/image/font/sfnt"
)

type Attributes struct {
	XMLName   xml.Name `xml:"attributes"`
	Clef      *Clef    `xml:"clef"`
	Key       *Key     `xml:"key"`
	Divisions *int     `xml:"divisions"`
	Time      *Time    `xml:"time"`
}

func (a Attributes) isMusicXMLElement() {}

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

// Render returns an svg.Group containing the clef and time signature glyphs.
func (a Attributes) Render(font *sfnt.Font) (svg.Group, error) {
	var elements []svg.SVGElement
	cursor := 0.0

	if a.Clef != nil {
		clefChar, err := a.Clef.Render(font)
		if err != nil {
			return svg.Group{}, fmt.Errorf("cannot render clef: %w", err)
		}
		elements = append(elements, svg.SVGElement{Character: &clefChar})
		w, err := clefChar.Width(font)
		if err != nil {
			return svg.Group{}, fmt.Errorf("cannot get clef width: %w", err)
		}
		cursor += w
	}

	if a.Key != nil {
		keyGroup, err := a.Key.Render(font)
		if err != nil {
			return svg.Group{}, fmt.Errorf("cannot render key signature: %w", err)
		}
		keyGroup.Transform(cursor, 0, 1)
		elements = append(elements, svg.SVGElement{Group: &keyGroup})
		w, err := keyGroup.Width(font)
		if err != nil {
			return svg.Group{}, fmt.Errorf("cannot get key signature width: %w", err)
		}
		cursor += w
	}

	if a.Time != nil {
		timeGroup, err := a.Time.Render(font)
		if err != nil {
			return svg.Group{}, fmt.Errorf("cannot render time signature: %w", err)
		}
		timeGroup.Transform(cursor, 0, 1)
		elements = append(elements, svg.SVGElement{Group: &timeGroup})
	}

	return svg.Group{
		Elements: elements,
		XOffset:  0,
		YOffset:  0,
		Scale:    1,
	}, nil
}
