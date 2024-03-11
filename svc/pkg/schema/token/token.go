package token

type TokenRequest struct {
	Code string `json:"code"`
}

type TokenResponse struct {
	Token string `json:"token"`
}
