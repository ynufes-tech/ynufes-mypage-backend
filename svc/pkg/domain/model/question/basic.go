package question

import "ynufes-mypage-backend/svc/pkg/domain/model/id"

type Basic struct {
	ID      id.QuestionID
	Text    string
	EventID id.EventID
	qType   Type
}

func NewBasic(id id.QuestionID, text string, eventID id.EventID, qType Type) Basic {
	return Basic{
		ID:      id,
		Text:    text,
		EventID: eventID,
		qType:   qType,
	}
}

func (b Basic) GetID() id.QuestionID {
	return b.ID
}

func (b Basic) GetText() string {
	return b.Text
}

func (b Basic) GetEventID() id.EventID {
	return b.EventID
}

func (b Basic) GetType() Type {
	return b.qType
}
