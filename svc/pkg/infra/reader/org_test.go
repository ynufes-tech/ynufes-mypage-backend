package reader

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

func TestOrg_GetByID(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	eventW := writer.NewEvent(fb.GetClient())

	event1 := event.Event{
		Name: "Event1",
	}
	assert.NoError(t, eventW.Create(context.Background(), &event1))

	org1 := org.Org{
		Name:   "test",
		IsOpen: false,
		Event:  event1,
	}
	orgW := writer.NewOrg(fb.GetClient())
	assert.NoError(t, orgW.Create(context.Background(), &org1))

	tests := []struct {
		name    string
		query   id.OrgID
		want    *org.Org
		wantErr error
	}{
		{
			name:    "Success",
			query:   org1.ID,
			want:    &org1,
			wantErr: nil,
		},
		{
			name:    "NotFound",
			want:    nil,
			query:   identity.IssueID(),
			wantErr: exception.ErrNotFound,
		},
	}
	r := NewOrg(fb.GetClient())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := r.GetByID(ctx, tt.query)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, got, tt.want)
		})
	}
}
