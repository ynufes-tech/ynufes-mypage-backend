package entity

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

const UserRootName = "Users"

type User struct {
	ID         id.UserID `json:"-"`
	UserDetail `json:"detail"`
}

func (u User) ToModel() (*user.User, error) {
	email, err := user.NewEmail(u.Email)
	if err != nil {
		return nil, err
	}
	gender, err := user.NewGender(u.Gender)
	if err != nil {
		return nil, err
	}
	ty, err := user.NewType(u.UserDetail.Type)
	if err != nil {
		return nil, err
	}
	return &user.User{
		ID: u.ID,
		Detail: user.Detail{
			Name: user.Name{
				FirstName:     u.NameFirst,
				LastName:      u.NameLast,
				FirstNameKana: u.NameFirstKana,
				LastNameKana:  u.NameLastKana,
			},
			Email:  email,
			Gender: gender,
			// TODO: add validation for StudentID, Type
			StudentID:  user.StudentID(u.StudentID),
			Type:       ty,
			PictureURL: user.PictureURL(u.PictureURL),
		},
	}, nil
}
