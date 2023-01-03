package query

import (
	"ynufes-mypage-backend/pkg/snowflake"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
)

type Event interface {
	GetByID(id snowflake.Snowflake) (*event.Event, error)
}
