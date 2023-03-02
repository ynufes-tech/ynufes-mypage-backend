package entity

type (
	Admin struct {
		IsSuperAdmin bool  `json:"admin-super_admin"`
		GrantedTime  int64 `json:"admin-granted_time"`
	}
)
