package svg

import (
	"encoding/xml"
	"strconv"
)

type SVG struct {
	XMLName  xml.Name `xml:"svg"`
	Width    int      `xml:"width,attr"`
	Height   int      `xml:"height,attr"`
	Elements []SVGElement
	Scale    float64
}

func encodeElement(e *xml.Encoder, el SVGElement) error {
	switch {
	case el.Rect != nil:
		if err := e.EncodeElement(el.Rect, xml.StartElement{Name: xml.Name{Local: "rect"}}); err != nil {
			return err
		}

	case el.Line != nil:
		if err := e.EncodeElement(el.Line, xml.StartElement{Name: xml.Name{Local: "line"}}); err != nil {
			return err
		}

	case el.Path != nil:
		if err := e.EncodeElement(el.Path, xml.StartElement{Name: xml.Name{Local: "path"}}); err != nil {
			return err
		}

	case el.Group != nil:
		if err := e.EncodeElement(el.Group, xml.StartElement{Name: xml.Name{Local: "g"}}); err != nil {
			return err
		}

	case el.Character != nil:
		path := el.Character.GetPath()
		if err := e.EncodeElement(path, xml.StartElement{Name: xml.Name{Local: "path"}}); err != nil {
			return err
		}
	}
	return nil
}

func (svg SVG) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "svg"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns"}, Value: "http://www.w3.org/2000/svg"},
		{Name: xml.Name{Local: "id"}, Value: "music-svg"},
		{Name: xml.Name{Local: "width"}, Value: strconv.Itoa(svg.Width)},
		{Name: xml.Name{Local: "height"}, Value: strconv.Itoa(svg.Height)},
		{Name: xml.Name{Local: "viewBox"}, Value: "0 0 " + strconv.Itoa(int(float64(svg.Width)/svg.Scale)) + " " + strconv.Itoa(int(float64(svg.Height)/svg.Scale))},
	}
	e.EncodeToken(start)

	for _, el := range svg.Elements {
		if err := encodeElement(e, el); err != nil {
			return err
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
