package environment

import (
	"upper.io/db.v3/lib/sqlbuilder"
)

// Settings, Config, .. können nachher hierrüber allen packages mitgegeben werden
type Env struct {
	DB sqlbuilder.Database
}
