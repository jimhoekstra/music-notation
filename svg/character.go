package svg

import (
	"fmt"

	"golang.org/x/image/font/sfnt"
)

type Character struct {
	D       string
	Scale   float64
	XOffset float64
	YOffset float64
}

func BuildCharacter(font *sfnt.Font, glyphName rune) (Character, error) {
	pathData, err := GetPathData(font, glyphName)
	if err != nil {
		return Character{}, err
	}

	return Character{
		D:       pathData,
		Scale:   1.0,
		XOffset: 0.0,
		YOffset: 0.0,
	}, nil
}

func getTransformString(scale float64, xOffset float64, yOffset float64) string {
	return fmt.Sprintf("scale(%f) translate(%f, %f)", scale, xOffset, yOffset)
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
