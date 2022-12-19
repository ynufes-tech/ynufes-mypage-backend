package entity

type Line struct {
	LineServiceID         string `firestore:"line-id"`
	LineProfileURL        string `firestore:"line-profile_url"`
	LineDisplayName       string `firestore:"line-display_name"`
	EncryptedAccessToken  string `firestore:"line-access_token"`
	EncryptedRefreshToken string `firestore:"line-refresh_token"`
}
