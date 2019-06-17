package FdfExtract

//Comment is an annotation found within a foxit document
type Comment interface {
	GetNote() []byte
	GetPageNumber() int
}
