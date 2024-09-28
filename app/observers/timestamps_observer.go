package observers

import (
	"time"

	"github.com/alaa-aqeel/school-managment-system/app/interfaces"
)

type TimestampsObserver[T interfaces.TimestampsInterface] struct{}

func (s *TimestampsObserver[T]) Creating(model T) {

	model.SetCreatedAt(time.Now())
	model.SetUpdatedAt(time.Now())
}

func (s *TimestampsObserver[T]) Updating(model T) {
	model.SetUpdatedAt(time.Now())
}

func (s *TimestampsObserver[T]) Deleting(model T) {

	// model.SetDeleatedAt(time.Now())
}
