package tag

import (
	"../_env"
	"github.com/oklog/ulid"
	"upper.io/db.v3"
)

// Interface zur DB
var tags db.Collection

func InitEnvironment(env *environment.Env) {
	tags = env.DB.Collection("tag")
}

type Tag struct {
	Id    ulid.ULID `json:"id,omitempty" db:"id,pk,omitempty"`
	label string    `json:"label,omitempty" db:"title,omitempty"`
}
