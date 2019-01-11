package uid

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// Erzeuge eine ULID
func GenerateULID() ulid.ULID {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	newID, _ := ulid.New(ulid.Timestamp(t), entropy)
	return newID
}
