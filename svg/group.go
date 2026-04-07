package svg

import "encoding/xml"

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
