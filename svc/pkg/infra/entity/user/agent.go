package entity

type (
	Agent struct {
		Roles []Role `json:"agent-roles"`
	}
	Role struct {
		ID          int64 `json:"id"`
		Level       int   `json:"level"`
		GrantedTime int64 `json:"granted_time"`
	}
)
