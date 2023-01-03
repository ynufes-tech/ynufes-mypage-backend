package event

import "ynufes-mypage-backend/pkg/snowflake"

type (
	Event struct {
		ID   snowflake.Snowflake
		Name string
	}
)
