package writer

import (
	"cloud.google.com/go/firestore"
	"context"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	entity "ynufes-mypage-backend/svc/pkg/infra/entity/form"
)

type Form struct {
	collection *firestore.CollectionRef
}

func NewForm(client *firestore.Client) *Form {
	return &Form{
		collection: client.Collection(entity.FormCollectionName),
	}
}

func (f Form) Create(ctx context.Context, target form.Form) error {
	sectionOrder := make([]int64, len(target.SectionIDs))
	for i := range target.SectionIDs {
		sectionOrder[i] = target.SectionIDs[i].GetValue()
	}

	sections := make([]entity.Section, 0, len(target.Sections))
	for id, section := range target.Sections {
		qIDs := make([]int64, len(section.QuestionIDs))
		for i := range section.QuestionIDs {
			qIDs[i] = section.QuestionIDs[i].GetValue()
		}
		cCustoms := make(map[string]int64, len(section.ConditionCustoms))
		for k, v := range section.ConditionCustoms {
			cCustoms[k.ExportID()] = v.GetValue()
		}
		sections = append(sections, entity.NewSection(
			id.GetValue(),
			qIDs,
			section.ConditionQuestion.GetValue(),
			cCustoms,
		))
	}

	formID := target.ID.ExportID()
	var roles = make([]int64, len(target.Roles))
	for i := 0; i < len(target.Roles); i++ {
		roles[i] = target.Roles[i].GetValue()
	}
	e := entity.NewForm(
		formID,
		target.EventID.GetValue(),
		target.Title,
		target.Summary,
		target.Description,
		roles,
		target.Deadline.UnixMilli(),
		target.IsOpen,
		sectionOrder,
		sections,
	)
	_, err := f.collection.Doc(formID).Create(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (f Form) Set(ctx context.Context, target form.Form) error {
	formID := target.ID.ExportID()
	var roles = make([]int64, len(target.Roles))
	for i := 0; i < len(target.Roles); i++ {
		roles[i] = target.Roles[i].GetValue()
	}
	sectionOrder := make([]int64, len(target.SectionIDs))
	for i := range target.SectionIDs {
		sectionOrder[i] = target.SectionIDs[i].GetValue()
	}

	sections := make([]entity.Section, 0, len(target.Sections))
	for id, section := range target.Sections {
		qIDs := make([]int64, len(section.QuestionIDs))
		for i := range section.QuestionIDs {
			qIDs[i] = section.QuestionIDs[i].GetValue()
		}
		cCustoms := make(map[string]int64, len(section.ConditionCustoms))
		for k, v := range section.ConditionCustoms {
			cCustoms[k.ExportID()] = v.GetValue()
		}
		sections = append(sections, entity.NewSection(
			id.GetValue(),
			qIDs,
			section.ConditionQuestion.GetValue(),
			cCustoms,
		))
	}
	e := entity.NewForm(
		formID,
		target.EventID.GetValue(),
		target.Title,
		target.Summary,
		target.Description,
		roles,
		target.Deadline.UnixMilli(),
		target.IsOpen,
		sectionOrder,
		sections,
	)
	_, err := f.collection.Doc(formID).Set(ctx, e)
	if err != nil {
		return err
	}
	return nil
}
