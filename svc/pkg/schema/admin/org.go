package admin

type CreateOrgResponse struct {
	EventID   string `json:"event_id"`
	EventName string `json:"event_name"`
	OrgID     string `json:"org_id"`
	OrgName   string `json:"org_name"`
	IsOpen    bool   `json:"is_open"`
}

type IssueOrgInviteTokenResponse struct {
	Token      string `json:"token"`
	OrgID      string `json:"org_id"`
	ValidUntil string `json:"valid_until"`
}
