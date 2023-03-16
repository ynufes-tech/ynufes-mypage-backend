package line

import (
	"ynufes-mypage-backend/pkg/cipher"
	"ynufes-mypage-backend/svc/pkg/config"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

type (
	LineUser struct {
		UserID                id.UserID
		LineServiceID         LineServiceID
		LineDisplayName       string
		EncryptedAccessToken  EncryptedAccessToken
		EncryptedRefreshToken EncryptedRefreshToken
	}
	LineServiceID         string
	EncryptedAccessToken  string
	EncryptedRefreshToken string
	PlainAccessToken      string
	PlainRefreshToken     string
)

var aes *cipher.AES

func init() {
	aes, _ = cipher.NewAES(config.Line.CipherKey)
}

func NewEncryptedAccessToken(s PlainAccessToken) EncryptedAccessToken {
	encrypted := aes.Encrypt(string(s))
	return EncryptedAccessToken(encrypted)
}

func NewEncryptedRefreshToken(s PlainRefreshToken) EncryptedRefreshToken {
	encrypted := aes.Encrypt(string(s))
	return EncryptedRefreshToken(encrypted)
}

func (l Line) AccessToken() (PlainAccessToken, error) {
	decrypted, err := aes.Decrypt(string(l.EncryptedAccessToken))
	if err != nil {
		return PlainAccessToken(""), err
	}
	return PlainAccessToken(decrypted), nil
}

func (l Line) RefreshToken() (PlainRefreshToken, error) {
	decrypted, err := aes.Decrypt(string(l.EncryptedRefreshToken))
	if err != nil {
		return PlainRefreshToken(""), err
	}
	return PlainRefreshToken(decrypted), nil
}
