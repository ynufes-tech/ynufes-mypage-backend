package util

import "ynufes-mypage-backend/svc/pkg/domain/model/util"

type IDManager interface {
	ImportID(id string) (util.ID, error)
	IssueID() util.ID
}
