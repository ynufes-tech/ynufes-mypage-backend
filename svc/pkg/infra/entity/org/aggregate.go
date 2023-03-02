package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
)

const OrgCollectionName = "Orgs"

type Org struct {
	ID        id.OrgID `json:"-"`
	EventID   int64    `json:"event_id"`
	EventName string   `json:"event_name"`
	Name      string   `json:"name"`
	Users     []int64  `json:"user_ids"`
	IsOpen    bool     `json:"is_open"`
}

func (o Org) ToModel() (*org.Org, error) {
	var users []id.UserID
	for i := range o.Users {
		users = append(users, identity.NewID(o.Users[i]))
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
