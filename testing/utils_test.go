package testing

import "testing"

func assertEqual(t *testing.T, a, b string) {
	t.Helper()

	equal, err := XMLEqual([]byte(a), []byte(b))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !equal {
		t.Error("expected identical XML strings to be equal")
	}
}

func assertNotEqual(t *testing.T, a, b string) {
	t.Helper()

	equal, err := XMLEqual([]byte(a), []byte(b))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if equal {
		t.Error("expected different XML strings to not be equal")
	}
}

func TestXMLEqual_IdenticalXML(t *testing.T) {
	a := `<note><pitch>C</pitch></note>`
	assertEqual(t, a, a)
}

func TestXMLEqual_IdenticalXMLWithDifferentFormatting(t *testing.T) {
	a := "<note><pitch>C</pitch></note>"
	b := `<note>
	<pitch>
		C
	</pitch>
</note>`

	assertEqual(t, a, b)
}

func TestXMLEqual_DifferentContent(t *testing.T) {
	a := `<note><pitch>C</pitch></note>`
	b := `<note><pitch>D</pitch></note>`
	assertNotEqual(t, a, b)
}

func TestXMLEqual_DifferentOrder(t *testing.T) {
	a := `<note><pitch>C</pitch><pitch>D</pitch></note>`
	b := `<note><pitch>D</pitch><pitch>C</pitch></note>`
	assertNotEqual(t, a, b)
}

func TestXMLEqual_InvalidXML(t *testing.T) {
	_, err := XMLEqual([]byte("not xml"), []byte(`<note/>`))
	if err == nil {
		t.Error("expected an error for invalid XML input")
	}
}
