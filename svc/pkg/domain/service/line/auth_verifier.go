package line

import linePkg "ynufes-mypage-backend/pkg/line"

type AuthVerifier interface {
	IssueNewState() string
	RevokeOldStates()
	RequestAccessToken(code string, state string) (*linePkg.AccessTokenResponse, error)
	VerifyAccessToken(accessToken string) (*linePkg.VerifyResponse, error)
}
