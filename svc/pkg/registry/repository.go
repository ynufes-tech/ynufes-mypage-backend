package registry

import (
	"cloud.google.com/go/firestore"
	"context"
	"ynufes-mypage-backend/pkg/setting"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

type Repository struct {
	fs *firestore.Client
}

func NewRepository() (Repository, error) {
	config := setting.Get()
	fs, err := firestore.NewClient(context.Background(), config.Infrastructure.Firestore.ProjectID)
	if err != nil {
		return Repository{}, err
	}
	return Repository{
		fs: fs,
	}, nil
}

func (repo Repository) NewUserQuery() query.User {
	return reader.NewUser(repo.fs)
}

func (repo Repository) NewUserCommand() command.User {
	return writer.NewUser(repo.fs)
}
