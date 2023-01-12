package command

type Event interface {
	Create(*Event) error
	Update(*Event) error
	Delete(*Event) error
}
