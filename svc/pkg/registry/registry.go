package registry

type Registry struct {
	repo Repository
	svc  Service
}

func (r Registry) Service() Service {
	return r.svc
}

func (r Registry) Repository() Repository {
	return r.repo
}
