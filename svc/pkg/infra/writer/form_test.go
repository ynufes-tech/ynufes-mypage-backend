package writer

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
)

func TestForm_SectionOrder(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	w := NewForm(fb.GetClient())
	targetForm := form.Form{
		EventID:     identity.IssueID(),
		Title:       "title",
		Summary:     "summary",
		Description: "test",
		Roles:       nil,
		Deadline:    time.Now().Add(time.Hour * 24 * 7),
		IsOpen:      false,
		Sections:    nil,
	}
	err := w.Create(context.Background(), &targetForm)
	assert.NoErrorf(t, err, "failed to create form: %v", err)

	targetSection := struct {
		Index     float64
		SectionID id.SectionID
	}{
		Index:     0.5,
		SectionID: identity.IssueID(),
	}
	err = w.AddSectionOrder(context.Background(), targetForm.ID, targetSection.SectionID, targetSection.Index)
	assert.NoErrorf(t, err, "failed to add section order: %v", err)

	targetSection1 := struct {
		Index     float64
		SectionID id.SectionID
	}{
		Index:     0.5,
		SectionID: targetSection.SectionID,
	}
	err = w.AddSectionOrder(context.Background(), targetForm.ID, targetSection1.SectionID, targetSection1.Index)
	assert.ErrorIs(t, err, exception.ErrAlreadyExists)
	fmt.Println("err in add section order: ", err)

	updateCases := []struct {
		Name      string
		Index     float64
		SectionID id.SectionID
		hasError  bool
	}{
		{
			Name:      "(Successful) Update section order",
			Index:     0.2,
			SectionID: targetSection.SectionID,
			hasError:  false,
		},
		{
			Name:      "(ErrorCase) Update section order",
			Index:     0.2,
			SectionID: identity.IssueID(),
			hasError:  true,
		},
	}
	for _, tt := range updateCases {
		err = w.UpdateSectionOrder(context.Background(), targetForm.ID, tt.SectionID, tt.Index)
		if tt.hasError {
			assert.ErrorIs(t, err, exception.ErrNotFound)
		} else {
			assert.NoErrorf(t, err, "illgal error in %s: %v", tt.Name, err)
		}
	}
}
