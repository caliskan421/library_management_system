package model

import (
	"time"

	"github.com/uptrace/bun"
)

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        string    `bun:"id,pk,type:uuid,default:gen_random_uuid()" json:"_id"`
	Name      string    `bun:"name,notnull" json:"name"`
	Email     string    `bun:"email,notnull,unique" json:"email"`
	Password  string    `bun:"password,notnull" json:"-"`
	Role      Role      `bun:"role,notnull,default:'user'" json:"role"`
	CreatedAt time.Time `bun:"created_at,notnull,default:current_timestamp" json:"createdAt"`
	UpdatedAt time.Time `bun:"updated_at,notnull,default:current_timestamp" json:"updatedAt"`
}
