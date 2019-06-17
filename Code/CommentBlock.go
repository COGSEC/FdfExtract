package FdfExtract

// CommentBlock contains meta-data and comments pulled from an FCE
type CommentBlock interface {
	GetSourceFilePath() []byte
	GetComments() []Comment
}
