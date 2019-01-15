package query

import (
	"encoding/json"
	"strings"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
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
	Count       uint64 `db:"c,pk,omitempty"`
	CurrentPage uint
	NextPage    uint
	PrevPage    uint
	FirstPage   uint
	LastPage    uint
}

type FieldSet struct {
	fields map[string]string
}

// Makes an empty fieldSet
func GetFieldSet() FieldSet {
	set := FieldSet{}
	set.fields = make(map[string]string)
	return set
}

// add something that matches your select query:  tag.id, id
func (set FieldSet) AddField(dbField string, structField string) {
	set.fields[structField] = dbField
	return
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

// Query Options for select statements
// fields, sort, limit, page, sind implementiert
// mit der dbMeta kann man sich eine Pagination bauen...
func ApplyRequestOptionsToSelect(q sqlbuilder.Selector, fieldSet FieldSet, options QueryOptions) (sqlbuilder.Selector, DBMeta) {
	type cnt struct {
		Count uint8 `db:"c,pk,omitempty"`
	}

	metaQ := q

	var requestedFields []interface{}
	if options.Fields != "" {
		fields := strings.Split(options.Fields, ",")

		for _, field := range fields {
			requestedFields = append(requestedFields, fieldSet.fields[field]+" as "+field)
		}
	} else {
		//append complete fieldset
		for field, dbfield := range fieldSet.fields {
			requestedFields = append(requestedFields, dbfield+" as "+field)
		}
	}

	q = q.Columns(requestedFields...)

	q = q.OrderBy(options.Sort)

	limit := int(PaginationDefault)
	if options.Limit != 0 {
		limit = int(options.Limit)
	}
	q = q.Limit(limit)
	if options.Page != 0 {
		q = q.Offset(limit*int(options.Page) - 1)
	}

	if options.Sort != "" {
		q = q.OrderBy(options.Sort)
	}

	// build meta
	metaQ = metaQ.Columns(db.Raw("COUNT(*) as c"))
	var meta DBMeta
	err := metaQ.One(&meta)

	meta.CurrentPage = 1
	if options.Page > 1 {
		meta.CurrentPage = options.Page
	}
	if err == nil {
		if 0 == limit {
			limit = 1
		}
		meta.LastPage = uint(int(meta.Count) / limit)
		meta.FirstPage = 1
		if meta.CurrentPage < meta.LastPage {
			meta.NextPage = meta.CurrentPage + 1
		}
		if meta.CurrentPage > 1 {
			meta.PrevPage = meta.CurrentPage - 1
		}
	}

	return q, meta
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
