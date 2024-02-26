package writer

import (
	"context"
	"encoding/json"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/user"
)

type User struct {
	ref *db.Ref
}

func NewUser(f *firebase.Firebase) User {
	return User{
		ref: f.Client(entity.UserRootName),
	}
}

// Create Note that new user will be created with no roles or authority.
func (u User) Create(ctx context.Context, model *user.User) error {
	if model.ID != nil && model.ID.HasValue() {
		return exception.ErrIDAlreadyAssigned
	}
	tid := identity.IssueID()
	e := entity.User{
		UserDetail: entity.UserDetail{
			NameFirst:     model.Detail.Name.FirstName,
			NameFirstKana: model.Detail.Name.FirstNameKana,
			NameLast:      model.Detail.Name.LastName,
			NameLastKana:  model.Detail.Name.LastNameKana,
			Gender:        int(model.Detail.Gender),
			StudentID:     string(model.Detail.StudentID),
			Email:         string(model.Detail.Email),
			Type:          int(model.Detail.Type),
			PictureURL:    string(model.Detail.PictureURL),
		},
	}
	err := u.ref.Child(tid.ExportID()).
		Set(ctx, e)
	if err != nil {
		return err
	}
	model.ID = tid
	return nil
}

func (u User) Set(ctx context.Context, model user.User) error {
	if model.ID == nil || !model.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}

	e := entity.User{
		UserDetail: entity.UserDetail{
			NameFirst:     model.Detail.Name.FirstName,
			NameFirstKana: model.Detail.Name.FirstNameKana,
			NameLast:      model.Detail.Name.LastName,
			NameLastKana:  model.Detail.Name.LastNameKana,
			Gender:        int(model.Detail.Gender),
			StudentID:     string(model.Detail.StudentID),
			Email:         string(model.Detail.Email),
			Type:          int(model.Detail.Type),
		},
	}
	err := u.ref.Child(model.ID.ExportID()).
		Set(ctx, e)
	return err
}

func (u User) UpdateUserDetail(ctx context.Context, tID id.UserID, detail user.Detail) error {
	// entity.UserDetail have optional fields,
	// so empty value will not be updated.
	e := entity.UserDetail{
		NameFirst:     detail.Name.FirstName,
		NameFirstKana: detail.Name.FirstNameKana,
		NameLast:      detail.Name.LastName,
		NameLastKana:  detail.Name.LastNameKana,
		Gender:        int(detail.Gender),
		StudentID:     string(detail.StudentID),
		Email:         string(detail.Email),
		Type:          int(detail.Type),
		PictureURL:    string(detail.PictureURL),
	}
	// marshal to json and unmarshal to map[string]interface{}
	// so that empty value will not be updated.
	jsonStr, err := json.Marshal(e)
	if err != nil {
		return err
	}
	var m map[string]interface{}
	if err = json.Unmarshal(jsonStr, &m); err != nil {
		return err
	}
	if err := u.ref.Child(tID.ExportID()).Child("detail").
		Update(ctx, m); err != nil {
		return err
	}
	return nil
}

func (u User) Delete(ctx context.Context, model user.User) error {
	err := u.ref.Child(model.ID.ExportID()).
		Delete(ctx)
	return err
}
