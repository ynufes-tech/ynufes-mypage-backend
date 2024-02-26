package registry

type Registry struct {
	repo Repository
	svc  Service
}

func New() (*Registry, error) {
	repo, err := NewRepository()
	if err != nil {
		return nil, err
	}
	svc := NewService(&repo)
	return &Registry{
		repo: repo,
		svc:  svc,
	}, nil
}

func (r Registry) Service() Service {
	return r.svc
}

func (r Registry) Repository() Repository {
	return r.repo
}
