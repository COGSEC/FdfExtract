package FdfExtract

import (
	"errors"
	"fmt"
)

// fdfComment is a struct which contains comment information
type fdfComment struct {
	content []byte
	page    int
}

// GetContent retrieves the content from FdfComment
func (comment fdfComment) GetContent() []byte {
	return comment.content
}

// GetPageNumber retrieves the page number from FdfComment
func (comment fdfComment) GetPageNumber() int {
	return comment.page
}

// NewComment returns an fdfComment as an exported Comment interface
func NewComment(content []byte, pagenumber int) (Comment, error) {
	err := checkCommentInput(content, pagenumber)
	if err != nil {
		return nil, err
	}
	return fdfComment{content, pagenumber}, err
}

//checkCommentInput checks content and pagenumber to ensure valid comments
func checkCommentInput(content []byte, pagenumber int) error {
	check := struct {
		contentExists         bool
		pagenumberNotNegative bool
	}{false, false}
	check.contentExists = (len(content) > 0)
	check.pagenumberNotNegative = pagenumber >= 0
	if check.contentExists && check.pagenumberNotNegative {
		return nil
	}
	errorString := fmt.Sprintf("Create Comment Failed: content exists [%t], page number not negative [%t]", check.contentExists, check.pagenumberNotNegative)
	return errors.New(errorString) // using fmt.sprintf instead of fmt.errorf for clarity
}
