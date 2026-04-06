package musicxml

import (
	"encoding/xml"
	"fmt"
)

type Key struct {
	XMLName xml.Name `xml:"key"`
	Fifths  int      `xml:"fifths"`
}

func (k Key) isMusicXMLElement() {}

func (k Key) Name() string { return "Key: " + fmt.Sprintf("%d", k.Fifths) }
