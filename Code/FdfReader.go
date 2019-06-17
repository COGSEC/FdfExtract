package FdfExtract

// FdfReader can read bytes in and output the number of bytes and an error if
// applicable
type FdfReader interface {
	ReadFdf([]byte) (int, error)
}
