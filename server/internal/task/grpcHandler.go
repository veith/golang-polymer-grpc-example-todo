package task

import (
	proto "../../../proto/task"
	"../pkg/hateoas"
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

func (s *serviceServer) CompleteTask(ctx context.Context, req *proto.CompleteTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)
	item, err := CompleteTaskItem(taskID)
	entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: hateoas.GenerateEntityHateoas(ctx, "/tasks", item.Id.String()).Links}

	return &entity, err
}

func (s *serviceServer) CreateTask(ctx context.Context, req *proto.CreateTaskRequest) (*proto.TaskEntity, error) {
	item, err := CreateTaskItem(MapProtoTaskToTask(req.Body))
	entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: hateoas.GenerateEntityHateoas(ctx, "/tasks", item.Id.String()).Links}
	return &entity, err
}

func (s *serviceServer) DeleteTask(ctx context.Context, req *proto.DeleteTaskRequest) (*empty.Empty, error) {
	taskID, _ := ulid.Parse(req.Id)
	err := DeleteTaskItem(taskID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not retrieve entity from the database: %s", err)
	}
	return &empty.Empty{}, nil
}

func (s *serviceServer) UpdateTask(ctx context.Context, req *proto.UpdateTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)

	item, err := UpdateTaskItem(taskID, MapProtoTaskToTask(req.Body))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Could not update entity: %s", err)
	}
	entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: hateoas.GenerateEntityHateoas(ctx, "/tasks", item.Id.String()).Links}
	return &entity, nil
}

func (s *serviceServer) GetTask(ctx context.Context, req *proto.GetTaskRequest) (*proto.TaskEntity, error) {
	taskID, _ := ulid.Parse(req.Id)
	item, err := GetTaskItem(taskID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Task not Found: %s", err)
	}
	entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: hateoas.GenerateEntityHateoas(ctx, "/tasks", item.Id.String()).Links}
	return &entity, nil
}

func (s *serviceServer) ListTask(ctx context.Context, req *proto.ListTaskRequest) (*proto.TaskCollection, error) {
	//token := ctx.Value("tokenInfo")

	opts := query.GetListOptionsFromRequest(req)
	items, dbMeta, err := ListTaskItems(opts)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}

	var collection []*proto.TaskEntity
	for _, item := range items {
		entity := proto.TaskEntity{Data: MapTaskToProtoTask(&item), Links: hateoas.GenerateEntityHateoas(ctx, "/tasks", item.Id.String()).Links}
		collection = append(collection, &entity)
	}

	return &proto.TaskCollection{Data: collection, Links: hateoas.GenerateCollectionHATEOAS(ctx, "/tasks", dbMeta).Links}, nil
}
