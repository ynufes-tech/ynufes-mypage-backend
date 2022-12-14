package command

type Form interface {
	Create(*Form) error
	Delete(*Form) error
	Grant(*Form)
}
