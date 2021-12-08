package gssdk

type Pagination struct {
	Total uint64
	Pages uint64
	PageSize int
	CurrPageNo uint64
	DocsInPage int
}

type Doc map[string]interface{}
