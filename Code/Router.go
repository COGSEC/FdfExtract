package FdfExtract

// Router must be able to identify the fdf format and pair it to an appropriate
// parse function
type Router interface {
	PickReader([]byte, map[string]Parser) Parser
}
