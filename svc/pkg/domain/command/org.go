package command

type Org interface {
	Create(*Org) error
	Update(*Org) error
	Delete(*Org) error
}
