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
