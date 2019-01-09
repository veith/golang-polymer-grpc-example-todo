package task

import (
	"../_env"
	"github.com/oklog/ulid"
	"upper.io/db.v3"
)

// Interface zur DB
var tasks db.Collection
var paginationDefault uint

func InitEnvironment(env *environment.Env) {
	tasks = env.DB.Collection("task")
	paginationDefault = 23
}

func CreateTaskItem(data *Task) (Task, error) {
	var item Task
	item.Id = GenerateULID()
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

func ListTaskItems(options QueryOptions) ([]Task, DBMeta, error) {

	res := tasks.Find()
	var meta DBMeta
	res, meta = ApplyRequestOptionsToQuery(res, options)
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
