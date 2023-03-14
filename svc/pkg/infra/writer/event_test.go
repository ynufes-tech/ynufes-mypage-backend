package writer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
)

func TestEvent_Create(t *testing.T) {
	tests := []struct {
		name    string
		target  *event.Event
		wantErr bool
	}{
		{
			name: "Success",
			target: &event.Event{
				Name: "TestEvent",
			},
			wantErr: false,
		},
	}
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	w := NewEvent(fb.GetClient())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr {
				assert.NoError(t, w.Create(context.Background(), tt.target))
			} else {
				assert.Error(t, w.Create(context.Background(), tt.target))
			}
		})
	}
	r := reader.NewEvent(fb.GetClient())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.wantErr {
				got, err := r.GetByID(context.Background(), tt.target.ID)
				assert.NoError(t, err)
				assert.Equal(t, tt.target, got)
			}
		})
	}
}

func TestEvent_Set(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	w := NewEvent(fb.GetClient())
	r := reader.NewEvent(fb.GetClient())
	eventExists := &event.Event{
		Name: "TestEvent",
	}
	assert.NoError(t, w.Create(context.Background(), eventExists))

	tests := []struct {
		name    string
		target  event.Event
		wantErr error
	}{
		{
			name: "CaseNew",
			target: event.Event{
				ID:   identity.IssueID(),
				Name: "TestEvent",
			},
			wantErr: nil,
		},
		{
			name: "CaseUpdate",
			target: event.Event{
				ID:   eventExists.ID,
				Name: "TestEvent2",
			},
			wantErr: nil,
		},
		{
			name: "CaseIDNotAssigned",
			target: event.Event{
				ID:   nil,
				Name: "TestEvent",
			},
			wantErr: exception.ErrIDNotAssigned,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				assert.NoError(t, w.Set(context.Background(), tt.target))
			} else {
				assert.Error(t, w.Set(context.Background(), tt.target))
			}
		})
	}
	for _, tt := range tests {
		if tt.wantErr == nil {
			rd, err := r.GetByID(context.Background(), tt.target.ID)
			assert.NoError(t, err)
			assert.Equal(t, tt.target, *rd)
		}
	}
}

func TestEvent_UpdateName(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	w := NewEvent(fb.GetClient())
	r := reader.NewEvent(fb.GetClient())
	eventExists := &event.Event{
		Name: "TestEvent",
	}
	assert.NoError(t, w.Create(context.Background(), eventExists))

	tests := []struct {
		id       id.EventID
		name     string
		updateTo string
		wantErr  error
	}{
		{
			name:     "CaseUpdate",
			id:       eventExists.ID,
			updateTo: "TestEventUpdated",
			wantErr:  nil,
		},
		{
			name:     "CaseIDNotAssigned",
			id:       nil,
			updateTo: "TestEventUpdated",
			wantErr:  exception.ErrIDNotAssigned,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr == nil {
				assert.NoError(t, w.UpdateName(context.Background(), tt.id, tt.updateTo))
			} else {
				assert.Error(t, w.UpdateName(context.Background(), tt.id, tt.updateTo))
			}
		})
	}
	for _, tt := range tests {
		if tt.wantErr == nil {
			rd, err := r.GetByID(context.Background(), tt.id)
			assert.NoError(t, err)
			assert.Equal(t, tt.updateTo, rd.Name)
		}
	}
}
