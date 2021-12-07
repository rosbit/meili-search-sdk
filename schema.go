package gssdk

import (
	"github.com/rosbit/gnet"
	"fmt"
	"time"
	"net/http"
)

type Schema struct {
	Uid string `json:"uid"`
	PrimaryKey string `json:"primaryKey"`
	CreateAt time.Time `json:"createdAt"`
	UpdateAt time.Time `json:"updatedAt"`
}

func CreateSchema(index string, pk string) (schema *Schema, err error) {
	var res Schema

	url := fmt.Sprintf("%s/indexes", SearcherBaseUrl)
	var status int
	if status, err = gnet.JSONCallJ(url, &res, gnet.Params(map[string]interface{}{
		"uid": index,
		"primaryKey": pk,
	})); err != nil {
		return
	}
	if status != http.StatusCreated {
		err = fmt.Errorf("status: %d", status)
		return
	}
	schema = &res
	return
}

func DeleteSchema(index string) (err error) {
	url := fmt.Sprintf("%s/indexes/%s", SearcherBaseUrl, index)
	var res struct {}
	var status int
	if status, err = gnet.HttpCallJ(url, &res, gnet.M("DELETE")); err != nil {
		return
	}
	if status != http.StatusNoContent {
		err = fmt.Errorf("status: %d", status)
		return
	}
	return
}

func GetSchema(index string) (schema *Schema, err error) {
	url := fmt.Sprintf("%s/indexes/%s", SearcherBaseUrl, index)
	var res Schema
	var status int
	if status, err = gnet.HttpCallJ(url, &res, gnet.DontReadRespBody()); err != nil {
		return
	}

	if status != http.StatusOK {
		err = fmt.Errorf("status: %d", status)
		return
	}
	schema = &res
	return
}

func UpdateSchema(index string, pk string) (schema *Schema, err error) {
	url := fmt.Sprintf("%s/indexes/%s", SearcherBaseUrl, index)
	var res Schema
	var status int
	if status, err = gnet.JSONCallJ(url, &res, gnet.M("PUT"), gnet.Params(map[string]interface{}{
		"primaryKey": pk,
	})); err != nil {
		return
	}
	if status != http.StatusOK {
		err = fmt.Errorf("status: %d", status)
		return
	}
	schema = &res
	return
}
