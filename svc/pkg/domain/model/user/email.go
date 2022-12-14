package user

type (
	Email string
)

// NewEmail TODO: add validation
func NewEmail(email string) (Email, error) {
	return Email(email), nil
}
