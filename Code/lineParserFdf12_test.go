package FdfExtract

import (
	"fmt"
	"testing"
)

// TestLineParserFdf12 tests lineParserFdf12.go
// NOTE: Not the best test cases, will have to be improved in the future **
// 			 Many of these were made in response to errors being received post-testing
//			 so it's a bit haphazard.
func TestLineParserFdf12(t *testing.T) {
	goodArray := []byte("<</CThis)/Page 2/RC( is a larger string Subtype/Contents(These are words within it)/Subtype")
	cmnt, err := lineParserFdf12(goodArray)
	if err != nil {
		t.Error(err.Error())
	}
	if string(cmnt.GetContent()) != "These are words within it" {
		t.Error("Incorrect String")
	}
	if cmnt.GetPageNumber() != 2 {
		t.Error("Incorrect Pagenumber")
	}
	goodArray2 := []byte("<</CThis)/Page 23/RC( is a larger string Subtype/Contents(These are words within it)/Subtype")
	cmnt2, err2 := lineParserFdf12(goodArray2)
	if err2 != nil {
		t.Error(err2.Error())
	}
	if string(cmnt2.GetContent()) != "These are words within it" {
		t.Error("Incorrect String")
	}
	if cmnt2.GetPageNumber() != 23 {
		t.Error("Incorrect Pagenumber")
	}
	badArray1 := []byte("<</mThis)/Page 23/RC( is a larger string Subtype/Contents(These are words within it)/Subtype")
	_, err3 := lineParserFdf12(badArray1)
	if err3 == nil {
		t.Error("failed to catch bad heading")
	}
	badArray2 := []byte("<</CThis)/Page 23/RC( is a larger string Subtype/Contens(These are words within it)/Subtype")
	_, err4 := lineParserFdf12(badArray2)
	if err4 == nil {
		t.Error("failed to produce error in response to badly formatted content")
	}
	badArray3 := []byte("<</CThis)/Page 23/C( is a larger string Subtype/Contents(These are words within it)/Subtype")
	_, err5 := lineParserFdf12(badArray3)
	if err5.Error() != "Page Number unable to be retrieved" {
		fmt.Println(err5.Error())
		t.Error("failed to produce error in response to badly formatted page number")
	}
	badArray4 := []byte("<</CThis)/Page 2a3/RC( is a larger string Subtype/Contents(These are words within it)/Subtype")
	_, err6 := lineParserFdf12(badArray4)
	if err6.Error() != "Page Number unable to be converted to int" {
		fmt.Println(err5.Error())
		t.Error("failed to produce error in response to badly formatted page number")
	}
	badArray5 := []byte("<</CThis)/Page 23/RC( is a larger string Subtype/Contents()/Subtype")
	_, err7 := lineParserFdf12(badArray5)
	if err7 == nil {
		t.Error("failed to produce error in response to empty comment")
	}
	badArray6 := []byte("<</C")
	_, err8 := lineParserFdf12(badArray6)
	if err8 == nil {
		t.Error("failed to produce error in response to line which was too small to produce content")
	}
}
