package user

type (
	User struct {
		ID        ID
		Detail    Detail
		Line      Line
		Dashboard Dashboard
	}
	ID int64
)
