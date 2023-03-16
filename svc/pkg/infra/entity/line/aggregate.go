package line

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/line"
	"ynufes-mypage-backend/svc/pkg/exception"
)

const LineRootName = "Lines"

type Line struct {
	LineServiceID         string `json:"-"`
	UserID                string `json:"user_id"`
	LineDisplayName       string `json:"display_name"`
	EncryptedAccessToken  string `json:"a_token"`
	EncryptedRefreshToken string `json:"r_token"`
}

func (e Line) ToModel() (*line.LineUser, error) {
	uid, err := identity.ImportID(e.UserID)
	if err != nil || e.LineServiceID == "" {
		return nil, exception.ErrIDNotAssigned
	}
	return &line.LineUser{
		UserID:                uid,
		LineServiceID:         line.LineServiceID(e.LineServiceID),
		LineDisplayName:       e.LineDisplayName,
		EncryptedAccessToken:  line.EncryptedAccessToken(e.EncryptedAccessToken),
		EncryptedRefreshToken: line.EncryptedRefreshToken(e.EncryptedRefreshToken),
	}, nil
}
