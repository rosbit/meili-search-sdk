package gssdk

import (
	"github.com/rosbit/gnet"
	"fmt"
	"log"
	// "os"
	"net/http"
)

func Search(index string, q string, options ...Option) (docs []Doc, resultHeader *ResultHeader, err error) {
	if len(q) == 0 {
		err = fmt.Errorf("param q expeted")
		return
	}
	query := getQueryOptions(q, options...)
	params := query.makeQuery()
	log.Printf("[info] params: %v\n", params)
	url := fmt.Sprintf("%s/indexes/%s/search", SearcherBaseUrl, index)

	var res struct {
		*ResultHeader
		Docs []Doc `json:"hits"`
	}
	var status int
	if status, err = gnet.JSONCallJ(url, &res, gnet.Params(params)/*, gnet.BodyLogger(os.Stderr)*/); err != nil {
		return
	}
	if status != http.StatusOK {
		err = fmt.Errorf("status: %d", status)
		return
	}
	docs = res.Docs
	resultHeader = res.ResultHeader
	return
}
