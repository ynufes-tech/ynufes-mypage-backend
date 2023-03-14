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
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
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
		Roles: []user.RoleID{
			user.RoleID(identity.IssueID()),
			user.RoleID(identity.IssueID()),
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
	assert.Equal(t, form1, *f)
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
		Roles: []user.RoleID{
			user.RoleID(identity.IssueID()),
			user.RoleID(identity.IssueID()),
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
		Roles: []user.RoleID{
			user.RoleID(identity.IssueID()),
			user.RoleID(identity.IssueID()),
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
		Roles: []user.RoleID{
			user.RoleID(identity.IssueID()),
			user.RoleID(identity.IssueID()),
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
	events1, err := formR.ListByEventID(ctx, event1.ID)
	assert.NoError(t, err)
	events2, err := formR.ListByEventID(ctx, event2.ID)
	assert.NoError(t, err)
	assert.Equal(t, []form.Form{form1, form2}, events1)
	assert.Equal(t, []form.Form{form3}, events2)
}
