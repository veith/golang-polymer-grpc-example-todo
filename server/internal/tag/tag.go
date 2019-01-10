package tag

import (
	"../_env"

	"upper.io/db.v3"
)

// Interface zur DB
var tags db.Collection
var paginationDefault uint

func InitEnvironment(env *environment.Env) {
	tags = env.DB.Collection("tag")
	paginationDefault = 23
}
