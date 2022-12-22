package entity

import "ynufes-mypage-backend/svc/pkg/domain/model/user"

type User struct {
	ID     string
	Status int `firestore:"status"`
	UserDetail
	Line
	UserDashboard
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
	id, err := user.Import(u.ID)
	return &user.User{
		ID:     id,
		Status: user.Status(u.Status),
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
			StudentID: user.StudentID(u.StudentID),
			Type:      user.Type(u.Type),
		},
		Line: user.Line{
			LineServiceID:         user.LineServiceID(u.LineServiceID),
			LineProfilePictureURL: user.LineProfilePictureURL(u.LineProfileURL),
			LineDisplayName:       u.LineDisplayName,
			EncryptedAccessToken:  user.EncryptedAccessToken(u.EncryptedAccessToken),
			EncryptedRefreshToken: user.EncryptedRefreshToken(u.EncryptedRefreshToken),
		},
	}, nil
}
