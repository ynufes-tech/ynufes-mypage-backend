package user

import "time"

type (
	Admin struct {
		IsSuperAdmin bool
		GrantedTime  time.Time
	}
)
