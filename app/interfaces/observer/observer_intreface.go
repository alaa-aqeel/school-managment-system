package observer

type Observer[T any] interface {
	Creating(model T)
	Updating(model T)
}
