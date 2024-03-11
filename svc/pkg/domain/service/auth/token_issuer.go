package auth

type TokenIssuer interface {
	IssueNewCode(id string) (string, error)
	IssueToken(code string) (string, error)
	RevokeOldCodes()
}
