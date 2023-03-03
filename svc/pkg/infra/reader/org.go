package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
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

func (o Org) GetByID(ctx context.Context, id id.OrgID) (*org.Org, error) {
	var orgEntity entity.Org
	oid := id.ExportID()
	err := o.ref.Child(oid).Get(ctx, &orgEntity)
	orgEntity.ID = id
	model, err := orgEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (o Org) ListByGrantedUserID(ctx context.Context, id id.UserID) ([]org.Org, error) {
	hits, err := o.ref.OrderByChild("user_ids/" + id.ExportID()).
		EqualTo(true).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	orgs := make([]org.Org, len(hits))
	for i := range hits {
		var orgEntity entity.Org
		if err := hits[i].Unmarshal(&orgEntity); err != nil {
			return nil, err
		}
		oid, err := identity.ImportID(hits[i].Key())
		if err != nil {
			return nil, err
		}
		orgEntity.ID = oid
		model, err := orgEntity.ToModel()
		if err != nil {
			return nil, err
		}
		orgs[i] = *model
	}
	return orgs, nil
}
