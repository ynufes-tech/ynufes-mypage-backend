package writer

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/line"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
)

func TestLine_Create(t *testing.T) {
	db := testutil.NewFirebaseTest()
	defer db.Reset()
	w := NewLine(db.GetClient())
	r := reader.NewLine(db.GetClient())
	testLineUsers := []line.LineUser{
		{
			UserID:                identity.IssueID(),
			LineServiceID:         "lineServiceID1",
			LineDisplayName:       "lineDisplayName1",
			EncryptedAccessToken:  "encryptedAccessToken1",
			EncryptedRefreshToken: "encryptedRefreshToken1",
		}, {
			UserID:                identity.IssueID(),
			LineServiceID:         "lineServiceID2",
			LineDisplayName:       "lineDisplayName2",
			EncryptedAccessToken:  "encryptedAccessToken2",
			EncryptedRefreshToken: "encryptedRefreshToken2",
		},
	}
	assert.NoError(t, w.Create(context.Background(), testLineUsers[0]))
	assert.NoError(t, w.Create(context.Background(), testLineUsers[1]))

	tests := []struct {
		name    string
		target  line.LineUser
		wantErr error
	}{
		{
			name: "Success",
			target: line.LineUser{
				UserID:                identity.IssueID(),
				LineServiceID:         "lineServiceID3",
				LineDisplayName:       "lineDisplayName3",
				EncryptedAccessToken:  "encryptedAccessToken3",
				EncryptedRefreshToken: "encryptedRefreshToken3",
			},
			wantErr: nil,
		}, {
			name:    "AlreadyExists",
			target:  testLineUsers[0],
			wantErr: exception.ErrAlreadyExists,
		}, {
			name: "UserIDNotAssigned",
			target: line.LineUser{
				UserID:                nil,
				LineServiceID:         "lineServiceID4",
				LineDisplayName:       "lineDisplayName4",
				EncryptedAccessToken:  "encryptedAccessToken4",
				EncryptedRefreshToken: "encryptedRefreshToken4",
			},
			wantErr: exception.ErrIDNotAssigned,
		}, {
			name: "EmptyLineServiceID",
			target: line.LineUser{
				UserID:                identity.IssueID(),
				LineServiceID:         "",
				LineDisplayName:       "lineDisplayName5",
				EncryptedAccessToken:  "encryptedAccessToken5",
				EncryptedRefreshToken: "encryptedRefreshToken5",
			},
			wantErr: exception.ErrIDNotAssigned,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := w.Create(context.Background(), tt.target)
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				found, err := r.GetByUserID(context.Background(), tt.target.UserID)
				assert.NoError(t, err)
				assert.Equal(t, tt.target, *found)
			}
		})
	}
}

func TestLine_Set(t *testing.T) {
	db := testutil.NewFirebaseTest()
	defer db.Reset()
	w := NewLine(db.GetClient())
	r := reader.NewLine(db.GetClient())
	testLineUsers := []line.LineUser{
		{
			UserID:                identity.IssueID(),
			LineServiceID:         "lineServiceID1",
			LineDisplayName:       "lineDisplayName1",
			EncryptedAccessToken:  "encryptedAccessToken1",
			EncryptedRefreshToken: "encryptedRefreshToken1",
		}, {
			UserID:                identity.IssueID(),
			LineServiceID:         "lineServiceID2",
			LineDisplayName:       "lineDisplayName2",
			EncryptedAccessToken:  "encryptedAccessToken2",
			EncryptedRefreshToken: "encryptedRefreshToken2",
		},
	}
	assert.NoError(t, w.Set(context.Background(), testLineUsers[0]))
	assert.NoError(t, w.Set(context.Background(), testLineUsers[1]))

	tests := []struct {
		name    string
		target  line.LineUser
		wantErr error
	}{
		{
			name: "Success",
			target: line.LineUser{
				UserID:                identity.IssueID(),
				LineServiceID:         "lineServiceID3",
				LineDisplayName:       "lineDisplayName3",
				EncryptedAccessToken:  "encryptedAccessToken3",
				EncryptedRefreshToken: "encryptedRefreshToken3",
			},
			wantErr: nil,
		}, {
			name:    "AlreadyExists",
			target:  testLineUsers[0],
			wantErr: nil,
		}, {
			name: "UserIDNotAssigned",
			target: line.LineUser{
				UserID:                nil,
				LineServiceID:         "lineServiceID4",
				LineDisplayName:       "lineDisplayName4",
				EncryptedAccessToken:  "encryptedAccessToken4",
				EncryptedRefreshToken: "encryptedRefreshToken4",
			},
			wantErr: exception.ErrIDNotAssigned,
		}, {
			name: "EmptyLineServiceID",
			target: line.LineUser{
				UserID:                identity.IssueID(),
				LineServiceID:         "",
				LineDisplayName:       "lineDisplayName5",
				EncryptedAccessToken:  "encryptedAccessToken5",
				EncryptedRefreshToken: "encryptedRefreshToken5",
			},
			wantErr: exception.ErrIDNotAssigned,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := w.Set(context.Background(), tt.target)
			assert.Equal(t, tt.wantErr, err)
			if tt.wantErr == nil {
				found, err := r.GetByUserID(context.Background(), tt.target.UserID)
				assert.NoError(t, err)
				assert.Equal(t, tt.target, *found)
			}
		})
	}
}
