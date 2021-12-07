package gssdk

import (
	"github.com/rosbit/gnet"
	"fmt"
)

func SetSearchableFields(index string, fields []string) (updateId int, err error) {
	return setSttting("searchableAttributes", index, fields)
}

func SetDisplayedFields(index string, fields []string) (updateId int, err error) {
	return setSttting("displayedAttributes", index, fields)
}

func SetFilterableFields(index string, fields []string) (updateId int, err error) {
	return setSttting("filterableAttributes", index, fields)
}

func SetSortableFields(index string, fields []string) (updateId int, err error) {
	return setSttting("sortableAttributes", index, fields)
}

func SetRankingRules(index string, rules []string) (updateId int, err error) {
	return setSttting("rankingRules", index, rules)
}

func SetStopWords(index string, words []string) (updateId int, err error) {
	return setSttting("stopWords", index, words)
}

func SetDistinctField(index string, field string) (updateId int, err error) {
	return setSttting("distinctAttribute", index, field)
}

func SetSynoyms(index string, synoyms map[string][]string) (updateId int, err error) {
	return setSttting("synonyms", index, synoyms)
}

func setSttting(attrName string, index string, attrValue interface{}) (updateId int, err error) {
	url := fmt.Sprintf("%s/indexes/%s/settings", SearcherBaseUrl, index)
	var res struct {
		UpdateId int `json:"updateId"`
	}
	_, err = gnet.JSONCallJ(url, &res, gnet.Params(map[string]interface{}{
		attrName: attrValue,
	}))
	updateId = res.UpdateId
	return
}

func GetSettings(index string) (attrs map[string]interface{}, err error) {
	url := fmt.Sprintf("%s/indexes/%s/settings", SearcherBaseUrl, index)
	_, err = gnet.HttpCallJ(url, &attrs)
	return
}

func ResetSettings(index string) (updateId int, err error) {
	url := fmt.Sprintf("%s/indexes/%s/settings", SearcherBaseUrl, index)
	var res struct {
		UpdateId int `json:"udpateId"`
	}
	if _, err = gnet.HttpCallJ(url, &res, gnet.M("DELETE")); err != nil {
		return
	}
	updateId = res.UpdateId
	return
}
