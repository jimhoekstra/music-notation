package musicxml

import (
	"strings"
	"testing"
)

func TestBuildXMLBytes_ContainsHeader(t *testing.T) {
	score := ScorePartWise{}
	result := string(BuildXMLBytes(score))

	if !strings.HasPrefix(result, Header) {
		t.Error("expected result to start with XML header")
	}
}

func TestBuildXMLBytes_ContainsRootElement(t *testing.T) {
	score := ScorePartWise{Version: "4.1"}
	result := string(BuildXMLBytes(score))

	if !strings.Contains(result, "<score-partwise") {
		t.Error("expected result to contain a <score-partwise> tag")
	}
	if !strings.Contains(result, `version="4.1"`) {
		t.Errorf("expected result to contain version attribute")
	}
}
