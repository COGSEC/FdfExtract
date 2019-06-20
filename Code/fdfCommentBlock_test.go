package FdfExtract

import (
	"testing"
)

// Tests struct and methods of fdfCommentBlock.go
func TestFdfCommentBlock(t *testing.T) {
	x := make([]Comment, 10)
	var err error
	for i := range x {
		x[i], err = NewComment([]byte("Test"), 2)
		if err != nil {
			break
		}
	}
	cb := NewCommentBlock([]byte("Test"), x)
	cmnts := cb.GetComments()
	path := string(cb.GetSourceFile())
	if path != "Test" {
		t.Error("Path failed to set or be retrieved")
	}
	for i := range x {
		if string(x[i].GetContent()) != string(cmnts[i].GetContent()) {
			t.Error("Comments failed to set or be retrieved")
		}
	}
}
