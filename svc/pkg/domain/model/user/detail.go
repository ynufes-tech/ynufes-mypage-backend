package user

type (
	Detail struct {
		Name  Name
		Email Email
	}
	Name struct {
		FirstName     string
		LastName      string
		FirstNameKana string
		LastNameKana  string
	}
)
