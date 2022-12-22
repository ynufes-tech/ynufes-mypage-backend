package user

type (
	User struct {
		ID        ID
		Status    Status
		Detail    Detail
		Line      Line
		Dashboard Dashboard
	}
	Status int
)

const (
	// StatusNew indicates that user is newly created and hasn't finished its basic registration.
	StatusNew Status = 0
	// StatusRegistered indicates that user has finished its basic registration.
	StatusRegistered Status = 1
)
