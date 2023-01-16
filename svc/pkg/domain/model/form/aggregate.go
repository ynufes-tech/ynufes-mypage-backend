package form

import "ynufes-mypage-backend/svc/pkg/domain/model/util"

type (
	Form struct {
		ID        ID
		Questions map[QID]Question
	}
	ID util.ID
)
