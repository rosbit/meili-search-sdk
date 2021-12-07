package gssdk

import (
	"github.com/rosbit/gnet"
	"fmt"
	"net/http"
)

func AddIndexDoc(index string, doc map[string]interface{}) (updateId int, err error) {
	return updateIndexDocs("POST", index, []map[string]interface{}{doc})
}

func AddIndexDocs(index string, docs []map[string]interface{}) (updateId int, err error) {
	return updateIndexDocs("POST", index, docs)
}

func UpdateIndexDoc(index string, doc map[string]interface{}) (updateId int, err error) {
	return updateIndexDocs("PUT", index, []map[string]interface{}{doc})
}

func UpdateIndexDocs(index string, docs []map[string]interface{}) (updateId int, err error) {
	return updateIndexDocs("PUT", index, docs)
}


func updateIndexDocs(method, index string, docs []map[string]interface{}) (updateId int, err error) {
	url := fmt.Sprintf("%s/indexes/%s/documents", SearcherBaseUrl, index)
	var res struct {
		UpdateId int `json:"updateId"`
	}
	var status int
	if status, err = gnet.JSONCallJ(url, &res, gnet.M(method), gnet.Params(docs)); err != nil {
		return
	}
	if status != http.StatusAccepted {
		err = fmt.Errorf("status: %d", status)
		return
	}
	updateId = res.UpdateId
	return
}

func DeleteIndexDoc(index string, docId string) (updateId int, err error) {
	url := fmt.Sprintf("%s/indexes/%s/documents/%s", SearcherBaseUrl, index, docId)
	var res struct {
		UpdateId int `json:"updateId"`
	}
	var status int
	if status, err = gnet.HttpCallJ(url, &res, gnet.M("DELETE")); err != nil {
		return
	}
	if status != http.StatusAccepted {
		err = fmt.Errorf("status: %d", status)
	}
	updateId = res.UpdateId
	return
}

func DeleteIndexDocs(index string, docIds []string) (updateId int, err error) {
	url := fmt.Sprintf("%s/indexes/%s/documents/delete-batch", SearcherBaseUrl, index)
	var res struct {
		UpdateId int `json:"updateId"`
	}
	var status int
	if status, err = gnet.JSONCallJ(url, &res, gnet.Params(docIds)); err != nil {
		return
	}
	if status != http.StatusAccepted {
		err = fmt.Errorf("status: %d", status)
	}
	updateId = res.UpdateId
	return
}

func DeleteAllIndexDocs(index string) (updateId int, err error) {
	url := fmt.Sprintf("%s/indexes/%s/documents", SearcherBaseUrl, index)
	var res struct {
		UpdateId int `json:"updateId"`
	}
	var status int
	if status, err = gnet.HttpCallJ(url, &res, gnet.M("DELETE")); err != nil {
		return
	}
	if status != http.StatusAccepted {
		err = fmt.Errorf("status: %d", status)
	}
	updateId = res.UpdateId
	return
}
