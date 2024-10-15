package uuid

func (s *IdObserver[T]) Creating(model T) {
	id, err := NewId()
	if err != nil {
		panic(err)
	}
	model.SetId(id)
}

func (s *IdObserver[T]) Updating(model T) {

}
