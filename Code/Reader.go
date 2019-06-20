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

// A Reader takes a filepath to an FDF document and outputs a CommentBlock
type Reader interface {
	Read(string) (CommentBlock, error) // takes filepath and returns CommentBlock
	enableConcurrency()                //set FdfReader to be Concurrent
	disableConcurrency()               //set FdfReader to be not be Concurrent
}
