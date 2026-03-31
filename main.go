package main

import (
	"github.com/jimhoekstra/music-notation/io"
	"github.com/jimhoekstra/music-notation/testing"
)

func main() {
	basic_score := testing.BuildSingleNoteScoreFromElements()
	io.SaveScoreAsMusicXml(basic_score, "result.xml")
}
