package response

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	Response struct {
		ID       ID
		OrgID    id.OrgID
		AuthorID id.UserID
		FormID   id.FormID
		Data     map[string]interface{}
	}
	ID util.ID
)
