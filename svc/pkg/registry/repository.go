package registry

import (
	"cloud.google.com/go/firestore"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
)

type Repository struct {
	fs *firestore.Client
}

func NewRepository(fs *firestore.Client) Repository {
	return Repository{
		fs: fs,
	}
}

func (repo Repository) NewUserQuery() query.User {
	return reader.NewUser(repo.fs)
}
