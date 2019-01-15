package task

import (
	proto "../../../proto/task"
	"../pkg/query"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/oklog/ulid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var RegisterServiceServer = proto.RegisterTaskServiceServer

// Gibt den grpc ServiceServer zurück
func GetServiceServer() proto.TaskServiceServer {
	var s serviceServer
	return &s
}

// serviceServer is used to implement taskServiceServer.
type serviceServer struct {
}

// [POST] ~/tasks:complete
func (s *serviceServer) CompleteTask(ctx context.Context, req *proto.CompleteTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)
	item, err := CompleteTaskItem(taskID)

	return mapTaskToProtoTaskEntity(ctx, item), err
}

// [POST] ~/task
func (s *serviceServer) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.TaskEntity, error) {
	item, err := CreateTaskItem(mapProtoTaskToTask(req.Body))
	return mapTaskToProtoTaskEntity(ctx, item), err
}

// [DELETE] ~/task/{id}
func (s *serviceServer) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*empty.Empty, error) {
	taskID, _ := ulid.Parse(req.Id)
	err := DeleteTaskItem(taskID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return &empty.Empty{}, nil
}

// [PATCH] ~/task/{id}
func (s *serviceServer) UpdateTask(ctx context.Context, req *proto.UpdateTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)

	item, err := UpdateTaskItem(taskID, mapProtoTaskToTask(req.Body))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}

	return mapTaskToProtoTaskEntity(ctx, item), nil
}

// [GET] ~/task/{id}
func (s *serviceServer) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)
	item, err := GetTaskItem(taskID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Task not Found: %s", err)
	}
	return mapTaskToProtoTaskEntity(ctx, item), nil
}

// [GET] ~/tasks
func (serviceServer) ListTask(ctx context.Context, req *proto.ListTaskRequest) (*proto.TaskCollection, error) {
	queryOptions := query.GetListOptionsFromRequest(req)
	taskList, dbMeta, err := ListTasks(queryOptions)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}

	return mapTaskListToProtoTaskCollection(ctx, taskList, dbMeta), nil
}
