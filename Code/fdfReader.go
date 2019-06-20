package FdfExtract

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

//fdfReader is the default FdfReader
type fdfReader struct {
	concurrency bool                  // enables/disables concurrency within the reader
	parsers     map[string]lineParser // maps doc headings to lineParsers
}

//EnableConcurrency enables concurrency in the Reader
func (reader *fdfReader) enableConcurrency() {
	reader.concurrency = true
}

//DisableConcurrency enables concurrency in the Reader
func (reader *fdfReader) disableConcurrency() {
	reader.concurrency = false
}

// Read a file from a file path and output a CommentBlock
func (reader *fdfReader) Read(path string) (CommentBlock, error) {
	if err := checkpath(path); err != nil { // check path for correct extention
		return nil, err // if incorrect extension return nil CommentBlock and err
	}
	fdf, err := os.Open(path) // if correct extension, open file at given path
	if err != nil {
		return nil, err // if error, return nil CommentBlock and err
	}
	defer fdf.Close()                             //defer file close
	scanner := bufio.NewScanner(fdf)              // create new scanner
	scanner.Scan()                                // scan in for heading
	parser, err := reader.router(scanner.Bytes()) // scan first line for doc-heading
	if err != nil {
		return nil, err // if heading not unrecognised, return nil CommentBlock and err
	}
	var comments []Comment // create comment array
	comments = readerNonConcurrent(scanner, parser)
	for i := range comments {
		fmt.Println(string(comments[i].GetContent()))
	}
	var output []Comment
	for i := range comments {
		output = append(output, comments[i])
	}
	return NewCommentBlock([]byte(path), output), nil
}

//router uses a doc heading to pick a LineParser
func (reader *fdfReader) router(firstLine []byte) (lineParser, error) {
	if _, ok := reader.parsers[string(firstLine)]; ok {
		return reader.parsers[string(firstLine)], nil
	}
	return nil, errors.New("Fdf Doc-Heading unrecognised")
}

// readerNonConcurrent is a non-concurrent version of fdfReader
func readerNonConcurrent(scanner *bufio.Scanner, parser lineParser) []Comment {
	var comments []Comment // create comment array
	for scanner.Scan() {
		cmnt, err := parser(copyBytes(scanner.Bytes()))
		if err == nil {
			comments = append(comments, cmnt)
		}
	}
	return comments
}

// NewReader returns a FdfReader Interface, which is the primary exported type
func NewReader() Reader {
	// init map
	var reader fdfReader
	reader.concurrency = false
	reader.parsers = make(map[string]lineParser) // instantiate map
	reader.parsers["%FDF-1.2"] = lineParserFdf12 // init attachment to lineParserFdf12 for docheading %FDF-1.2
	// IF ADDITIONAL LINEPARSERS BECOME AVAILABLE, INIT THEM HERE ^^^ //
	return &reader
}
