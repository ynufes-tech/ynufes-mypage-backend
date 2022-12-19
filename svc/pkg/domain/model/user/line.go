package user

import (
	"ynufes-mypage-backend/pkg/cipher"
	"ynufes-mypage-backend/svc/pkg/config"
)

type (
	LineServiceID         string
	LineProfilePictureURL string
	EncryptedAccessToken  string
	EncryptedRefreshToken string
	PlainAccessToken      string
	PlainRefreshToken     string
	Line                  struct {
		LineServiceID         LineServiceID
		LineProfilePictureURL LineProfilePictureURL
		LineDisplayName       string
		EncryptedAccessToken  EncryptedAccessToken
		EncryptedRefreshToken EncryptedRefreshToken
	}
)

var aes *cipher.AES

func init() {
	aes, _ = cipher.NewAES(config.Line.CipherKey)
}

func NewEncryptedAccessToken(s PlainAccessToken) (EncryptedAccessToken, error) {
	encrypted := aes.Encrypt(string(s))
	return EncryptedAccessToken(encrypted), nil
}

func NewEncryptedRefreshToken(s PlainRefreshToken) (EncryptedRefreshToken, error) {
	encrypted := aes.Encrypt(string(s))
	return EncryptedRefreshToken(encrypted), nil
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
