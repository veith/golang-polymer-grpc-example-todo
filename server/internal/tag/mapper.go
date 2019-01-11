package tag

import (
	proto "../../../proto/tag"
)

func mapTagToProto(tag *Tag) *proto.Tag {
	out := &proto.Tag{}
	out.Id = tag.Id.String()
	out.Label = tag.label
	return out
}

func mapTagCollectionToProtoTagCollection(tags []*Tag) *proto.TagCollectionResponse {
	tagCollection := &proto.TagCollectionResponse{}

	return tagCollection
}
