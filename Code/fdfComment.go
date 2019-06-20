// Copyright 2019 Richard J. Cordes
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
