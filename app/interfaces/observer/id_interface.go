package observer

type SetIdInterface interface {
	SetId(id string)
}

type IdInterface interface {
	SetId(id string)
	GetId() string
}
