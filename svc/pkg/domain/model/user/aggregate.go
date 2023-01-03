package user

type (
	User struct {
		ID     ID
		Status Status
		Detail Detail
		Line   Line
	}
	Status int
)

const (
	// StatusNew indicates that user is newly created and hasn't finished its basic registration.
	StatusNew Status = 1
	// StatusRegistered indicates that user has finished its basic registration.
	StatusRegistered Status = 2
)

func (u User) IsValid() bool {
	return u.ID != 0 && u.Line.LineServiceID != ""
}
