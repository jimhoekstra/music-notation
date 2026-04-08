package svg

import (
	"fmt"

	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

type Character struct {
	r       rune
	D       string
	Scale   float64
	XOffset float64
	YOffset float64
}

func (c Character) GetAdvance(font *sfnt.Font) (fixed.Int26_6, error) {
	return GetGlyphAdvance(font, c.r)
}

func BuildCharacter(font *sfnt.Font, glyphName rune) (Character, error) {
	pathData, err := GetPathData(font, glyphName)
	if err != nil {
		return Character{}, err
	}

	return Character{
		r:       glyphName,
		D:       pathData,
		Scale:   1.0,
		XOffset: 0.0,
		YOffset: 0.0,
	}, nil
}

func getTransformString(scale float64, xOffset float64, yOffset float64) string {
	return fmt.Sprintf("translate(%f, %f) scale(%f)", xOffset, yOffset, scale)
}

func (c Character) GetPath() Path {
	return Path{
		D:         c.D,
		Fill:      "black",
		Transform: getTransformString(c.Scale, c.XOffset, c.YOffset),
	}
}

func (c *Character) Transform(x, y, scale float64) {
	c.XOffset += x
	c.YOffset += y
	c.Scale *= scale
}

// Width returns the advance width of the character glyph in SVG coordinate units.
func (c Character) Width(font *sfnt.Font) (float64, error) {
	advance, err := c.GetAdvance(font)
	if err != nil {
		return 0, err
	}
	return float64(advance) / 64.0, nil
}
