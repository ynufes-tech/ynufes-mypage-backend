package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
)

const OrgRootName = "Orgs"

type Org struct {
	ID        id.OrgID        `json:"-"`
	EventID   int64           `json:"event_id"`
	EventName string          `json:"event_name"`
	Name      string          `json:"name"`
	Users     map[string]bool `json:"user_ids"`
	IsOpen    bool            `json:"is_open"`
}

func (o Org) ToModel() (*org.Org, error) {
	users := make([]id.UserID, 0, len(o.Users))
	for k, v := range o.Users {
		if !v {
			continue
		}
		u, err := identity.ImportID(k)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return &org.Org{
		ID: o.ID,
		Event: event.Event{
			ID:   identity.NewID(o.EventID),
			Name: o.EventName,
		},
		Name:   o.Name,
		Users:  users,
		IsOpen: o.IsOpen,
	}, nil
}
