package entity

type Line struct {
	LineServiceID         string `json:"id"`
	LineProfileURL        string `json:"profile_url"`
	LineDisplayName       string `json:"display_name"`
	EncryptedAccessToken  string `json:"access_token"`
	EncryptedRefreshToken string `json:"refresh_token"`
}
