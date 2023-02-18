package user

import (
	"time"
)

type (
	Admin struct {
		IsSuperAdmin bool
		GrantedTime  *time.Time
		// if IsSuperAdmin is false, GrantedTime should be nil.
	}
)
