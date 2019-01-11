package tag

import (
	proto "../../../proto/tag"
	"context"
)

var RegisterServiceServer = proto.RegisterTagServiceServer

// Gibt den grpc ServiceServer zurück
func GetServiceServer() proto.TagServiceServer {
	var s serviceServer
	return &s
}

func NewProtoTags() *proto.Tag {
	return &proto.Tag{}
}

// taskServiceServer is used to implement taskServiceServer.
type serviceServer struct {
}

func (serviceServer) CreateTag(ctx context.Context, req *proto.CreateTagRequest) (*proto.TagEntityResponse, error) {
	panic("implement me")
}

func (serviceServer) GetTag(ctx context.Context, req *proto.GetTagRequest) (*proto.TagEntityResponse, error) {
	panic("implement me")
}

func (serviceServer) ListAllTags(ctx context.Context, req *proto.ListTagsRequest) (*proto.TagCollectionResponse, error) {
	res := tags.Find()

	var tagItems []Tag
	err := res.All(&tagItems)

	return &proto.TagCollectionResponse{}, err

}

func (serviceServer) ListTagsFromTask(ctx context.Context, req *proto.ListTagsRequest) (*proto.TagCollectionResponse, error) {
	res := tags.Find()

	var tagItems []Tag
	err := res.All(&tagItems)

	return &proto.TagCollectionResponse{}, err
}

func (serviceServer) DeleteTag(ctx context.Context, req *proto.DeleteTagRequest) (*proto.TagEntityResponse, error) {
	panic("implement me")
}

func (serviceServer) UpdateTag(ctx context.Context, req *proto.UpdateTagRequest) (*proto.TagEntityResponse, error) {
	panic("implement me")
}
