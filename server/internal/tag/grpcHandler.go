// In this file are the grpc handlers. Our API talks grpc...

package tag

import (
	proto "../../../proto/tag"
	"../pkg/dberrors"
	"../pkg/query"
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/oklog/ulid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var RegisterServiceServer = proto.RegisterTagServiceServer

func GetServiceServer() proto.TagServiceServer {
	var s serviceServer
	return &s
}

type serviceServer struct {
}

func (serviceServer) CreateTag(ctx context.Context, req *proto.CreateTagRequest) (*proto.TagEntity, error) {
	tag, err := CreateTag(mapProtoTagToTag(req.Item))
	if err != nil {
		if dberrors.FindErrorByMessageString(err, "uniq") {
			return nil, status.Errorf(codes.AlreadyExists, "Constraint violation: %s", err)
		}
		return nil, err
	}
	return mapTagToProtoTagEntity(ctx, tag), nil
}

func (serviceServer) GetTag(ctx context.Context, req *proto.GetTagRequest) (*proto.TagEntity, error) {
	tagID, _ := ulid.Parse(req.Id)
	item, err := GetTag(tagID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Tag not Found: %s", err)
	}
	return mapTagToProtoTagEntity(ctx, item), nil
}

func (serviceServer) ListAllTags(ctx context.Context, req *proto.ListTagsRequest) (*proto.TagCollection, error) {
	queryOptions := query.GetListOptionsFromRequest(req)
	tagList, dbMeta, err := ListTags(queryOptions)

	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}

	return mapTagListToProtoTagCollection(ctx, tagList, dbMeta), nil
}

func (serviceServer) ListTagsFromTask(ctx context.Context, req *proto.ListTagsRequest) (*proto.TagCollection, error) {
	queryOptions := query.GetListOptionsFromRequest(req)
	tagList, dbMeta, err := ListTags(queryOptions)
	if err != nil {
		return nil, status.Errorf(codes.Unavailable, "Data Error: %s", err)
	}

	return mapTagListToProtoTagCollection(ctx, tagList, dbMeta), nil
}

func (serviceServer) DeleteTag(ctx context.Context, req *proto.DeleteTagRequest) (*empty.Empty, error) {
	tagID, _ := ulid.Parse(req.Id)
	err := DeleteTag(tagID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Tag not Found: %s", err)
	}
	return &empty.Empty{}, nil
}

func (serviceServer) UpdateTag(ctx context.Context, req *proto.UpdateTagRequest) (*proto.TagEntity, error) {
	tagID, _ := ulid.Parse(req.Id)
	tag, err := UpdateTag(tagID, mapProtoTagToTag(req.Item))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Tag not Found: %s", err)
	}
	return mapTagToProtoTagEntity(ctx, tag), nil
}
