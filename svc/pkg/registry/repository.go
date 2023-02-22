package registry

import (
	"cloud.google.com/go/firestore"
	firestorePkg "ynufes-mypage-backend/pkg/firestore"
	"ynufes-mypage-backend/svc/pkg/domain/command"
	"ynufes-mypage-backend/svc/pkg/domain/query"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

type Repository struct {
	fs *firestore.Client
}

func NewRepository() (Repository, error) {
	fs := firestorePkg.New()
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

func (repo Repository) NewEventQuery() query.Event {
	return reader.NewEvent(repo.fs)
}

func (repo Repository) NewEventCommand() command.Event {
	return writer.NewEvent(repo.fs)
}

func (repo Repository) NewOrgQuery() query.Org {
	return reader.NewOrg(repo.fs)
}

func (repo Repository) NewOrgCommand() command.Org {
	return writer.NewOrg(repo.fs)
}

func (repo Repository) NewFormCommand() command.Form {
	return writer.NewForm(repo.fs)
}

func (repo Repository) NewFormQuery() query.Form {
	return reader.NewForm(repo.fs)
}

func (repo Repository) NewQuestionQuery() query.Question {
	return reader.NewQuestion(repo.fs)
}

func (repo Repository) NewQuestionCommand() command.Question {
	return writer.NewQuestion(repo.fs)
}
