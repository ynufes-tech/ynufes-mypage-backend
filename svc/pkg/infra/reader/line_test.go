package reader

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/line"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

func TestLine(t *testing.T) {
	fb := testutil.NewFirebaseTest()
	defer fb.Reset()
	lineUsers := []line.LineUser{
		{
			UserID:                identity.IssueID(),
			LineServiceID:         "LineServiceID1",
			LineDisplayName:       "LineDisplayName1",
			EncryptedAccessToken:  "EncryptedAccessToken1",
			EncryptedRefreshToken: "EncryptedRefreshToken1",
		}, {
			UserID:                identity.IssueID(),
			LineServiceID:         "LineServiceID2",
			LineDisplayName:       "LineDisplayName2",
			EncryptedAccessToken:  "EncryptedAccessToken2",
			EncryptedRefreshToken: "EncryptedRefreshToken2",
		},
	}
	w := writer.NewLine(fb.GetClient())
	for i := range lineUsers {
		assert.NoError(t, w.Create(context.Background(), lineUsers[i]))
	}
	r := NewLine(fb.GetClient())
	testsByUsers := []struct {
		name    string
		userID  id.UserID
		want    *line.LineUser
		wantErr error
	}{
		{
			name:    "Success1",
			userID:  lineUsers[0].UserID,
			want:    &lineUsers[0],
			wantErr: nil,
		}, {
			name:    "Success2",
			userID:  lineUsers[1].UserID,
			want:    &lineUsers[1],
			wantErr: nil,
		}, {
			name:    "NotFound",
			userID:  identity.IssueID(),
			want:    nil,
			wantErr: exception.ErrNotFound,
		}, {
			name:    "InvalidID",
			userID:  nil,
			want:    nil,
			wantErr: exception.ErrIDNotAssigned,
		},
	}
	for _, tt := range testsByUsers {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetByUserID(context.Background(), tt.userID)
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
	testsLineServiceID := []struct {
		name    string
		lineID  line.LineServiceID
		want    *line.LineUser
		wantErr error
	}{
		{
			name:    "Success1",
			lineID:  lineUsers[0].LineServiceID,
			want:    &lineUsers[0],
			wantErr: nil,
		}, {
			name:    "Success2",
			lineID:  lineUsers[1].LineServiceID,
			want:    &lineUsers[1],
			wantErr: nil,
		}, {
			name:    "NotFound",
			lineID:  "IDNotExists",
			want:    nil,
			wantErr: exception.ErrNotFound,
		}, {
			name:    "InvalidID",
			lineID:  "",
			want:    nil,
			wantErr: exception.ErrIDNotAssigned,
		},
	}
	for _, tt := range testsLineServiceID {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetByLineServiceID(context.Background(), tt.lineID)
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
