package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/response"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

const ResponseCollectionName = "Responses"

type (
	Response struct {
		ID       user.ID    `json:"-"`
		OrgID    int64      `json:"org_id"`
		AuthorID int64      `json:"author_id"`
		FormID   int64      `json:"form_id"`
		Data     [][]string `json:"data"`
	}
)

func (r Response) ToModel() (*response.Response, error) {
	return &response.Response{
		ID:       r.ID,
		OrgID:    identity.NewID(r.OrgID),
		AuthorID: identity.NewID(r.AuthorID),
		FormID:   identity.NewID(r.FormID),
		Data:     r.Data,
	}, nil
}
