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
	"strconv"
	"strings"
)

// lineParser12 is used on fdf docs with heading "%FDF-1.2"
func lineParserFdf12(line []byte) (Comment, error) {
	if len(line) < 5 {
		return nil, errors.New("Empty Line") // nothing on this line
	}
	if !checkbytes(line[0:4], []byte("<</C")) {
		return nil, errors.New("Bad Heading") // line starter incorrect
	}
	delimiters := []struct { // for readability
		sWord []byte
		eWord []byte
	}{
		{[]byte("/Contents("), []byte(")/Subtype")},
		{[]byte("/Page "), []byte("/RC(")},
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
	cmnt, err := NewComment(copyBytes(cleanformatting(content)), pagenumber)
	if err != nil {
		return nil, err
	}
	return cmnt, nil
}

func cleanformatting(bytes []byte) []byte {
	s := string(bytes)
	s = strings.Replace(s, string([]byte{92, 114, 92, 110}), string("\n"), -1)
	return copyBytes([]byte(s))
}
