package tag

import (
	"../pkg/environment"
	"../pkg/query"
	"../pkg/uid"
	"context"
	"github.com/oklog/ulid"
	"upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

var env *environment.Environment

func Register() {
	// Env provides the DB Session and maybe some config
	env = environment.Env
}

type Tag struct {
	Id    ulid.ULID `json:"id,omitempty" db:"id,pk,omitempty"`
	Label string    `json:"label,omitempty" db:"label,omitempty"`
}

func CreateTag(newTag *Tag) (*Tag, error) {
	// generate an id
	newTag.Id = uid.GenerateULID()

	err := env.DB.Tx(context.Background(), func(tx sqlbuilder.Tx) error {
		// Use `tx` like you would normally use `sess` (Env.DB is sess).
		tags := tx.Collection("tag")
		// id interface not needed, we create the ids ourself
		_, err := tags.Insert(newTag)
		if err != nil {
			// Rollback the transaction by returning any error.
			return err
		}
		// Commit the transaction by returning nil.
		return nil
	})

	return newTag, err
}

func DeleteTag(id ulid.ULID) error {
	tag := &Tag{}
	tags := env.DB.Collection("tag")
	res := tags.Find(db.Cond{"id": id})
	if err := res.One(tag); err != nil {
		return err
	}
	err := res.Delete()
	return err
}

func GetTag(id ulid.ULID) (*Tag, error) {
	var tag *Tag
	tags := env.DB.Collection("tag")
	res := tags.Find(db.Cond{"id": id})
	err := res.One(tag)
	return tag, err
}

func UpdateTag(id ulid.ULID, data *Tag) (*Tag, error) {
	tag := &Tag{}
	err := env.DB.Tx(context.Background(), func(tx sqlbuilder.Tx) error {

		// fields to update
		tag.Id = id
		tag.Label = data.Label

		tags := tx.Collection("tag")
		res := tags.Find(db.Cond{"id": id})

		err := res.Update(tag)
		if err != nil {
			return err
		}
		// Commit the transaction by returning nil.
		return err
	})
	// read your write
	tag, err = GetTag(id)
	return tag, err
}

func ListTags(options query.QueryOptions) ([]*Tag, query.DBMeta, error) {
	tags := env.DB.Collection("tag")
	res := tags.Find()
	var meta query.DBMeta
	res, meta = query.ApplyRequestOptionsToQuery(res, options)
	var tagList []*Tag
	err := res.All(&tagList)
	return tagList, meta, err
}
