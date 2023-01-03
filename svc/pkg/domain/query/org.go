package query

import (
	"ynufes-mypage-backend/pkg/snowflake"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
)

type Org interface {
	GetByID(id snowflake.Snowflake) (*org.Org, error)
}
