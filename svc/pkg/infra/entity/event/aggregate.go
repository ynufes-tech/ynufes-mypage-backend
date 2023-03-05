package entity

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

const EventRootName = "Events"

type Event struct {
	ID   id.EventID `json:"-"`
	Name string     `json:"name"`
}

func (e Event) ToModel() (*event.Event, error) {
	return &event.Event{
		ID:   e.ID,
		Name: e.Name,
	}, nil
}
