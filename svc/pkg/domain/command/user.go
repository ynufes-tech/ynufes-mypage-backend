package command

type User interface {
	Create(*User) error
	UpdateLineAuth(*User) error
	UpdateAll(*User) error
	Delete(*User) error
}
