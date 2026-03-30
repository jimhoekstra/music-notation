package io

import (
	"os"

	"github.com/jimhoekstra/music-notation/musicxml"
)

// SaveScoreAsMusicXml writes a ScorePartWise score as a
// MusicXML file to the given path.
func SaveScoreAsMusicXml(score musicxml.ScorePartWise, path string) {
	os.WriteFile(path, musicxml.BuildXMLBytes(score), 0644)
}

// LoadInputFile reads the contents of a file at the given path
// and returns it as a string. This expects files to contain
// musical notation in a format that the parser can process.
func LoadInputFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
