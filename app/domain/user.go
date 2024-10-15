package domain

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

func (u *User) GetCreatedAt() time.Time {
	return u.CreatedAt
}

func (u *User) SetCreatedAt(t time.Time) {
	u.CreatedAt = t
}

func (u *User) GetUpdatedAt() time.Time {
	return u.UpdatedAt
}

func (u *User) SetUpdatedAt(t time.Time) {
	u.UpdatedAt = t
}

func (u *User) SetDeleatedAt(t time.Time) {
	u.DeleatedAt = t
}

func (u *User) SetLastLoginAt(t time.Time) {
	u.LastLoginAt = t
}

func (u *User) SetId(id string) {
	u.ID = id
}

func (u *User) GetId() string {
	return u.ID
}
