package FdfExtract

import (
	"errors"
	"strconv"
)

// lineParser12 is used on fdf docs with heading "%FDF-1.2"
func lineParserFdf12(line []byte) (Comment, error) {
	if len(line) < 5 {
		return nil, errors.New("Empty Line") // nothing on this line
	}
	if !checkbytes(line[0:4], []byte("<</C")) {
		return nil, errors.New("Bad Heading") // line starter incorrect
	}
	delimiters := []struct {
		sWord []byte
		eWord []byte
	}{
		{[]byte("/Contents("), []byte(")/Subtype")},
		{[]byte(")/Page "), []byte("/RC(")},
	}
	content, ok := extract(line, delimiters[0].sWord, delimiters[0].eWord)
	if !ok {
		return nil, errors.New("Content unable to be retrieved")
	}
	pagenum, ok := extract(line, delimiters[1].sWord, delimiters[1].eWord)
	if !ok {
		return nil, errors.New("Page Number unable to be retrieved")
	}
	pagenumber, err := strconv.Atoi(string(pagenum))
	if err != nil {
		return nil, errors.New("Page Number unable to be converted to int")
	}
	cmnt, err := NewComment(content, pagenumber)
	if err != nil {
		return nil, err
	}
	return cmnt, nil
}
