package FdfExtract

// fdfCommentBlock contains comments pulled from FCE
type fdfCommentBlock struct {
	sourceFile []byte
	comments   []Comment
}

// GetSourceFile from fdfCommentBlock
func (cb fdfCommentBlock) GetSourceFile() []byte {
	return cb.sourceFile
}

// GetComments from fdfCommentBlock
func (cb fdfCommentBlock) GetComments() []Comment {
	return cb.comments
}

// NewCommentBlock creates a new comment block
func NewCommentBlock(sourceFile []byte, comments []Comment) CommentBlock {
	return fdfCommentBlock{sourceFile, comments}
}
