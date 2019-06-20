package FdfExtract

// A Reader takes a filepath to an FDF document and outputs a CommentBlock
type Reader interface {
	Read(string) (CommentBlock, error) // takes filepath and returns CommentBlock
	enableConcurrency()                //set FdfReader to be Concurrent
	disableConcurrency()               //set FdfReader to be not be Concurrent
}
