package engine

type Request struct {
	Url string
	ParserFunc func(contents []byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items 	 []interface{}
}