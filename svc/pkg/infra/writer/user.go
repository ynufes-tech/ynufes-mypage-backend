package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"strconv"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/user"
)

type User struct {
	Collection *firestore.CollectionRef
}

func NewUser(c *firestore.CollectionRef) User {
	return User{
		Collection: c,
	}
}

func (u User) Create(ctx context.Context, model user.User) error {
	e := entity.User{
		//ID is not required as it will not be used by firestore
		//ID: int64(model.ID),
		UserDetail: entity.UserDetail{
			NameFirst:     model.Detail.Name.FirstName,
			NameFirstKana: model.Detail.Name.FirstNameKana,
			NameLast:      model.Detail.Name.LastName,
			NameLastKana:  model.Detail.Name.LastNameKana,
			Gender:        int(model.Detail.Gender),
			StudentID:     string(model.Detail.StudentID),
			Email:         string(model.Detail.Email),
			Type:          model.Detail.Type,
		},
		LineAuthorization: entity.LineAuthorization{
			LineServiceID:         model.Line.LineServiceID,
			EncryptedAccessToken:  model.Line.EncryptedAccessToken,
			EncryptedRefreshToken: model.Line.EncryptedRefreshToken,
		},
		UserDashboard: entity.UserDashboard{
			Grants: model.Dashboard.Grants,
		},
	}
	//NOTE: Create fails if the document already exists
	_, err := u.Collection.Doc(strconv.FormatInt(int64(model.ID), 10)).
		Create(ctx, e)
	if err != nil {
		return err
	}
	return nil
}