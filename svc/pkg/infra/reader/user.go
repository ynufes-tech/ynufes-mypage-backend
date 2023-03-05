package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"fmt"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/user"
)

type (
	User struct {
		ref *db.Ref
	}
)

func NewUser(c *firebase.Firebase) User {
	return User{
		ref: c.Client(entity.UserRootName),
	}
}

func (u User) GetByID(ctx context.Context, id id.UserID) (*user.User, error) {
	if !id.HasValue() {
		return nil, exception.ErrIDNotAssigned
	}
	r, err := u.ref.OrderByKey().EqualTo(id.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	if len(r) > 1 {
		fmt.Printf("multiple user found with id: %s", id)
	}
	var userEntity entity.User
	if err := r[0].Unmarshal(&userEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user entity: %w", err)
	}
	userEntity.ID = id
	fmt.Printf("userEntity: %+v\n", userEntity)
	model, err := userEntity.ToModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert user entity to model: %w", err)
	}
	return model, nil
}

func (u User) GetByLineServiceID(ctx context.Context, id user.LineServiceID) (*user.User, error) {
	var userEntity entity.User
	r, err := u.ref.OrderByChild("line/id").
		EqualTo(string(id)).LimitToFirst(1).
		GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(r) == 0 {
		return nil, exception.ErrNotFound
	}
	if len(r) > 1 {
		fmt.Printf("multiple user found with line service id: %s\n", id)
	}
	if err := r[0].Unmarshal(&userEntity); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user entity: %w", err)
	}
	userEntity.ID, err = identity.ImportID(r[0].Key())
	if err != nil {
		return nil, err
	}
	model, err := userEntity.ToModel()
	if err != nil {
		return nil, err
	}
	return model, nil
}
