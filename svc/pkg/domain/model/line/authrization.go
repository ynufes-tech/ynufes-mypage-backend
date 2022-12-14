package line

import "time"

type (
	EncryptedAccessToken  string
	EncryptedRefreshToken string
	Authorization         struct {
		EncryptedAccessToken  EncryptedAccessToken
		EncryptedRefreshToken EncryptedRefreshToken
		CreatedAt             time.Time
		UpdatedAt             time.Time
	}
)
