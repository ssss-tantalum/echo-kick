package model

import (
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID       int64  `bun:"id,pk,autoincrement"`
	Username string `bun:"username,notnull"`
	Email    string `bun:"email,notnull,unique"`
}
