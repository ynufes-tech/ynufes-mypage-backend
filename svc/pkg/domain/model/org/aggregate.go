package org

import "ynufes-mypage-backend/pkg/snowflake"

type Org struct {
	ID      snowflake.Snowflake
	EventID snowflake.Snowflake
	Name    string
}
