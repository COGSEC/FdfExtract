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

// CommentBlock includes methods to get comments pulled from an FCE and return
// the path they were received from
type CommentBlock interface {
	GetSourceFile() []byte
	GetComments() []Comment
}
