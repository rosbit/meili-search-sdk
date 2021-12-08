package gssdk

import (
	"strings"
	"fmt"
)

type query struct {
	q string
	sorting []string
	f []string
	page, pageSize int
	fl []string
}

type Option func(*query)
func getQueryOptions(qStr string, options ...Option) *query {
	q := &query{q: qStr}
    for _, o := range options {
        o(q)
    }
    return q
}
func (q *query) makeQuery() map[string]interface{} {
	res := make(map[string]interface{})
	if len(q.q) > 0 {
		res["q"] = q.q
	}
	if q.pageSize <= 0 || q.pageSize > 100 {
		q.pageSize = 20
	}
	res["limit"] = q.pageSize
	if q.page > 0 {
		res["offset"] = (q.page-1) * q.pageSize
	} else {
		q.page = 1
	}
	if len(q.f) > 0 {
		res["filter"] = strings.Join(q.f, " AND ")
	}
	if len(q.sorting) > 0 {
		res["sort"] = q.sorting
	}
	if len(q.fl) > 0 {
		res["attributesToRetrieve"] = q.fl
	}
	return res
}

func Sorting(field string, isAsc ...bool) Option {
	return func(q *query) {
		scend := func()string{
			if len(isAsc) > 0 && isAsc[0]  {
				return "asc"
			}
			return "desc"
		}()
		q.sorting = append(q.sorting, fmt.Sprintf("%s:%s", field, scend))
	}
}

func Filter(field string, vals []string) Option {
	return func(q *query) {
		if len(vals) == 0 {
			return
		}
		q.f = append(q.f, fmt.Sprintf("%s:%s", field, strings.Join(vals, ",")))
	}
}

func Page(page int) Option {
	return func(q *query) {
		q.page = page
	}
}

func PageSize(pageSize int) Option {
	return func(q *query) {
		q.pageSize = pageSize
	}
}

func OutputFields(fieldNames []string) Option {
	return func(q *query) {
		if len(fieldNames) > 0 {
			q.fl = append(q.fl, fieldNames...)
		}
	}
}
