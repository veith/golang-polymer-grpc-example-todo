package task

import (
	proto "../../../proto/task"
	"../pkg/hateoas"

	"../pkg/query"
	"context"
	"github.com/oklog/ulid"
)

// Converts a protobuf Task in a Task
func mapProtoTaskToTask(obj *proto.Task) *Task {
	ulid, _ := ulid.Parse(obj.Id)
	return &Task{ulid, obj.Title, obj.Description, int32(obj.Completed)}
}

// Maps an array with Tasks to a protobuf Task Collection
// DBMeta and Context  is used for HATEOAS
func mapTaskListToProtoTaskCollection(ctx context.Context, taskList []*Task, dbMeta query.DBMeta) *proto.TaskCollection {
	var tasks []*proto.TaskEntity
	for _, task := range taskList {
		taskEntity := mapTaskToProtoTaskEntity(ctx, task)
		tasks = append(tasks, taskEntity)
	}

	taskCollection := &proto.TaskCollection{Data: tasks, Links: hateoas.GenerateCollectionHATEOAS(ctx, "/tasks", dbMeta).Links}
	return taskCollection
}

// Maps a Task to a protobuf Task Entity
// DBMeta and Context  is used for HATEOAS
func mapTaskToProtoTaskEntity(ctx context.Context, task *Task) *proto.TaskEntity {
	taskEntity := &proto.TaskEntity{}
	taskEntity.Data = mapTaskToProtoTask(task)

	hts := hateoas.GenerateEntityHateoas(ctx, "/tasks", task.Id.String())

	hts.AddSubresource("tags")

	taskEntity.Links = hts.Links
	return taskEntity
}

// Maps a Task to a protobuf Task
func mapTaskToProtoTask(task *Task) *proto.Task {
	out := &proto.Task{}
	out.Id = task.Id.String()
	out.Title = task.Title
	out.Description = task.Description
	out.Completed = proto.Complete(task.Completed)
	return out
}
