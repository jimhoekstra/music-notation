package testing

import (
	"os"
	"testing"

	"github.com/jimhoekstra/music-notation/musicxml"
)

func TestBuildSingleNoteScoreFromElements(t *testing.T) {
	score := BuildSingleNoteScoreFromElements()
	score_as_bytes := musicxml.BuildXMLBytes(score)

	expected, err := os.ReadFile("single_note.xml")
	if err != nil {
		t.Fatalf("Error reading expected XML file: %v", err)
	}

	equal, err := XMLEqual(score_as_bytes, expected)
	if err != nil {
		t.Fatalf("Error comparing XML: %v", err)
	}

	if !equal {
		t.Errorf("Generated MusicXML does not match expected output.")
	}
}
