package db

import (
	"github.com/notnulldev/e-com-2/types"
)

type UserDb interface {
	GetUserById(id uint) (types.User, error)
}
