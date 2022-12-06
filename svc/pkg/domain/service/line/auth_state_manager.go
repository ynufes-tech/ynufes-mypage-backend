package line

type AuthStateManager interface {
	IssueNewState() string
	VerifyState(state string) error
}
