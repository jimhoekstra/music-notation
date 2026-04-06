package musicxml

import (
	"encoding/xml"
)

type Element interface {
	isMusicXMLElement()
	Name() string
}

type EmptyElement struct {
	XMLName xml.Name
}

func (e EmptyElement) isMusicXMLElement() {}
func (e EmptyElement) Name() string       { return "EmptyElement" }
