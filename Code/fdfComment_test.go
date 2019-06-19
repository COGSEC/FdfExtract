package FdfExtract

import (
	"errors"
	"testing"
)

// TestCheckCommentInput tests comment input checker in fdfComment.go
func TestCheckCommentInput(t *testing.T) {
	testcases := []struct {
		testCaseName string
		content      []byte
		pagenum      int
		err          error
	}{
		{"Inputs Okay", []byte("Test Note"), 23, nil},
		{"bad page number", []byte("Test Note"), -23, errors.New("Create Comment Failed: content exists [true], page number not negative [false]")},
		{"bad content", []byte(""), 23, errors.New("Create Comment Failed: content exists [false], page number not negative [true]")},
		{"bad content and page number", []byte(""), -23, errors.New("Create Comment Failed: content exists [false], page number not negative [false]")},
	}
	for i := range testcases {
		if checkCommentInput(testcases[i].content, testcases[i].pagenum) != nil {
			if checkCommentInput(testcases[i].content, testcases[i].pagenum).Error() != testcases[i].err.Error() {
				t.Errorf("%s FAILED: content: %s  pagenum: %d, error expected: %v || error gotten: %v", testcases[i].testCaseName, testcases[i].content, testcases[i].pagenum, testcases[i].err, checkCommentInput(testcases[i].content, testcases[i].pagenum))
			}
		} else {
			if testcases[i].err != nil {
				t.Errorf("%s FAILED: content: %s  pagenum: %d, error expected: %v || error gotten: %v", testcases[i].testCaseName, testcases[i].content, testcases[i].pagenum, testcases[i].err, checkCommentInput(testcases[i].content, testcases[i].pagenum))
			}
		}
	}
}

// TestCreateComment tests CreateComment in fdfComment.go
func TestCreateComment(t *testing.T) {
	testcases := []struct {
		testCaseName string
		content      []byte
		pagenum      int
		err          error
		cmnt         fdfComment
	}{
		{"Inputs Okay", []byte("Test Note"), 23, nil, fdfComment{[]byte("Test Note"), 23}},
		{"bad page number", []byte("Test Note"), -23, errors.New("Create Comment Failed: content exists [true], page number not negative [false]"), fdfComment{[]byte("Test Note"), -23}},
		{"bad content", []byte(""), 23, errors.New("Create Comment Failed: content exists [false], page number not negative [true]"), fdfComment{[]byte(""), 23}},
		{"bad content and page number", []byte(""), -23, errors.New("Create Comment Failed: content exists [false], page number not negative [false]"), fdfComment{[]byte(""), -23}},
	}
	for i := range testcases {
		cmnt, err := NewComment(testcases[i].content, testcases[i].pagenum)
		if err != nil {
			if err.Error() != testcases[i].err.Error() {
				t.Errorf("%s FAILED", testcases[i].testCaseName)
			}
		} else {
			if (string(testcases[i].content) != string(cmnt.GetContent())) || (testcases[i].pagenum != cmnt.GetPageNumber()) {
				t.Errorf("%s FAILED", testcases[i].testCaseName)
			}
		}
	}
}

// TestfdfComment
func TestFdfComment(t *testing.T) {
	tText := "TestComment"
	cmnt, err := NewComment([]byte(tText), 3)
	if err != nil {
		t.Error("Failed, check create comment")
	}
	if string(cmnt.GetContent()) != tText {
		t.Error("GetContent Failed")
	}
	if cmnt.GetPageNumber() != 3 {
		t.Error("GetPageNum Failed")
	}
}
