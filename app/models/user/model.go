package user

import (
	"time"
)

type User struct {
	ID          string
	Username    string
	Password    string
	IsActive    bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeleatedAt  time.Time
	LastLoginAt time.Time
}

func (u *User) SetCreatedAt(t time.Time) {
	u.CreatedAt = t
}

func (u *User) SetUpdatedAt(t time.Time) {
	u.UpdatedAt = t
}

func (u *User) SetId(id string) {
	u.ID = id
}
