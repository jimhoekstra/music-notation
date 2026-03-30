package musicxml

import (
	"encoding/xml"
)

const Header string = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE score-partwise PUBLIC
    "-//Recordare//DTD MusicXML 4.1 Partwise//EN"
    "http://www.musicxml.org/dtds/partwise.dtd">
`

// buildXMLBytes marshals a ScorePartWise score to a
// MusicXML byte slice, prepending the XML declaration
// and DOCTYPE header.
func BuildXMLBytes(score ScorePartWise) []byte {
	output, err := xml.MarshalIndent(score, "", "  ")
	if err != nil {
		panic(err)
	}

	result := []byte(Header + string(output))
	return result
}
