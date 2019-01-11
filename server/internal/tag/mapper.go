package tag

import (
	proto "../../../proto/tag"
	"../pkg/hateoas"
	"../pkg/query"
	"context"
	"github.com/oklog/ulid"
)

// ProtoTag in Tag umwandeln
func mapProtoTagToTag(pbTag *proto.Tag) *Tag {
	ulid, _ := ulid.Parse(pbTag.Id)
	return &Tag{ulid, pbTag.Label}
}

// Array mit Tags in eine Collection mappen
func mapTagListToTagCollection(ctx context.Context, tagList []*Tag, dbMeta query.DBMeta) *proto.TagCollection {
	var tags []*proto.TagEntity
	for _, tag := range tagList {
		tagEntity := mapTagToTagEntity(ctx, tag)
		tags = append(tags, tagEntity)
	}

	tagCollection := &proto.TagCollection{Data: tags, Links: hateoas.GenerateCollectionHATEOAS(ctx, "/tags", dbMeta).Links}
	return tagCollection
}

// Tag in eine TagEntity mappen
func mapTagToTagEntity(ctx context.Context, tag *Tag) *proto.TagEntity {
	tagEntity := &proto.TagEntity{}
	tagEntity.Data = mapTagToProto(tag)
	tagEntity.Links = hateoas.GenerateEntityHateoas(ctx, "/tags", tag.Id.String()).Links
	return tagEntity
}

// tag in ProtoTag umwandeln
func mapTagToProto(tag *Tag) *proto.Tag {
	out := &proto.Tag{}
	out.Id = tag.Id.String()
	out.Label = tag.Label
	return out
}
