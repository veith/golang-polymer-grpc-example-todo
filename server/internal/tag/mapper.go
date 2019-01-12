package tag

import (
	proto "../../../proto/tag"
	"../pkg/hateoas"
	"../pkg/query"
	"context"
	"github.com/oklog/ulid"
)

// Converts a protobuf Tag in a Tag
func mapProtoTagToTag(pbTag *proto.Tag) *Tag {
	ulid, _ := ulid.Parse(pbTag.Id)
	return &Tag{ulid, pbTag.Label}
}

// Maps an array with Tags to a protobuf Tag Collection
// DBMeta and Context  is used for HATEOAS
func mapTagListToProtoTagCollection(ctx context.Context, tagList []*Tag, dbMeta query.DBMeta) *proto.TagCollection {
	var tags []*proto.TagEntity
	for _, tag := range tagList {
		tagEntity := mapTagToProtoTagEntity(ctx, tag)
		tags = append(tags, tagEntity)
	}

	tagCollection := &proto.TagCollection{Data: tags, Links: hateoas.GenerateCollectionHATEOAS(ctx, "/tags", dbMeta).Links}
	return tagCollection
}

// Maps a Tag to a protobuf Tag Entity
// DBMeta and Context  is used for HATEOAS
func mapTagToProtoTagEntity(ctx context.Context, tag *Tag) *proto.TagEntity {
	tagEntity := &proto.TagEntity{}
	tagEntity.Data = mapTagToProto(tag)
	tagEntity.Links = hateoas.GenerateEntityHateoas(ctx, "/tags", tag.Id.String()).Links
	return tagEntity
}

// Maps a Tag to a protobuf Tag
func mapTagToProto(tag *Tag) *proto.Tag {
	out := &proto.Tag{}
	out.Id = tag.Id.String()
	out.Label = tag.Label
	return out
}
