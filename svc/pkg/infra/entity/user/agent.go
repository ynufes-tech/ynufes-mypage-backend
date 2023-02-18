package entity

type (
	Agent struct {
		Roles []Role `firestore:"agent-roles"`
	}
	Role struct {
		ID          int64 `firestore:"id"`
		Level       int   `firestore:"level"`
		GrantedTime int64 `firestore:"granted_time"`
	}
)
