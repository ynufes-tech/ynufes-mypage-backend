package entity

type (
	Admin struct {
		IsSuperAdmin bool  `json:"super_admin"`
		GrantedTime  int64 `json:"granted_time"`
	}
)
