package musicxml

// BuildHelloWorld returns a ScorePartWise score containing a
// single whole C4 note in 4/4 time with a treble clef.
func BuildHelloWorld() ScorePartWise {
	attributes := Attributes{
		Divisions: 1,
		Key: Key{
			Fifths: 0,
		},
		Time: Time{
			Beats:    4,
			BeatType: 4,
		},
		Clef: Clef{
			Sign: TrebleClef,
			Line: 2,
		},
	}

	note := Note{
		Pitch: Pitch{
			Step:   "C",
			Octave: 4,
		},
		Duration: 4,
		Type:     "whole",
	}

	measure := Measure{
		Number:     1,
		Attributes: attributes,
		Notes:      []Note{note},
	}

	part_list := PartList{
		ScoreParts: []ScorePart{{
			ID:       "P1",
			PartName: "Music",
		}},
	}

	part := Part{
		ID:       "P1",
		Measures: []Measure{measure},
	}

	score_partwise := ScorePartWise{
		PartList: part_list,
		Version:  "4.1",
		Parts:    []Part{part},
	}

	return score_partwise
}
