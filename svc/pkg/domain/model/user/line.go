package user

type (
	LineServiceID         string
	EncryptedAccessToken  string
	EncryptedRefreshToken string
	Line                  struct {
		LineServiceID         LineServiceID
		EncryptedAccessToken  EncryptedAccessToken
		EncryptedRefreshToken EncryptedRefreshToken
	}
)
