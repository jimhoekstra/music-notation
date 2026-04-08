package musicxml

import (
	"encoding/xml"
	"fmt"

	"github.com/jimhoekstra/music-notation/svg"
	"golang.org/x/image/font/sfnt"
)

type Note struct {
	XMLName  xml.Name `xml:"note"`
	Chord    *Chord   `xml:"chord,omitempty"`
	Pitch    Pitch    `xml:"pitch"`
	Duration int      `xml:"duration"`
	Type     string   `xml:"type"`
}

func (n Note) isMusicXMLElement() {}

func (n Note) Name() string {
	return "Note: " + n.Pitch.Step + fmt.Sprintf("%d", n.Pitch.Octave) + " (" + n.Type + ")"
}

// noteheadGlyph returns the SMuFL glyph rune for a notehead based on note type.
// Whole notes use an open notehead (U+E0A2), half notes use a half notehead
// (U+E0A3), and all shorter durations use a filled notehead (U+E0A4).
func noteheadGlyph(noteType string) rune {
	switch noteType {
	case "whole":
		return rune(0xe0a2)
	case "half":
		return rune(0xe0a3)
	default:
		return rune(0xe0a4)
	}
}

// stepIndex maps a pitch step letter to a diatonic index (C=0 through B=6).
func stepIndex(step string) int {
	switch step {
	case "C":
		return 0
	case "D":
		return 1
	case "E":
		return 2
	case "F":
		return 3
	case "G":
		return 4
	case "A":
		return 5
	case "B":
		return 6
	default:
		return 0
	}
}

// staffPosition returns a linear diatonic position for a pitch, where C4 = 0,
// D4 = 1, ..., B4 = 6, C5 = 7, etc.
// TODO: does it make sense to have C4 as the reference point?
func staffPosition(p Pitch) int {
	return (p.Octave-4)*7 + stepIndex(p.Step)
}

// pitchToStaveLine maps a pitch to a stave line position assuming treble clef.
// Integer values (1–5) are the staff lines; half-values (e.g. 1.5) are spaces.
// E4 is line 1, G4 is line 2, B4 is line 3, D5 is line 4, F5 is line 5.
// TODO: accept a Clef argument and adjust the reference pitch accordingly.
func pitchToStaveLine(p Pitch) float64 {
	return float64(staffPosition(p)) / 2.0
}

// stemLength is the standard stem length in SVG coordinate units (3.5 staff spaces).
const stemLength = 875

// needsStem returns true for all note types that require a stem.
func needsStem(noteType string) bool {
	return noteType != "whole"
}

// stemUp returns true when the note should have a stem pointing upward.
// Notes at or below the middle line (B4 = stave line 3.0) get a stem up.
func stemUp(staveLine float64) bool {
	return staveLine <= 3.0
}

// buildStem constructs a stem svg.Line from the notehead position to 3.5 staff
// spaces above (stem up) or below (stem down). For stem up the line attaches to
// the right side of the notehead; for stem down it attaches to the left side.
// TODO: organize the magic numbers here a bit better.
func buildStem(noteheadY, noteheadWidth float64, up bool) svg.Line {
	x := int(noteheadWidth) - 13
	if !up {
		x = 13
	}
	y1 := int(noteheadY) - 40
	y2 := y1 - stemLength + 40
	if !up {
		y1 = int(noteheadY) + 40
		y2 = y1 + stemLength - 40
	}
	return svg.Line{
		X1:     x,
		Y1:     y1,
		X2:     x,
		Y2:     y2,
		Stroke: "black",
		Width:  25,
	}
}

// flagGlyph returns the SMuFL glyph rune for a flag and true when the note type
// requires one. Up and down stems use different glyph variants.
// Eighth: U+E240/E241, 16th: U+E242/E243, 32nd: U+E244/E245.
func flagGlyph(noteType string, up bool) (rune, bool) {
	switch noteType {
	case "eighth":
		if up {
			return rune(0xe240), true
		}
		return rune(0xe241), true
	case "16th":
		if up {
			return rune(0xe242), true
		}
		return rune(0xe243), true
	case "32nd":
		if up {
			return rune(0xe244), true
		}
		return rune(0xe245), true
	default:
		return 0, false
	}
}

// RenderHead returns the notehead glyph positioned at the correct vertical
// staff position for the note's pitch, along with the Y coordinate used for
// placement. The Y value is needed by the caller when building a stem or flag
// for a chord containing this note.
func (n Note) RenderHead(font *sfnt.Font) (svg.Character, float64, error) {
	glyph := noteheadGlyph(n.Type)

	notehead, err := svg.BuildCharacter(font, glyph)
	if err != nil {
		return svg.Character{}, 0, fmt.Errorf("cannot render notehead for %s: %w", n.Type, err)
	}

	staveLine := pitchToStaveLine(n.Pitch)
	y := staveLineToY(staveLine)
	notehead.Transform(0, y, 1)

	return notehead, y, nil
}

// renderStemAndFlag builds the stem line and, when applicable, the flag glyph
// for a note or chord. noteheadY is the Y position of the note at the stem
// base (lowest note for stem-up, highest for stem-down), noteheadWidth is the
// width of that notehead, and noteType drives the flag selection.
func renderStemAndFlag(font *sfnt.Font, noteheadY, noteheadWidth float64, noteType string, up bool) ([]svg.SVGElement, error) {
	stem := buildStem(noteheadY, noteheadWidth, up)
	elements := []svg.SVGElement{{Line: &stem}}

	if flagRune, hasFlag := flagGlyph(noteType, up); hasFlag {
		flagChar, err := svg.BuildCharacter(font, flagRune)
		if err != nil {
			return nil, fmt.Errorf("cannot render flag for %s: %w", noteType, err)
		}
		flagX := noteheadWidth - 25
		flagY := noteheadY - stemLength
		if !up {
			flagX = 0
			flagY = noteheadY + stemLength
		}
		flagChar.Transform(flagX, flagY, 1)
		elements = append(elements, svg.SVGElement{Character: &flagChar})
	}

	return elements, nil
}

// Render returns an svg.Group containing the notehead, stem, and flag (where
// applicable) for a single note.
func (n Note) Render(font *sfnt.Font) (svg.Group, error) {
	notehead, y, err := n.RenderHead(font)
	if err != nil {
		return svg.Group{}, err
	}

	elements := []svg.SVGElement{{Character: &notehead}}

	if needsStem(n.Type) {
		noteheadWidth, err := notehead.Width(font)
		if err != nil {
			return svg.Group{}, fmt.Errorf("cannot get notehead width: %w", err)
		}

		staveLine := pitchToStaveLine(n.Pitch)
		up := stemUp(staveLine)
		stemElements, err := renderStemAndFlag(font, y, noteheadWidth, n.Type, up)
		if err != nil {
			return svg.Group{}, err
		}
		elements = append(elements, stemElements...)
	}

	return svg.Group{
		Elements: elements,
		XOffset:  0,
		YOffset:  0,
		Scale:    1,
	}, nil
}
