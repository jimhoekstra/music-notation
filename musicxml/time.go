package musicxml

import (
	"encoding/xml"
	"fmt"
)

type Time struct {
	XMLName  xml.Name `xml:"time"`
	Beats    int      `xml:"beats"`
	BeatType int      `xml:"beat-type"`
}

func (t Time) isMusicXMLElement() {}

func (t Time) Name() string { return "Time: " + fmt.Sprintf("%d/%d", t.Beats, t.BeatType) }
