package interfaces

import "time"

type TimestampsInterface interface {
	SetCreatedAt(t time.Time)
	SetUpdatedAt(t time.Time)
}
