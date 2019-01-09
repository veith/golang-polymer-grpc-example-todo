package task

import (
	"../../../proto/task"
	"encoding/json"
	"github.com/gogo/protobuf/types"
	"github.com/oklog/ulid"
	"github.com/veith/protos/date"
	"github.com/veith/protos/rest"
	//"google.golang.org/genproto/googleapis/type/date"
	"strconv"
)

type Hateoas struct {
	Links []*rest.Link
}

func MapTaskToProtoTask(ob1 *Task) *task.Task {
	var t types.Timestamp
	var date date.Date
	var q struct{}
	ob2 := task.Task{ob1.Id.String(), ob1.Title, ob1.Description, task.Complete(ob1.Completed), &date, &t, &t, q, []byte{}, 0}
	return &ob2
}
func MapProtoTaskToTask(ob1 *task.Task) *Task {
	id, _ := ulid.Parse(ob1.Id)
	ob2 := Task{id, ob1.Title, ob1.Description, int32(ob1.Completed)}
	return &ob2
}

// links einem HTS hinzufügen
func (h *Hateoas) AddLink(rel, contenttype, href string, method rest.Link_Method) {
	link := rest.Link{Rel: rel, Href: href, Type: contenttype, Method: method}
	h.Links = append(h.Links, &link)
}

// Optionen für Listenelemente kommen aus dem proto als beliebiger Typ daher, jedoch immer in der gleichen nummerierung
// diese werden in die QueryOptions Form gebracht, damit upper sauber damit umgehen kann.
func GetListOptionsFromRequest(options interface{}) QueryOptions {
	tmp, _ := json.Marshal(options)
	var opts QueryOptions
	json.Unmarshal(tmp, &opts)
	return opts
}

// hateoas anhand DBMEta für eine Collection erzeugen
func GenerateCollectionHATEOAS(dbMeta DBMeta) Hateoas {
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
