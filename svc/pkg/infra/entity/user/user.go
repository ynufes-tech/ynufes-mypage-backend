package entity

type (
	ID   int64
	User struct {
		ID ID `firestore:"id"`
	}
)
