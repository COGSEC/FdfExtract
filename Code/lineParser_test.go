package FdfExtract

import (
	"testing"
)

// TestCheckbytes tests unexported checkbytes in lineParser.go
func TestCheckbytes(t *testing.T) {
	a, b := []byte("test1/"), []byte("test/2")
	c, d := []byte("test1/"), []byte("test1/")
	e, f := []byte("test1"), []byte("test1/")
	if checkbytes(a, b) {
		t.Error("false positive")
	}
	if !checkbytes(c, d) {
		t.Error("false negative")
	}
	if checkbytes(e, f) {
		t.Error("false positive")
	}
}

// TestSearchforbytes tests unexported searchforbytes in lineParser.go
func TestSearchforbytes(t *testing.T) {
	a := []byte("This is a larger string Subtype/Content(These are words within it)")
	b := []byte("Subtype/Content(")
	c := []byte("Subtype/Conent(")
	x, y := searchforbytes(a, b)
	if !y {
		t.Error("False Negative")
	}
	if x.startPos != 24 {
		t.Error("Wrong Start Idx Found")
	}
	if x.endPos != 40 {
		t.Error("Wrong end Idx Found")
	}
	x, y = searchforbytes(a, c)
	if y {
		t.Error("False Positive")
	}
	if x.startPos != -1 {
		t.Error("Wrong Idx Returned, should be -1")
	}
}

// TestExtract tests unexported extract in lineParser.go
func TestExtract(t *testing.T) {
	a := []byte("Example sWord/THISISTHEPHRASE/eWord End")
	x, ok := extract(a, []byte("sWord/"), []byte("/eWord"))
	if !ok {
		t.Error("False Negative")
	}
	if string(x) != "THISISTHEPHRASE" {
		t.Error("Failed to pull correct indexes")
	}
	// Next Testcase
	a = []byte("Example sWord/THISISTHEPHRASE/eWord End")
	_, ok = extract(a, []byte("sWordd/"), []byte("/eWord"))
	if ok {
		t.Error("False Positive")
	}
}
