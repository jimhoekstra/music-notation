package testing

import (
	"github.com/jimhoekstra/music-notation/musicxml"
)

// BuildSingleNoteScoreFromElements returns a ScorePartWise score containing a
// single measure with a single whole note, built directly from Go structs.
func BuildSingleNoteScoreFromElements() musicxml.ScorePartWise {
	attributes := musicxml.Attributes{
		Key: &musicxml.Key{
			Fifths: 1,
		},
		Clef: &musicxml.Clef{
			Sign: musicxml.TrebleClef,
			Line: 2,
		},
		Divisions: func() *int { i := 1; return &i }(),
		Time: &musicxml.Time{
			Beats:    4,
			BeatType: 4,
		},
	}

	note := musicxml.Note{
		Pitch: musicxml.Pitch{
			Step:   "G",
			Octave: 4,
		},
		Duration: 4,
		Type:     "whole",
	}

	measure := musicxml.Measure{
		Number: 1,
		Elements: []musicxml.MeasureElement{
			{Attributes: &attributes},
			{Note: &note},
		},
	}

	part_list := musicxml.PartList{
		ScoreParts: []musicxml.ScorePart{{
			ID:       "violin",
			PartName: "Violin",
		}},
	}

	part := musicxml.Part{
		ID:       "violin",
		Measures: []musicxml.Measure{measure},
	}

	score_partwise := musicxml.ScorePartWise{
		PartList: part_list,
		Version:  "4.1",
		Parts:    []musicxml.Part{part},
	}

	return score_partwise
}
