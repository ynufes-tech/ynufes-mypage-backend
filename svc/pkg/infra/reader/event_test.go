package reader

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

func TestEvent_GetByID(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	w := writer.NewEvent(fb.GetClient())
	r := NewEvent(fb.GetClient())
	eventExists := &event.Event{
		Name: "TestEvent",
	}
	assert.NoError(t, w.Create(context.Background(), eventExists))
	tests := []struct {
		name    string
		query   id.EventID
		want    *event.Event
		wantErr error
	}{
		{
			name:    "Success",
			query:   eventExists.ID,
			want:    eventExists,
			wantErr: nil,
		},
		{
			name:    "NotFound",
			query:   identity.IssueID(),
			want:    nil,
			wantErr: exception.ErrNotFound,
		},
		{
			name:    "NO ID",
			query:   nil,
			want:    nil,
			wantErr: exception.ErrIDNotAssigned,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetByID(context.Background(), tt.query)
			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
