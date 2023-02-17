package entity

type (
	Admin struct {
		IsSuperAdmin bool  `firestore:"admin-super_admin"`
		GrantedTime  int64 `firestore:"admin-granted_time"`
	}
)
