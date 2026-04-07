package svg

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
)

func LoadFont(path string) (*sfnt.Font, error) {
	fontBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	font, err := sfnt.Parse(fontBytes)
	if err != nil {
		return nil, err
	}

	return font, nil
}

func coordToString(coord fixed.Point26_6) string {
	return fmt.Sprintf("%d, %d", int(coord.X)/64, int(coord.Y)/64)
}

func GetPathData(font *sfnt.Font, glyphName rune) (string, error) {
	var buf sfnt.Buffer
	glyphIndex, err := font.GlyphIndex(&buf, glyphName)
	if err != nil {
		return "", err
	}

	segments, err := font.LoadGlyph(&buf, glyphIndex, fixed.Int26_6(1000<<6), nil)
	if err != nil {
		return "", err
	}

	svgPath := ""

	for _, segment := range segments {
		p := segment.Args
		switch segment.Op {
		case sfnt.SegmentOpMoveTo:
			svgPath += fmt.Sprintf("M %s ", coordToString(p[0]))
		case sfnt.SegmentOpLineTo:
			svgPath += fmt.Sprintf("L %s ", coordToString(p[0]))
		case sfnt.SegmentOpQuadTo:
			svgPath += fmt.Sprintf("Q %s %s ", coordToString(p[0]), coordToString(p[1]))
		case sfnt.SegmentOpCubeTo:
			svgPath += fmt.Sprintf("C %s %s %s ", coordToString(p[0]), coordToString(p[1]), coordToString(p[2]))
		}
	}

	svgPath += "Z"
	return svgPath, nil
}

func GetGlyphAdvance(f *sfnt.Font, glyphName rune) (fixed.Int26_6, error) {
	var buf sfnt.Buffer
	glyphIndex, err := f.GlyphIndex(&buf, glyphName)
	if err != nil {
		return 0, err
	}

	advance, err := f.GlyphAdvance(&buf, glyphIndex, fixed.Int26_6(1000<<6), font.HintingNone)
	if err != nil {
		return 0, err
	}

	return advance, nil
}
