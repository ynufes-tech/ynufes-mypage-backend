package registry

type Registry struct {
	//repo Repository
	svc Service
}

func (r Registry) Service() Service {
	return r.svc
}
