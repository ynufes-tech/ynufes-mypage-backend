package writer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/staff"
)

func TestStaff_Set(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	us := testutil.Users()
	c := fb.GetClient()
	defer fb.Reset()

	userc := NewUser(c)
	ctx := context.Background()
	for i := range us {
		err := userc.Create(ctx, &us[i])
		assert.NoError(t, err)
	}

	tests := []struct {
		name    string
		staff   staff.Staff
		wantErr error
	}{
		{
			name: "normal",
			staff: staff.Staff{
				UserID:  us[0].ID,
				IsAdmin: false,
			},
			wantErr: nil,
		},
		{
			name: "normal as admin",
			staff: staff.Staff{
				UserID:  us[1].ID,
				IsAdmin: true,
			},
			wantErr: nil,
		},
		{
			// this testcase is for clarifying the writer spec.
			// implement constraint would be better,
			// but it also would worsen the performance.
			name: "user not exists, it does not return error",
			staff: staff.Staff{
				UserID:  identity.IssueID(),
				IsAdmin: false,
			},
			wantErr: nil,
		},
	}
	for i := range tests {
		tt := tests[i]
		t.Run(tt.name, func(t *testing.T) {
			w := NewStaff(c)
			err := w.Set(ctx, tt.staff)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
