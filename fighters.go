package rockpapershit

import (
	guuid "github.com/google/uuid"
	"github.com/kudarap/rockpapershit/xerror"
)

var (
	ErrFighterNotFound = xerror.Error("not_found")
	ErrNotFound        = xerror.Error("not_found")
)

type Fighter struct {
	ID        guuid.UUID
	FirstName string
	LastName  string
}
