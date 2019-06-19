package FdfExtract

// fdfCommentBlock contains comments pulled from FCE
type fdfCommentBlock struct {
	sourceFile []byte
	comments   []Comment
}

// GetSourceFile
func (cb fdfCommentBlock) GetSourceFile() []byte {
	return cb.sourceFile
}

// GetComments
func (cb fdfCommentBlock) GetComments() []Comment {
	return cb.comments
}

//GetComments

// NewCommentBlock creates a new comment block
func NewCommentBlock(sourceFile []byte, comments []Comment) CommentBlock {
	return fdfCommentBlock{sourceFile, comments}
}
