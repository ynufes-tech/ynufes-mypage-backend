package reader

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/form"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

func TestForm_GetByID(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	formW := writer.NewForm(fb.GetClient())
	formR := NewForm(fb.GetClient())

	ctx := context.Background()
	event1 := event.Event{
		Name: "TestEvent1",
	}
	eventW := writer.NewEvent(fb.GetClient())
	assert.NoError(t, eventW.Create(ctx, &event1))

	deadline := time.Now().Add(time.Hour * 24 * 7).UnixMilli()
	form1 := form.Form{
		EventID:     event1.ID,
		Title:       "FormTitle1",
		Summary:     "FormSummary1",
		Description: "FormDescription1",
		Roles: []id.RoleID{
			id.RoleID(identity.IssueID()),
			id.RoleID(identity.IssueID()),
		},
		Deadline: time.UnixMilli(deadline),
		IsOpen:   false,
		// Index of sections will be reordered and reassigned on creation.
		Sections: map[id.SectionID]float64{
			id.SectionID(identity.IssueID()): 3,
			id.SectionID(identity.IssueID()): 2,
			id.SectionID(identity.IssueID()): 1,
			id.SectionID(identity.IssueID()): 4,
			id.SectionID(identity.IssueID()): 0,
		},
	}

	assert.NoError(t, formW.Create(ctx, &form1))
	f, err := formR.GetByID(ctx, form1.ID)
	if err != nil {
		return
	}
	checkFormEqual(t, form1, *f)
}

func TestForm_ListByEventID(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	formW := writer.NewForm(fb.GetClient())
	formR := NewForm(fb.GetClient())

	ctx := context.Background()
	event1 := event.Event{
		Name: "TestEvent1",
	}
	event2 := event.Event{
		Name: "TestEvent2",
	}
	eventW := writer.NewEvent(fb.GetClient())
	assert.NoError(t, eventW.Create(ctx, &event1))
	assert.NoError(t, eventW.Create(ctx, &event2))

	deadline1 := time.Now().Add(time.Hour * 24 * 7).UnixMilli()
	deadline2 := time.Now().Add(time.Hour * 24 * 2).UnixMilli()
	deadline3 := time.Now().Add(time.Hour * 24 * 3).UnixMilli()
	form1 := form.Form{
		EventID:     event1.ID,
		Title:       "FormTitle1",
		Summary:     "FormSummary1",
		Description: "FormDescription1",
		Roles: []id.RoleID{
			id.RoleID(identity.IssueID()),
			id.RoleID(identity.IssueID()),
		},
		Deadline: time.UnixMilli(deadline1),
		IsOpen:   false,
		// Index of sections will be reordered and reassigned on creation.
		Sections: map[id.SectionID]float64{
			id.SectionID(identity.IssueID()): 1,
			id.SectionID(identity.IssueID()): 4,
			id.SectionID(identity.IssueID()): 2,
			id.SectionID(identity.IssueID()): 3,
			id.SectionID(identity.IssueID()): 0,
		},
	}
	form2 := form.Form{
		EventID:     event1.ID,
		Title:       "FormTitle2",
		Summary:     "FormSummary2",
		Description: "FormDescription2",
		Roles: []id.RoleID{
			id.RoleID(identity.IssueID()),
			id.RoleID(identity.IssueID()),
		},
		Deadline: time.UnixMilli(deadline2),
		IsOpen:   false,
		// Index of sections will be reordered and reassigned on creation.
		Sections: map[id.SectionID]float64{
			id.SectionID(identity.IssueID()): 4,
			id.SectionID(identity.IssueID()): 3,
			id.SectionID(identity.IssueID()): 0,
			id.SectionID(identity.IssueID()): 1,
			id.SectionID(identity.IssueID()): 2,
		},
	}
	form3 := form.Form{
		EventID:     event2.ID,
		Title:       "FormTitle3",
		Summary:     "FormSummary3",
		Description: "FormDescription3",
		Roles: []id.RoleID{
			id.RoleID(identity.IssueID()),
			id.RoleID(identity.IssueID()),
		},
		Deadline: time.UnixMilli(deadline3),
		IsOpen:   false,
		// Index of sections will be reordered and reassigned on creation.
		Sections: map[id.SectionID]float64{
			id.SectionID(identity.IssueID()): 3,
			id.SectionID(identity.IssueID()): 2,
			id.SectionID(identity.IssueID()): 1,
			id.SectionID(identity.IssueID()): 4,
			id.SectionID(identity.IssueID()): 0,
		},
	}
	assert.NoError(t, formW.Create(ctx, &form1))
	assert.NoError(t, formW.Create(ctx, &form2))
	assert.NoError(t, formW.Create(ctx, &form3))
	forms1, err := formR.ListByEventID(ctx, event1.ID)
	assert.NoError(t, err)
	forms2, err := formR.ListByEventID(ctx, event2.ID)
	assert.NoError(t, err)
	checkFormsEqual(t, []form.Form{form1, form2}, forms1)
	checkFormsEqual(t, []form.Form{form3}, forms2)
}

func checkFormsEqual(t *testing.T, f1, f2 []form.Form) {
	assert.Equal(t, len(f1), len(f2))
	forms := make(map[id.FormID]form.Form, len(f1))
	for _, v := range f1 {
		forms[v.ID] = v
	}
	for _, v := range f2 {
		checkFormEqual(t, v, forms[v.ID])
	}
}

func checkFormEqual(t *testing.T, f1, f2 form.Form) {
	assert.Equal(t, f1.ID, f2.ID)
	assert.Equal(t, f1.EventID, f2.EventID)
	assert.Equal(t, f1.Title, f2.Title)
	assert.Equal(t, f1.Summary, f2.Summary)
	assert.Equal(t, f1.Description, f2.Description)
	checkRolesEqual(t, f1.Roles, f2.Roles)
	assert.Equal(t, f1.Deadline, f2.Deadline)
	assert.Equal(t, f1.IsOpen, f2.IsOpen)
	assert.Equal(t, len(f1.Sections), len(f2.Sections))
	for k, v := range f1.Sections {
		assert.Equal(t, v, f2.Sections[k])
	}
}

func checkRolesEqual(t *testing.T, r1, r2 []id.RoleID) {
	assert.Equal(t, len(r1), len(r2))
	roles := make(map[id.RoleID]struct{}, len(r1))
	for _, v := range r1 {
		roles[v] = struct{}{}
	}
	for _, v := range r2 {
		_, ok := roles[v]
		assert.True(t, ok)
	}
}
