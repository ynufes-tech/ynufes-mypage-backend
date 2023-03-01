package question

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
)

type Basic struct {
	ID      ID
	Text    string
	EventID event.ID
	qType   Type
}

func NewBasic(id ID, text string, eventID event.ID, qType Type) Basic {
	return Basic{
		ID:      id,
		Text:    text,
		EventID: eventID,
		qType:   qType,
	}
}

func (b Basic) GetID() ID {
	return b.ID
}

func (b Basic) GetText() string {
	return b.Text
}

func (b Basic) GetEventID() event.ID {
	return b.EventID
}

func (b Basic) GetType() Type {
	return b.qType
}
