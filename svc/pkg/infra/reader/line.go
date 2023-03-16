package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
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

func (l Line) GetByUserID(ctx context.Context, userID id.UserID) (*line.LineUser, error) {
	if userID == nil || !userID.HasValue() {
		return nil, exception.ErrIDNotAssigned
	}
	r, err := l.ref.OrderByChild("user_id").EqualTo(userID.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	if len(r) > 1 {
		fmt.Printf("multiple line found with user id: %s", userID)
	}
	var lineEntity entity.Line
	if err := r[0].Unmarshal(&lineEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal line entity: %w", err)
	}
	lineEntity.LineServiceID = r[0].Key()
	model, err := lineEntity.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert line entity to model: %w", err)
	}
	return model, nil
}

func (l Line) GetByLineServiceID(ctx context.Context, lineID line.LineServiceID) (*line.LineUser, error) {
	if lineID == "" {
		return nil, exception.ErrIDNotAssigned
	}
	r, err := l.ref.OrderByKey().EqualTo(lineID).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	var lineEntity entity.Line
	if err := r[0].Unmarshal(&lineEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal line entity: %w", err)
	}
	lineEntity.LineServiceID = string(lineID)
	model, err := lineEntity.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert line entity to model: %w", err)
	}
	return model, nil
}
