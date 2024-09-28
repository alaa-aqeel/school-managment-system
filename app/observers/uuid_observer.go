package observers

import (
	"github.com/alaa-aqeel/school-managment-system/app/interfaces"
	"github.com/alaa-aqeel/school-managment-system/helpers"
)

type IdObserver[T interfaces.SetIdInterface] struct{}

func (s *IdObserver[T]) Creating(model T) {
	model.SetId(helpers.UUID())
}

func (s *IdObserver[T]) Updating(model T) {

}
