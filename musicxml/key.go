package musicxml

import (
	"encoding/xml"
	"fmt"

	"github.com/jimhoekstra/music-notation/svg"
	"golang.org/x/image/font/sfnt"
)

type Key struct {
	XMLName xml.Name `xml:"key"`
	Fifths  int      `xml:"fifths"`
}

func (k Key) isMusicXMLElement() {}

func (k Key) Name() string { return "Key: " + fmt.Sprintf("%d", k.Fifths) }

// keyAccidentalGlyph returns the SMuFL glyph rune for a sharp (U+E262) or
// flat (U+E260) key accidental.
func keyAccidentalGlyph(sharp bool) rune {
	if sharp {
		return rune(0xe262)
	}
	return rune(0xe260)
}

// trebleSharpYPositions returns the y-offsets for each sharp in a treble-clef
// key signature, in circle-of-fifths order: F C G D A E B.
func trebleSharpYPositions() []float64 {
	// TODO: Verify exact y-offsets against the staff coordinate system.
	return []float64{0, 375, -125, 250, 625, 125, 500}
}

// trebleFlatYPositions returns the y-offsets for each flat in a treble-clef
// key signature, in circle-of-fifths order: B E A D G C F.
func trebleFlatYPositions() []float64 {
	// TODO: Verify exact y-offsets against the staff coordinate system.
	return []float64{500, 125, 625, 250, 750, 375, 875}
}

// renderAccidentals builds n accidental SVGElements from glyph, spaced
// horizontally by xStep and positioned vertically using yPositions.
func renderAccidentals(font *sfnt.Font, glyph rune, yPositions []float64, n int) ([]svg.SVGElement, error) {
	char, err := svg.BuildCharacter(font, glyph)
	if err != nil {
		return nil, fmt.Errorf("failed to build accidental glyph %d: %w", glyph, err)
	}
	xStep, err := char.GetAdvance(font)
	if err != nil {
		return nil, fmt.Errorf("failed to get advance width for accidental glyph: %w", err)
	}

	elements := make([]svg.SVGElement, 0, n)
	for i := range n {
		c := char
		c.Transform(float64(i)*float64(xStep)/64.0, yPositions[i], 1)
		elements = append(elements, svg.SVGElement{Character: &c})
	}
	return elements, nil
}

// Render returns an svg.Group containing all accidentals for the key signature.
// Currently uses treble-clef positions; clef-awareness is a future TODO.
func (k Key) Render(font *sfnt.Font) (svg.Group, error) {
	if k.Fifths == 0 {
		return svg.Group{}, nil
	}

	sharp := k.Fifths > 0
	count := k.Fifths
	if count < 0 {
		count = -count
	}

	var yPositions []float64
	if sharp {
		yPositions = trebleSharpYPositions()
	} else {
		yPositions = trebleFlatYPositions()
	}

	elements, err := renderAccidentals(font, keyAccidentalGlyph(sharp), yPositions, count)
	if err != nil {
		return svg.Group{}, fmt.Errorf("cannot render key signature: %w", err)
	}

	return svg.Group{
		Elements: elements,
		XOffset:  0,
		YOffset:  0,
		Scale:    1,
	}, nil
}
