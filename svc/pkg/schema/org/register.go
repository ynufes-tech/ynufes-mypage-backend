package org

type RegisterRequest struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Added     bool   `json:"added"`
	OrgID     string `json:"org_id"`
	OrgName   string `json:"org_name"`
	EventID   string `json:"event_id"`
	EventName string `json:"event_name"`
}
