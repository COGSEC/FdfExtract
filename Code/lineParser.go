package FdfExtract

import (
	"golang.org/x/text/language"
	"golang.org/x/text/search"
)

// LineParser functions takes bytes, parse them, and returns a Comment
type lineParser func([]byte) (Comment, bool)

/*
The following are utility functions expected to be held in common between many
types of lineParsers
*/

//check bytes compares two sets of bytes for consistency
func checkbytes(bytes []byte, check []byte) bool {
	if len(bytes) != len(check) { // if bytes unequal length, cannot be identical
		return false
	}
	for i := range bytes {
		if bytes[i] != check[i] {
			return false
		}
	}
	return true
}

// posVector is a simple structure which is intended to contain the start and
// end indexes of a string
type posVector struct {
	startPos int
	endPos   int
}

// searchforbytes searches an array of bytes for an array of bytes and returns
// the idx at which it was found and whether or not it was found. If it was not
// found it will return -1.
func searchforbytes(array []byte, term []byte) (posVector, bool) {
	s := search.New(language.English, search.IgnoreCase)
	idx, idx2 := s.Index(array, term)
	if idx == -1 {
		return posVector{idx, idx2}, false
	}
	return posVector{idx, idx2}, true
}

// extract pulls data from between two phrases within an array of bytes
func extract(array []byte, sWord []byte, eWord []byte) ([]byte, bool) {
	sPos, ok := searchforbytes(array, sWord)
	if !ok {
		return nil, false // comment formatted badly
	}
	ePos, ok := searchforbytes(array, eWord)
	if !ok {
		return nil, false // comment formatted badly
	}
	return array[sPos.endPos:ePos.startPos], true
}
