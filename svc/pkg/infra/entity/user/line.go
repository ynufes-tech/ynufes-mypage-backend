package entity

type LineAuthorization struct {
	LineServiceID         string `firestore:"line-id"`
	EncryptedAccessToken  string `firestore:"line-access_token"`
	EncryptedRefreshToken string `firestore:"line-refresh_token"`
}
