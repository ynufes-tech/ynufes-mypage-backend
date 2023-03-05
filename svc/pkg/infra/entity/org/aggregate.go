package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
)

const OrgRootName = "Orgs"

type Org struct {
	ID        id.OrgID `json:"-"`
	EventID   string   `json:"event_id"`
	EventName string   `json:"event_name"`
	Name      string   `json:"name"`
	IsOpen    bool     `json:"is_open"`
}

func (o Org) ToModel() (*org.Org, error) {
	eid, err := identity.ImportID(o.EventID)
	if err != nil {
		return nil, err
	}
	return &org.Org{
		ID: o.ID,
		Event: event.Event{
			ID:   eid,
			Name: o.EventName,
		},
		Name:   o.Name,
		IsOpen: o.IsOpen,
	}, nil
}
