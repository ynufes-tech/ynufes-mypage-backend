package org

import "ynufes-mypage-backend/svc/pkg/schema/form"

type OrgResponse struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	EventID string             `json:"event_id"`
	Forms   []form.FormSummary `json:"forms"`
}
