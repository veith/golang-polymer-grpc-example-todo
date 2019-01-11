package task

import (
	"../_env"
	"../pkg/query"
	"../pkg/uid"
	"github.com/oklog/ulid"
	"upper.io/db.v3"
)

// Interface zur Env
var env *environment.Environment

// tasks collection
var tasks db.Collection

func Register() {
	env = environment.Env
	tasks = env.DB.Collection("task")
}

// Task
type Task struct {
	Id          ulid.ULID `json:"id,omitempty" db:"id,pk,omitempty"`
	Title       string    `json:"title,omitempty" db:"title,omitempty"`
	Description string    `json:"description,omitempty" db:"description,omitempty"`
	Completed   int32     `json:"completed,omitempty" db:"completed,omitempty"`
}

func CreateTaskItem(data *Task) (Task, error) {
	var item Task
	item.Id = uid.GenerateULID()
	item.Title = data.Title
	item.Description = data.Description
	if data.Completed != 0 {
		item.Completed = data.Completed
	} else {
		item.Completed = 1
	}
	// id interface not needed, we create the ids ourself
	_, err := tasks.Insert(&item)
	//fire("item.generated",&item)
	return item, err
}

func ListTaskItems(options query.QueryOptions) ([]Task, query.DBMeta, error) {

	res := tasks.Find()
	var meta query.DBMeta
	res, meta = query.ApplyRequestOptionsToQuery(res, options)
	var items []Task
	err := res.All(&items)

	return items, meta, err
}

func CompleteTaskItem(id ulid.ULID) (Task, error) {
	var item Task
	item.Completed = 2
	return UpdateTaskItem(id, &item)
}

func DeleteTaskItem(id ulid.ULID) error {
	var item Task
	res := tasks.Find(db.Cond{"id": id})
	if err := res.One(&item); err != nil {
		return err
	}
	err := res.Delete()
	return err
}

func GetTaskItem(id ulid.ULID) (Task, error) {
	var item Task
	res := tasks.Find(db.Cond{"id": id})
	err := res.One(&item)
	return item, err
}

func UpdateTaskItem(id ulid.ULID, data *Task) (Task, error) {
	var item Task
	res := tasks.Find(db.Cond{"id": id})
	// fields to update
	item.Id = id
	item.Title = data.Title
	item.Description = data.Description
	item.Completed = data.Completed

	if err := res.Update(&item); err != nil {
		return Task{}, err
	}
	// read your write
	err := res.One(&item)
	return item, err
}
