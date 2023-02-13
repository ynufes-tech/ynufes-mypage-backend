package entity

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/response"
)

const ResponseCollectionName = "Responses"

type (
	Response struct {
		ID       string     `firestore:"-"`
		OrgID    int64      `firestore:"org_id"`
		AuthorID int64      `firestore:"author_id"`
		FormID   int64      `firestore:"form_id"`
		Data     [][]string `firestore:"data"`
	}
)

func (r Response) ToModel() (*response.Response, error) {
	id, err := identity.ImportID(r.ID)
	if err != nil {
		return nil, err
	}
	return &response.Response{
		ID:       id,
		OrgID:    identity.NewID(r.OrgID),
		AuthorID: identity.NewID(r.AuthorID),
		FormID:   identity.NewID(r.FormID),
		Data:     r.Data,
	}, nil
}
