package command

type User interface {
	Create(*User) error
	Delete(*User) error
}
