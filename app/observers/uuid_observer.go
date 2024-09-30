package observers

import (
	"github.com/alaa-aqeel/school-managment-system/app/interfaces"
	"github.com/gofrs/uuid"
)

type IdObserver[T interfaces.SetIdInterface] struct{}

func (s *IdObserver[T]) Creating(model T) {
	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	model.SetId(id.String())
}

func (s *IdObserver[T]) Updating(model T) {

}
