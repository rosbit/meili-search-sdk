package gssdk

import (
	"encoding/json"
)

type Pagination struct {
	Total uint64
	Pages uint64
	PageSize int
	CurrPageNo uint64
	DocsInPage int
}

type Doc = json.RawMessage
