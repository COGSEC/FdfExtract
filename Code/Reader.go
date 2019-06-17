package FdfExtract

// FdfReader takes a filepath to an FDF document and outputs a CommentBlock
type FdfReader interface {
	ReadFdf(string) CommentBlock // takes filepath and returns CommentBlock
	SetConcurrency(bool)         //set FdfReader to be Concurrent
}
