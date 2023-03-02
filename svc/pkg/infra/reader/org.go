package reader

import (
	"cloud.google.com/go/firestore"
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/org"
)

type Org struct {
	collection *firestore.CollectionRef
}

func NewOrg(c *firestore.Client) Org {
	return Org{
		collection: c.Collection(entity.OrgCollectionName),
	}
}

func (o Org) GetByID(ctx context.Context, id id.OrgID) (*org.Org, error) {
	var orgEntity entity.Org
	oid := id.ExportID()
	snap, err := o.collection.Doc(oid).Get(ctx)
	if err != nil {
		return nil, err
	}
	err = snap.DataTo(&orgEntity)
	orgEntity.ID = id
	model, err := orgEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (o Org) ListByGrantedUserID(ctx context.Context, id id.UserID) ([]org.Org, error) {
	var orgs []org.Org
	uid := id.GetValue()
	iter := o.collection.Where("user_ids", "array-contains", uid).Documents(ctx)
	for {
		var orgEntity entity.Org
		snap, err := iter.Next()
		if err != nil {
			break
		}
		err = snap.DataTo(&orgEntity)
		if err != nil {
			return nil, err
		}
		orgEntity.ID = id
		model, err := orgEntity.ToModel()
		if err != nil {
			return nil, err
		}
		orgs = append(orgs, *model)
	}
	return orgs, nil
}
