package musicxml

import (
	"encoding/xml"
	"errors"

	"github.com/jimhoekstra/music-notation/svg"
	"golang.org/x/image/font/sfnt"
)

type ClefSign string

const (
	TrebleClef ClefSign = "G"
	BassClef   ClefSign = "F"
)

type Clef struct {
	XMLName xml.Name `xml:"clef"`
	Sign    ClefSign `xml:"sign"`
	Line    int      `xml:"line"`
}

func (c Clef) isMusicXMLElement() {}

func (c Clef) Name() string { return "Clef: " + string(c.Sign) }

func staveLineToY(line float64) float64 {
	return 1250 - (line * 250)
}

func (c Clef) Render(font *sfnt.Font) (svg.Character, error) {
	switch c.Sign {
	case TrebleClef:
		glyphName := rune(0xe050)
		character, err := svg.BuildCharacter(font, glyphName)
		if err != nil {
			return svg.Character{}, err
		}
		character.Transform(0, staveLineToY(float64(c.Line)), 1)
		return character, nil

	case BassClef:
		glyphName := rune(0xe062)
		character, err := svg.BuildCharacter(font, glyphName)
		if err != nil {
			return svg.Character{}, err
		}
		character.Transform(0, staveLineToY(float64(c.Line)), 1)
		return character, nil

	default:
		return svg.Character{}, errors.New("cannot render unknown clef sign: " + string(c.Sign))
	}
}
