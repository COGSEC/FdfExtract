package FdfExtract

//Comment has methods to pull annotation info found within a foxit document
type Comment interface {
	GetContent() []byte
	GetPageNumber() int
}
