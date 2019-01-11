package hateoas

import (
	"../query"
	"github.com/veith/protos/rest"
	"strconv"
)

type Hateoas struct {
	Links []*rest.Link
}

// links einem HTS hinzufügen
func (h *Hateoas) AddLink(rel, contenttype, href string, method rest.Link_Method) {
	link := rest.Link{Rel: rel, Href: href, Type: contenttype, Method: method}
	h.Links = append(h.Links, &link)
}

// hateoas anhand DBMEta für eine Collection erzeugen
func GenerateCollectionHATEOAS(dbMeta query.DBMeta) Hateoas {
	//todo Link_Get,.. nach REST schieben
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage), 10), rest.Link_GET)
	if dbMeta.PrevPage != 0 {
		h.AddLink("prev", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage-1), 10), rest.Link_GET)
	}
	if dbMeta.NextPage != 0 {
		h.AddLink("next", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage+1), 10), rest.Link_GET)
	}
	h.AddLink("first", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.FirstPage), 10), rest.Link_GET)
	h.AddLink("last", "application/json", "http://localhost:8888/tasks?page="+strconv.FormatUint(uint64(dbMeta.LastPage), 10), rest.Link_GET)
	h.AddLink("create", "application/json", "http://localhost:8888/tasks", rest.Link_POST)
	return h
}

func GenerateEntityHateoas(id string) Hateoas {
	//todo check gegen spec machen
	var h Hateoas
	h.AddLink("self", "application/json", "http://localhost:8888/tasks/"+id, rest.Link_GET)
	h.AddLink("delete", "application/json", "http://localhost:8888/tasks/"+id, rest.Link_DELETE)
	h.AddLink("update", "application/json", "http://localhost:8888/tasks/"+id, rest.Link_PATCH)
	h.AddLink("parent", "application/json", "http://localhost:8888/tasks", rest.Link_GET)
	h.AddLink("complete", "application/json", "http://localhost:8888/tasks"+id+":complete", rest.Link_POST)
	return h
}
