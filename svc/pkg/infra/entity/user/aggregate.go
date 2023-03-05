package entity

import (
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

const UserRootName = "Users"

type User struct {
	ID         id.UserID `json:"-"`
	UserDetail `json:"detail"`
	Line       `json:"line"`
	Admin      `json:"admin"`
	Agent      `json:"agent"`
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
	roles := make([]user.Role, len(u.Roles))
	for i, role := range u.Roles {
		lv, err := user.NewRoleLevel(role.Level)
		if err != nil {
			return nil, err
		}
		roles[i] = user.Role{
			ID:          identity.NewID(role.ID),
			Level:       lv,
			GrantedTime: time.UnixMilli(role.GrantedTime),
		}
	}
	var adminGrantedTime *time.Time
	if u.IsSuperAdmin {
		t := time.UnixMilli(u.GrantedTime)
		adminGrantedTime = &t
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
			StudentID: user.StudentID(u.StudentID),
			Type:      ty,
		},
		Line: user.Line{
			LineServiceID:         user.LineServiceID(u.LineServiceID),
			LineProfilePictureURL: user.LineProfilePictureURL(u.LineProfileURL),
			LineDisplayName:       u.LineDisplayName,
			EncryptedAccessToken:  user.EncryptedAccessToken(u.EncryptedAccessToken),
			EncryptedRefreshToken: user.EncryptedRefreshToken(u.EncryptedRefreshToken),
		},
		Admin: user.Admin{
			IsSuperAdmin: u.IsSuperAdmin,
			GrantedTime:  adminGrantedTime,
		},
		Agent: user.Agent{
			Roles: roles,
		},
	}, nil
}
