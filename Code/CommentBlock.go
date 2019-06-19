package FdfExtract

// CommentBlock includes methods to get comments pulled from an FCE and return
// the path they were received from
type CommentBlock interface {
	GetSourceFile() []byte
	GetComments() []Comment
}
