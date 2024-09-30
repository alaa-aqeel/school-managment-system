package interfaces

type Observer[T any] interface {
	Creating(model T)
	Updating(model T)
}
