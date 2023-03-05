package writer

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/section"
	"ynufes-mypage-backend/svc/pkg/exception"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/section"
)

type Section struct {
	ref *db.Ref
}

func NewSection(client *firebase.Firebase) Section {
	return Section{
		ref: client.Client(entity.SectionRootName),
	}
}

func (s Section) Create(ctx context.Context, targetSection *section.Section) error {
	if targetSection.ID != nil && targetSection.ID.HasValue() {
		return exception.ErrIDAlreadyAssigned
	}
	newID := identity.IssueID()

	questions := make(map[string]float64, len(targetSection.QuestionIDs))
	for qid, index := range targetSection.QuestionIDs {
		questions[qid.ExportID()] = index
	}

	customs := make(map[string]string, len(targetSection.ConditionCustoms))
	for key, value := range targetSection.ConditionCustoms {
		customs[key.ExportID()] = value.ExportID()
	}

	err := s.ref.Child(newID.ExportID()).Set(ctx,
		entity.NewSection(
			newID,
			targetSection.FormID,
			questions,
			targetSection.ConditionQuestion.ExportID(),
			customs,
		),
	)
	if err != nil {
		return err
	}
	targetSection.ID = newID
	return nil
}

func (s Section) Set(ctx context.Context, targetSection section.Section) error {
	if targetSection.ID == nil || !targetSection.ID.HasValue() {
		return exception.ErrIDNotAssigned
	}

	questions := make(map[string]float64, len(targetSection.QuestionIDs))
	for qid, index := range targetSection.QuestionIDs {
		questions[qid.ExportID()] = index
	}

	customs := make(map[string]string, len(targetSection.ConditionCustoms))
	for key, value := range targetSection.ConditionCustoms {
		customs[key.ExportID()] = value.ExportID()
	}

	err := s.ref.Child(targetSection.ID.ExportID()).Set(ctx,
		entity.NewSection(
			targetSection.ID,
			targetSection.FormID,
			questions,
			targetSection.ConditionQuestion.ExportID(),
			customs,
		),
	)
	if err != nil {
		return err
	}
	return nil
}

func (s Section) LinkQuestion(ctx context.Context, secID id.SectionID, qID id.QuestionID, pos float64) error {
	if qID == nil || !qID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	if err := s.ref.Child(secID.ExportID()).Child("questions").
		Child(qID.ExportID()).Set(ctx, pos); err != nil {
		return err
	}
	return nil
}

func (s Section) UnlinkQuestion(ctx context.Context, secID id.SectionID, qID id.QuestionID) error {
	if qID == nil || !qID.HasValue() {
		return exception.ErrIDNotAssigned
	}
	if err := s.ref.Child(secID.ExportID()).Child("questions").
		Child(qID.ExportID()).Delete(ctx); err != nil {
		return err
	}
	return nil
}
