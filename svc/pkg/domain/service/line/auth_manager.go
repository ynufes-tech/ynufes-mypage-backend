package line

type AuthManager interface {
	Auth(code string) (string, error)
	Verify(token string) (string, error)
}
