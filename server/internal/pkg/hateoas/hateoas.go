package hateoas

import (
	"../query"
	"context"
	"github.com/veith/protos/rest"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
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
func GenerateCollectionHATEOAS(ctx context.Context, path string, dbMeta query.DBMeta) Hateoas {
	md, _ := metadata.FromIncomingContext(ctx)
	baseURL := strings.Join(md["api-base-url"], "")

	var h Hateoas
	h.AddLink("self", "application/json", baseURL+path+"?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage), 10), rest.Link_GET)
	if dbMeta.PrevPage != 0 {
		h.AddLink("prev", "application/json", baseURL+path+"?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage-1), 10), rest.Link_GET)
	}
	if dbMeta.NextPage != 0 {
		h.AddLink("next", "application/json", baseURL+path+"?page="+strconv.FormatUint(uint64(dbMeta.CurrentPage+1), 10), rest.Link_GET)
	}
	h.AddLink("first", "application/json", baseURL+path+"?page="+strconv.FormatUint(uint64(dbMeta.FirstPage), 10), rest.Link_GET)
	h.AddLink("last", "application/json", baseURL+path+"?page="+strconv.FormatUint(uint64(dbMeta.LastPage), 10), rest.Link_GET)
	h.AddLink("create", "application/json", baseURL+path, rest.Link_POST)
	return h
}

func GenerateEntityHateoas(ctx context.Context, path string, id string) Hateoas {
	//todo check gegen spec machen
	md, _ := metadata.FromIncomingContext(ctx)
	baseURL := strings.Join(md["api-base-url"], "")

	var h Hateoas
	h.AddLink("self", "application/json", baseURL+path+"/"+id, rest.Link_GET)
	h.AddLink("delete", "application/json", baseURL+path+"/"+id, rest.Link_DELETE)
	h.AddLink("update", "application/json", baseURL+path+"/"+id, rest.Link_PATCH)
	h.AddLink("parent", "application/json", baseURL+path, rest.Link_GET)
	h.AddLink("complete", "application/json", baseURL+path+"/"+id+":complete", rest.Link_POST)
	return h
}
