package svg

import (
	"encoding/xml"
	"strconv"
)

type Rect struct {
	XMLName xml.Name `xml:"rect"`
	X       int      `xml:"x,attr"`
	Y       int      `xml:"y,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Fill    string   `xml:"fill,attr"`
}

type Line struct {
	XMLName xml.Name `xml:"line"`
	X1      int      `xml:"x1,attr"`
	Y1      int      `xml:"y1,attr"`
	X2      int      `xml:"x2,attr"`
	Y2      int      `xml:"y2,attr"`
	Stroke  string   `xml:"stroke,attr"`
	Width   int      `xml:"stroke-width,attr"`
}

type Path struct {
	XMLName   xml.Name `xml:"path"`
	D         string   `xml:"d,attr"`
	Fill      string   `xml:"fill,attr"`
	Transform string   `xml:"transform,attr"`
}

type SVGElement struct {
	Rect *Rect
	Line *Line
	Path *Path
	// Character is a special utility element that gets converted to a Path
	// during XML marshaling.
	Character *Character
}

type SVG struct {
	XMLName  xml.Name `xml:"svg"`
	Width    int      `xml:"width,attr"`
	Height   int      `xml:"height,attr"`
	Elements []SVGElement
}

func (svg SVG) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "svg"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "xmlns"}, Value: "http://www.w3.org/2000/svg"},
		{Name: xml.Name{Local: "width"}, Value: strconv.Itoa(svg.Width)},
		{Name: xml.Name{Local: "height"}, Value: strconv.Itoa(svg.Height)},
		{Name: xml.Name{Local: "viewBox"}, Value: "0 0 " + strconv.Itoa(svg.Width) + " " + strconv.Itoa(svg.Height)},
	}
	e.EncodeToken(start)

	for _, el := range svg.Elements {
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

		case el.Character != nil:
			path := el.Character.GetPath()
			if err := e.EncodeElement(path, xml.StartElement{Name: xml.Name{Local: "path"}}); err != nil {
				return err
			}
		}
	}

	return e.EncodeToken(xml.EndElement{Name: start.Name})
}
