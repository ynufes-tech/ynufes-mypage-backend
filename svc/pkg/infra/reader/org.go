package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/org"
)

type Org struct {
	ref *db.Ref
}

func NewOrg(f *firebase.Firebase) Org {
	return Org{
		ref: f.Client(entity.OrgRootName),
	}
}

func (o Org) GetByID(ctx context.Context, oid id.OrgID) (*org.Org, error) {
	var orgEntity entity.Org
	r, err := o.ref.OrderByKey().
		EqualTo(oid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	if len(r) > 1 {
		fmt.Printf("multiple org found with id: %s\n", oid)
	}
	if err := r[0].Unmarshal(&orgEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal org entity: %w", err)
	}
	orgEntity.ID = oid
	m, err := orgEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return m, nil
}
