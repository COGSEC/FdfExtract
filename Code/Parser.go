package FdfExtract

// Parser functions take a potential annotation line and output a Comment
type Parser func([]byte) Comment
