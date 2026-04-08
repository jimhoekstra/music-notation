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

// Render returns an svg.Group containing the notehead glyph positioned at the
// correct vertical staff position for the note's pitch.
func (n Note) Render(font *sfnt.Font) (svg.Group, error) {
	glyph := noteheadGlyph(n.Type)

	notehead, err := svg.BuildCharacter(font, glyph)
	if err != nil {
		return svg.Group{}, fmt.Errorf("cannot render notehead for %s: %w", n.Type, err)
	}

	y := staveLineToY(pitchToStaveLine(n.Pitch))

	notehead.Transform(0, y, 1)

	return svg.Group{
		Elements: []svg.SVGElement{
			{Character: &notehead},
		},
		XOffset: 0,
		YOffset: 0,
		Scale:   1,
	}, nil
}
