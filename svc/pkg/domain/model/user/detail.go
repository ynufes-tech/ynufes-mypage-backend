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

// IsValid TODO: implement validation for StudentID
func (s StudentID) IsValid() bool {
	return s != ""
}

func (e Email) IsValid() bool {
	return e != ""
}

func (n Name) HasAllValue() bool {
	return n.FirstName != "" && n.LastName != "" && n.FirstNameKana != "" && n.LastNameKana != ""
}

func (d Detail) MeetsBasicRequirement() bool {
	return d.Name.HasAllValue() && d.StudentID.IsValid() && d.Email.IsValid()
}
