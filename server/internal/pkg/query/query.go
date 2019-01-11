package query

import (
	"encoding/json"
	"strings"
	"upper.io/db.v3"
)

// Anfrageoptionen für upper
type QueryOptions struct {
	Fields  string `json:"fields,omitempty"` // partial representation
	Sort    string `json:"sort,omitempty"`
	Filter  string `json:"filter,omitempty"`
	Count   bool   `json:"count,omitempty"` // return count in meta
	Sum     string `json:"sum,omitempty"`   // calculate sum
	Context string `json:"context,omitempty"`
	Limit   uint   `json:"limit,omitempty"`  // set page limit to limit
	Page    uint   `json:"page,omitempty"`   // pagination
	Cursor  uint   `json:"cursor,omitempty"` // for cursor pagination
}

type DBMeta struct {
	Count       uint64
	CurrentPage uint
	NextPage    uint
	PrevPage    uint
	FirstPage   uint
	LastPage    uint
}

var PaginationDefault uint

func init() {
	PaginationDefault = 23
}

// Optionen für Listenelemente kommen aus dem proto als beliebiger Typ daher, jedoch immer in der gleichen nummerierung
// diese werden in die QueryOptions Form gebracht, damit upper sauber damit umgehen kann.
func GetListOptionsFromRequest(options interface{}) QueryOptions {
	tmp, _ := json.Marshal(options)
	var opts QueryOptions
	json.Unmarshal(tmp, &opts)
	return opts
}

// Query Options für auf das db.Result anwenden.
// fields, sort, limit, page, sind implementiert
// mit der dbMeta kann man sich eine Pagination bauen...
func ApplyRequestOptionsToQuery(res db.Result, options QueryOptions) (db.Result, DBMeta) {
	var meta DBMeta
	if options.Limit != 0 {
		res = res.Paginate(options.Limit)
	} else {
		res = res.Paginate(PaginationDefault)
	}

	if options.Fields != "" {
		fields := strings.Split(options.Fields, ",")
		s := make([]interface{}, len(fields))
		for i, field := range fields {
			s[i] = field
		}
		res = res.Select(s...)
	}

	if options.Sort != "" {
		res = res.OrderBy(options.Sort)
	}

	meta.CurrentPage = 1
	if options.Page > 0 {
		meta.CurrentPage = options.Page
		res = res.Page(options.Page)
	}
	pages, _ := res.TotalPages()
	meta.LastPage = pages
	meta.FirstPage = 1
	if meta.CurrentPage < meta.LastPage {
		meta.NextPage = meta.CurrentPage + 1
	}
	if meta.CurrentPage > 1 {
		meta.PrevPage = meta.CurrentPage - 1
	}

	if options.Count {
		meta.Count, _ = res.Count()
	}

	return res, meta
}
