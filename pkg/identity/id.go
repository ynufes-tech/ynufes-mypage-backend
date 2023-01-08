package identity

import (
	"strconv"
	"ynufes-mypage-backend/pkg/snowflake"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	ID        snowflake.Snowflake
	IDManager struct {
	}
)

func NewIDManager() IDManager {
	return IDManager{}
}

func (IDManager) IssueID() util.ID {
	return ID(snowflake.NewSnowflake())
}

func (IDManager) ImportID(id string) (util.ID, error) {
	result, err := strconv.ParseInt(id, 36, 64)
	if err != nil {
		return ID(0), err
	}
	return ID(result), nil
}

func NewID(id int64) ID {
	return ID(id)
}

func (i ID) ExportID() string {
	return strconv.FormatInt(int64(i), 36)
}

func (i ID) HasValue() bool {
	return i != 0
}
