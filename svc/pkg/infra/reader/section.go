package reader

import (
	"context"
	"firebase.google.com/go/v4/db"
	"ynufes-mypage-backend/pkg/firebase"
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

func (s Section) GetSectionByID(ctx context.Context, tid id.SectionID) (*section.Section, error) {
	if tid == nil || !tid.HasValue() {
		return nil, exception.ErrIDNotAssigned
	}
	hits, err := s.ref.OrderByKey().
		EqualTo(tid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	if len(hits) == 0 {
		return nil, exception.ErrNotFound
	}
	var e entity.Section
	if err := hits[0].Unmarshal(&e); err != nil {
		return nil, err
	}
	e.ID = tid
	m, err := e.ToModel()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (s Section) ListSectionsByFormID(ctx context.Context, fid id.FormID) ([]section.Section, error) {
	if fid == nil || !fid.HasValue() {
		return nil, exception.ErrIDNotAssigned
	}
	hits, err := s.ref.OrderByChild("form_id").
		EqualTo(fid.ExportID()).GetOrdered(ctx)
	if err != nil {
		return nil, err
	}
	var sections []section.Section
	for _, hit := range hits {
		var e entity.Section
		if err := hit.Unmarshal(&e); err != nil {
			return nil, err
		}
		e.ID = fid
		m, err := e.ToModel()
		if err != nil {
			return nil, err
		}
		sections = append(sections, *m)
	}
	return sections, nil
}
