package user

type (
	Detail struct {
		Name      Name
		Email     Email
		Gender    Gender
		StudentID StudentID
		Type      Type
	}
)

// TODO: add validation for Name, StudentID, Type
type (
	Name struct {
		FirstName     string
		LastName      string
		FirstNameKana string
		LastNameKana  string
	}
	StudentID string
)
