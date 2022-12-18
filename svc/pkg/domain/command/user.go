package command

import "ynufes-mypage-backend/svc/pkg/domain/model/user"

type User interface {
	Create(*user.User) error
	UpdateLineAuth(*user.User) error
	UpdateAll(*user.User) error
	Delete(*user.User) error
}
