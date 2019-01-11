package task

import (
	"../../../proto/task"
	"github.com/gogo/protobuf/types"
	"github.com/oklog/ulid"
	"github.com/veith/protos/date"
)

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
