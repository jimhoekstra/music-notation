package musicxml

import (
	"encoding/xml"
	"fmt"

	"github.com/jimhoekstra/music-notation/svg"
	"golang.org/x/image/font/sfnt"
)

type Time struct {
	XMLName  xml.Name `xml:"time"`
	Beats    int      `xml:"beats"`
	BeatType int      `xml:"beat-type"`
}

func (t Time) isMusicXMLElement() {}

func (t Time) Name() string { return "Time: " + fmt.Sprintf("%d/%d", t.Beats, t.BeatType) }

// timeSignatureGlyph returns the SMuFL glyph rune for a single-digit time
// signature numeral (0–9). Digits 0–9 map to code points 0xe080–0xe089.
func timeSignatureGlyph(n int) rune {
	return rune(0xe080 + n)
}

// Render returns an svg.Group containing two characters: Beats on top and
// BeatType below it, positioned as a standard time signature.
func (t Time) Render(font *sfnt.Font) (svg.Group, error) {
	beatsGlyph := timeSignatureGlyph(t.Beats)
	beatTypeGlyph := timeSignatureGlyph(t.BeatType)

	beatsChar, err := svg.BuildCharacter(font, beatsGlyph)
	if err != nil {
		return svg.Group{}, fmt.Errorf("cannot render time signature beats %d: %w", t.Beats, err)
	}

	beatTypeChar, err := svg.BuildCharacter(font, beatTypeGlyph)
	if err != nil {
		return svg.Group{}, fmt.Errorf("cannot render time signature beat type %d: %w", t.BeatType, err)
	}

	beatsChar.Transform(0, 250, 1)
	beatTypeChar.Transform(0, 750, 1)

	return svg.Group{
		Elements: []svg.SVGElement{
			{Character: &beatsChar},
			{Character: &beatTypeChar},
		},
		XOffset: 0,
		YOffset: 0,
		Scale:   1,
	}, nil
}
