package main

import (
	gssdk "github.com/rosbit/meili-search-sdk"
	"flag"
	"os"
	"fmt"
	"strings"
	"encoding/json"
)

func init() {
	searcherUrl := os.Getenv("SEARCHER_URL")
	if len(searcherUrl) == 0 {
		panic("env SEARCHER_URL not found")
	}
	gssdk.SetSearcherBaseUrl(searcherUrl)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s <command> <options>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Where <command> are:\n")
		fmt.Fprintf(os.Stderr, " create-schema\n")
		fmt.Fprintf(os.Stderr, " delete-schema, dump-schema\n")
		fmt.Fprintf(os.Stderr, " add-index-doc update-index-doc\n")
		fmt.Fprintf(os.Stderr, " delete-index-doc\n")
		fmt.Fprintf(os.Stderr, " search\n")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create-schema":
		createSchema()
	case "delete-schema":
		deleteSchema()
	case "dump-schema":
		dumpSchema()
	case "add-index-doc":
		addIndexDoc()
	case "update-index-doc":
		updateIndexDoc()
	case "delete-index-doc":
		deleteIndexDoc()
	case "search":
		search()
	default:
		fmt.Fprintf(os.Stderr, "unknown command %s\n", os.Args[1])
		os.Exit(2)
	}
}

func createSchema() {
	f := flag.NewFlagSet("create-schema", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	pk := f.String("pk", "", "specify primary key field name")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*pk) == 0 {
		fmt.Fprintf(os.Stderr, "%s create-schema -index=xxx -pk=xxx\n", os.Args[0])
		os.Exit(4)
	}
	schema, err := gssdk.CreateSchema(*index, *pk)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("schema: %#v\n", schema)
	}
}

func deleteSchema() {
	f := flag.NewFlagSet("delete-schema", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 {
		fmt.Fprintf(os.Stderr, "%s delete-schema -index=xxx\n", os.Args[0])
		os.Exit(4)
	}
	err := gssdk.DeleteSchema(*index)
	fmt.Printf("err: %v\n", err)
}

func dumpSchema() {
	f := flag.NewFlagSet("dump-schema", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 {
		fmt.Fprintf(os.Stderr, "%s dump-schema -index=xxx\n", os.Args[0])
		os.Exit(4)
	}
	schema, err := gssdk.GetSchema(*index)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(5)
	}
	fmt.Printf("%#v\n", schema)
}

func addIndexDoc() {
	f := flag.NewFlagSet("add-index-doc", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	doc := f.String("doc", "", "specify doc with A JSON object")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*doc) == 0 {
		fmt.Fprintf(os.Stderr, "%s add-index-doc -index=xxx -doc={JSON}\n", os.Args[0])
		os.Exit(4)
	}
	var d map[string]interface{}
	if err := json.Unmarshal([]byte(*doc), &d); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	updateId, err := gssdk.AddIndexDoc(*index, d)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(6)
	}
	fmt.Printf("updateId: %d\n", updateId)
}

func updateIndexDoc() {
	f := flag.NewFlagSet("update-index-doc", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	doc := f.String("doc", "", "specify doc with A JSON object")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*doc) == 0 {
		fmt.Fprintf(os.Stderr, "%s update-index-doc -index=xxx -doc={JSON}\n", os.Args[0])
		os.Exit(4)
	}
	var d map[string]interface{}
	if err := json.Unmarshal([]byte(*doc), &d); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(5)
	}
	updateId, err := gssdk.UpdateIndexDoc(*index, d)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(6)
	}
	fmt.Printf("updateId: %d\n", updateId)
}

func deleteIndexDoc() {
	f := flag.NewFlagSet("delete-index-doc", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	docId := f.String("doc-id", "", "specify docId returned when calling add/update-index-doc")
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0 || len(*docId) == 0 {
		fmt.Fprintf(os.Stderr, "%s delete-index-doc -index=xxx -doc-id=xxx\n", os.Args[0])
		os.Exit(4)
	}
	updateId, err := gssdk.DeleteIndexDoc(*index, *docId)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("updateId: %d\n", updateId)
	}
}

func search() {
	var options []gssdk.Option

	f := flag.NewFlagSet("search", flag.ExitOnError)
	index := f.String("index", "", "specify index name")
	q := f.String("q", "", "specify query string")
	f.Func("s", "specify sorting, format: field[:asc]", func(s string)error{
		if len(s) == 0 {
			return nil
		}
		ss := strings.Split(s, ":")
		switch len(ss) {
		case 0: break
		case 1: options = append(options, gssdk.Sorting(ss[0]))
		default:
			scend := ss[1] == "asc"
			options = append(options, gssdk.Sorting(ss[0], scend))
		}
		return nil
	})
	offset := f.Int("offset", 0, "specify offset")
	limit := f.Int("limit", 0, "specify limit")
	f.Func("f", "specify field filter, format: field:val1,val2,...", func(filter string)error{
		if len(filter) == 0 {
			return nil
		}
		fs := strings.Split(filter, ":")
		switch len(fs) {
		case 2:
			vs := strings.FieldsFunc(fs[1], func(c rune)bool{
				return c == ',' || c == ' ' || c == '\t'
			})
			if len(vs) == 0 {
				return nil
			}
			options = append(options, gssdk.Filter(fs[0], vs))
		default:
		}
		return nil
	})
	f.Func("fl", "specify field list", func(fl string)error{
		if len(fl) == 0 {
			return nil
		}
		fls := strings.FieldsFunc(fl, func(c rune)bool{
			return c == ',' || c == ' ' || c == '\t'
		})
		if len(fls) == 0 {
			return nil
		}
		options = append(options, gssdk.OutputFields(fls))
		return nil
	})
	if err := f.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(3)
	}
	if len(*index) == 0  || len(*q) == 0 {
		fmt.Fprintf(os.Stderr, "%s search -index=xxx -q=xx -s=xx -offset=xx -limit=xx -fl=xxx\n", os.Args[0])
		os.Exit(4)
	}
	if *offset> 0 {
		options = append(options, gssdk.Offset(*offset))
	}
	if *limit > 0 {
		options = append(options, gssdk.Limit(*limit))
	}

	docs, header, err := gssdk.Search(*index, *q, options...)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(5)
	}
	jEnc := json.NewEncoder(os.Stdout)
	jEnc.Encode(header)
	jEnc.Encode(docs)
}
