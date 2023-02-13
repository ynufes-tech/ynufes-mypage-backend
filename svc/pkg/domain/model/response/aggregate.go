package response

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/domain/model/util"
)

type (
	Response struct {
		ID       ID
		OrgID    org.ID
		AuthorID user.ID
		FormID   form.ID
		Data     [][]string
	}
	ID util.ID
)
