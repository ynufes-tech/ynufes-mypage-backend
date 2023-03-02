package entity

type Line struct {
	LineServiceID         string `json:"line-id"`
	LineProfileURL        string `json:"line-profile_url"`
	LineDisplayName       string `json:"line-display_name"`
	EncryptedAccessToken  string `json:"line-access_token"`
	EncryptedRefreshToken string `json:"line-refresh_token"`
}
