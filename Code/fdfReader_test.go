package FdfExtract

import (
	"fmt"
	"testing"
)

//TestNewReader tests exported NewReader func in fdfreader.go
func TestNewReader(t *testing.T) {
	reader := NewReader()
	reader.disableConcurrency()
}

// TestReader1 Tests reader // good comments good path
func TestReader1(t *testing.T) {
	path := "_testfiles/workingFdf0.fdf"
	reader := NewReader()
	_, err := reader.Read(path)
	if err == nil {
		t.Error("reader failed to catch error")
	}
}

// TestReader1 Tests reader // bad headings
func TestReader2(t *testing.T) {
	path := "_testfiles/workingFdf.fdf"
	reader := NewReader()
	cmnts, err := reader.Read(path)
	if err != nil {
		fmt.Println(err.Error())
		t.Error("reader failed")
	} else if len(cmnts.GetComments()) != 4 {
		t.Error("reader missed comments")
	}
	cmntlist := cmnts.GetComments()
	if (string(cmntlist[0].GetContent()[0:5])) != "$NOTE" {
		t.Error("comment not grabbed correctly")
	}
}

// TestReader3 Tests reader // no file
func TestReader3(t *testing.T) {
	path := "_testfiles/workinFdf1.fdf"
	reader := NewReader()
	_, err := reader.Read(path)
	if err == nil {
		t.Error("reader failed to catch error")
	}
}

// TestReader4 Tests reader // bad extension
func TestReader4(t *testing.T) {
	path := "_testfiles/workinFdf1.txt"
	reader := NewReader()
	_, err := reader.Read(path)
	if err == nil {
		t.Error("reader failed to catch error")
	}
}
