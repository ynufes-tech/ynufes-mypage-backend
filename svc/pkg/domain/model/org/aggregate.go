package org

import "ynufes-mypage-backend/pkg/snowflake"

type (
	Org struct {
		ID      ID
		EventID snowflake.Snowflake
		Name    string
	}
	ID snowflake.Snowflake
)
