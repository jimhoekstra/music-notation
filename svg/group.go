package svg

import (
	"encoding/xml"
	"fmt"

	"golang.org/x/image/font/sfnt"
)

type Group struct {
	XMLName  xml.Name `xml:"g"`
	Elements []SVGElement
	XOffset  float64
	YOffset  float64
	Scale    float64
}

func (g *Group) Transform(x, y, scale float64) {
	g.XOffset += x
	g.YOffset += y
	g.Scale *= scale
}

// Width returns the total width of the group in SVG coordinate units.
// It is computed as the group's XOffset plus the maximum right edge
// (XOffset + Width) across all character and nested group elements.
func (g Group) Width(font *sfnt.Font) (float64, error) {
	maxRight := 0.0
	for _, el := range g.Elements {
		switch {
		case el.Character != nil:
			w, err := el.Character.Width(font)
			if err != nil {
				return 0, fmt.Errorf("cannot get character width: %w", err)
			}
			right := el.Character.XOffset + w
			if right > maxRight {
				maxRight = right
			}
		case el.Group != nil:
			w, err := el.Group.Width(font)
			if err != nil {
				return 0, fmt.Errorf("cannot get group width: %w", err)
			}
			right := el.Group.XOffset + w
			if right > maxRight {
				maxRight = right
			}
		}
	}
	return maxRight, nil
}

func (g Group) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "g"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "transform"}, Value: getTransformString(g.Scale, g.XOffset, g.YOffset)},
	}
	e.EncodeToken(start)

	for _, el := range g.Elements {
		if err := encodeElement(e, el); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
