package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/line"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/line"
)

type Line struct {
	ref *db.Ref
}

func NewLine(fb *firebase.Firebase) Line {
	return Line{
		ref: fb.Client(entity.LineRootName),
	}
}

func (w Line) Create(ctx context.Context, lineUser line.LineUser) error {
	if lineUser.LineServiceID == "" || lineUser.UserID == nil || !lineUser.UserID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	err := w.ref.Child(string(lineUser.LineServiceID)).
		Transaction(ctx, func(node db.TransactionNode) (interface{}, error) {
			var e entity.Line
			if err := node.Unmarshal(&e); err != nil {
				return nil, err
			}
			if e.UserID != "" {
				return nil, exception.ErrAlreadyExists
			}
			e = entity.Line{
				LineServiceID:         string(lineUser.LineServiceID),
				UserID:                lineUser.UserID.ExportID(),
				LineDisplayName:       lineUser.LineDisplayName,
				EncryptedAccessToken:  string(lineUser.EncryptedAccessToken),
				EncryptedRefreshToken: string(lineUser.EncryptedRefreshToken),
			}
			return e, nil
		})
	if err != nil {
		return err
	}
	return nil
}

func (w Line) Set(ctx context.Context, lineUser line.LineUser) error {
	if lineUser.UserID == nil || !lineUser.UserID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	if lineUser.LineServiceID == "" {
		return exception.ErrIDNotAssigned
	}
	err := w.ref.Child(string(lineUser.LineServiceID)).
		Set(ctx, entity.Line{
			LineServiceID:         string(lineUser.LineServiceID),
			UserID:                lineUser.UserID.ExportID(),
			LineDisplayName:       lineUser.LineDisplayName,
			EncryptedAccessToken:  string(lineUser.EncryptedAccessToken),
			EncryptedRefreshToken: string(lineUser.EncryptedRefreshToken),
		})
	if err != nil {
		return err
	}
	return err
}
