package task

import (
	"../pkg/environment"
	"../pkg/query"
	"../pkg/uid"
	"context"
	"github.com/oklog/ulid"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
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
	Title       string    `json:"title,omitempty" db:"title"`
	Description string    `json:"description,omitempty" db:"description"`
	Completed   int32     `json:"completed,omitempty" db:"completed,omitempty"`
}

func CreateTaskItem(data *Task) (*Task, error) {
	item := &Task{}
	item.Id = uid.GenerateULID()
	item.Title = data.Title
	item.Description = data.Description
	if data.Completed != 0 {
		item.Completed = data.Completed
	} else {
		item.Completed = 1
	}
	// id interface not needed, we create the ids ourself
	_, err := tasks.Insert(item)
	//fire("item.generated",&item)
	return item, err
}

func ListTasks(options query.QueryOptions) ([]*Task, query.DBMeta, error) {

	res := tasks.Find()
	var meta query.DBMeta
	res, meta = query.ApplyRequestOptionsToQuery(res, options)
	var items []*Task
	err := res.All(&items)

	return items, meta, err
}

func CompleteTaskItem(id ulid.ULID) (*Task, error) {
	item := &Task{}
	item.Completed = 2
	return UpdateTaskItem(id, item)
}

func DeleteTaskItem(id ulid.ULID) error {
	item := &Task{}
	res := tasks.Find(db.Cond{"id": id})
	if err := res.One(item); err != nil {
		return err
	}
	err := res.Delete()
	return err
}

func GetTaskItem(id ulid.ULID) (*Task, error) {
	item := &Task{}
	res := tasks.Find(db.Cond{"id": id})
	err := res.One(item)
	return item, err
}

func UpdateTaskItem(id ulid.ULID, data *Task) (*Task, error) {
	task := &Task{}
	err := env.DB.Tx(context.Background(), func(tx sqlbuilder.Tx) error {

		// fields to update
		task.Id = id

		task.Title = data.Title
		task.Description = data.Description
		task.Completed = data.Completed

		tasks := tx.Collection("task")
		res := tasks.Find(db.Cond{"id": id})

		err := res.Update(task)
		if err != nil {
			return err
		}
		// Commit the transaction by returning nil.
		return err
	})
	// read your write
	task, err = GetTaskItem(id)
	return task, err
}
