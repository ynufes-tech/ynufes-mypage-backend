package org

import "ynufes-mypage-backend/svc/pkg/schema/form"

type OrgResponse struct {
	ID        string             `json:"org_id"`
	Name      string             `json:"org_name"`
	EventID   string             `json:"event_id"`
	EventName string             `json:"event_name"`
	Forms     []form.FormSummary `json:"forms"`
}

func NewOrgResponse(orgID, orgName, eventID, eventName string, forms []form.FormSummary) OrgResponse {
	return OrgResponse{
		ID:        orgID,
		Name:      orgName,
		EventID:   eventID,
		EventName: eventName,
		Forms:     forms,
	}
}
