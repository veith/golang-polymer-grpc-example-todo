package tag

import (
	proto "../../../proto/tag"
	"../pkg/hateoas"
	"../pkg/query"
	"github.com/oklog/ulid"
)

// ProtoTag in Tag umwandeln
func mapProtoTagToTag(pbTag *proto.Tag) *Tag {
	ulid, _ := ulid.Parse(pbTag.Id)
	return &Tag{ulid, pbTag.Label}
}

// Array mit Tags in eine Collection mappen
func mapTagListToTagCollection(tagList []*Tag, dbMeta query.DBMeta) *proto.TagCollection {
	var tags []*proto.TagEntity
	for _, tag := range tagList {
		tagEntity := mapTagToTagEntity(tag)
		tags = append(tags, tagEntity)
	}

	tagCollection := &proto.TagCollection{Data: tags, Links: hateoas.GenerateCollectionHATEOAS(dbMeta).Links}
	return tagCollection
}

// Tag in eine TagEntity mappen
func mapTagToTagEntity(tag *Tag) *proto.TagEntity {
	tagEntity := &proto.TagEntity{}
	tagEntity.Data = mapTagToProto(tag)
	tagEntity.Links = hateoas.GenerateEntityHateoas(tag.Id.String()).Links
	return tagEntity
}

// tag in ProtoTag umwandeln
func mapTagToProto(tag *Tag) *proto.Tag {
	out := &proto.Tag{}
	out.Id = tag.Id.String()
	out.Label = tag.Label
	return out
}
