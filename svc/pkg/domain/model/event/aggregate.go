package event

import "ynufes-mypage-backend/pkg/snowflake"

type (
	Event struct {
		ID   ID
		Name string
	}
	ID snowflake.Snowflake
)
