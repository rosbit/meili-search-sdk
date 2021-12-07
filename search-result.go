package gssdk

type ResultHeader struct {
	Total uint64 `json:"nbHits"`
	Offset uint64 `json:"offset"`
	Limit int `json:"limit"`
	Exhaustive bool `json:"exhaustiveNbHits"`
	SearchTimeMs int `json:"processingTimeMs"`
	Query string `json:"query"`
}

type Doc map[string]interface{}
