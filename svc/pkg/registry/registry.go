package registry

type Registry struct {
	repo Repository
	svc  Service
}

func New() (*Registry, error) {
	repository, err := NewRepository()
	if err != nil {
		return nil, err
	}
	svc := NewService()
	return &Registry{
		repo: repository,
		svc:  svc,
	}, nil
}

func (r Registry) Service() Service {
	return r.svc
}

func (r Registry) Repository() Repository {
	return r.repo
}
