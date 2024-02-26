package reader

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/staff"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

func TestStaff_GetStaffByUserID(t *testing.T) {
	fb := testutil.NewFirebaseTest()

	r := NewStaff(fb.GetClient())
	w := writer.NewStaff(fb.GetClient())
	userW := writer.NewUser(fb.GetClient())
	ctx := context.Background()

	us := testutil.Users()
	for i := range us {
		err := userW.Create(ctx, &us[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	err := w.Set(ctx, staff.Staff{
		UserID:  us[0].ID,
		IsAdmin: true,
	})
	assert.NoError(t, err)

	tests := []struct {
		name    string
		uid     id.UserID
		want    *staff.Staff
		wantErr error
	}{
		{
			name: "normal",
			uid:  us[0].ID,
			want: &staff.Staff{
				UserID:  us[0].ID,
				IsAdmin: true,
			},
			wantErr: nil,
		},
		{
			name:    "not found",
			uid:     us[1].ID,
			want:    nil,
			wantErr: exception.ErrNotFound,
		},
	}

	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetStaffByUserID(ctx, tt.uid)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
