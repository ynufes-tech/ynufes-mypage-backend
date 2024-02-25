package writer

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/reader"
)

func TestRelationOrgUser_CreateOrgUser(t *testing.T) {
	ctx := context.Background()
	fbt := testutil.NewFirebaseTest()
	defer fbt.Reset()
	orgID := id.OrgID(identity.IssueID())
	userID := id.UserID(identity.IssueID())
	fmt.Println("firebase test created")
	w := NewRelationOrgUser(fbt.GetClient())
	err := w.CreateOrgUser(ctx, orgID, userID)
	assert.NoError(t, err)

	relationR := reader.NewRelationOrgUser(fbt.GetClient())
	orgs, err := relationR.ListOrgIDsByUserID(ctx, userID)
	assert.NoError(t, err)
	assert.Equal(t, id.OrgIDs{orgID}, orgs)
	users, err := relationR.ListUserIDsByOrgID(ctx, orgID)
	assert.NoError(t, err)
	assert.Equal(t, []id.UserID{userID}, users)
}

func TestRelationOrgUser_DeleteOrgUser(t *testing.T) {
	type relationOrgUser struct {
		OrgID  id.OrgID
		UserID id.UserID
	}

	rSimple := relationOrgUser{OrgID: id.OrgID(identity.IssueID()), UserID: id.UserID(identity.IssueID())}
	rMultiple := relationOrgUser{OrgID: id.OrgID(identity.IssueID()), UserID: id.UserID(identity.IssueID())}
	relations := []relationOrgUser{rSimple, rMultiple, rMultiple, rMultiple}
	tests := []struct {
		name    string
		give    relationOrgUser
		wantErr error
	}{
		{
			name:    "normal delete",
			give:    rSimple,
			wantErr: nil,
		},
		{
			name:    "delete with multiple connections",
			give:    rMultiple,
			wantErr: nil,
		},
		{
			name:    "not exist - 1",
			give:    relationOrgUser{OrgID: id.OrgID(identity.IssueID()), UserID: id.UserID(identity.IssueID())},
			wantErr: exception.ErrNotFound,
		},
		{
			name:    "not exist - 2",
			give:    relationOrgUser{OrgID: relations[0].OrgID, UserID: id.UserID(identity.IssueID())},
			wantErr: exception.ErrNotFound,
		},
		{
			name:    "not exist - 3",
			give:    relationOrgUser{OrgID: id.OrgID(identity.IssueID()), UserID: relations[0].UserID},
			wantErr: exception.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := testutil.NewFirebaseTest()
			defer fb.Reset()

			ctx := context.Background()
			w := NewRelationOrgUser(fb.GetClient())

			for _, r := range relations {
				assert.NoError(t, w.CreateOrgUser(ctx, r.OrgID, r.UserID))
			}

			err := w.DeleteOrgUser(ctx, tt.give.OrgID, tt.give.UserID)
			assert.ErrorIs(t, err, tt.wantErr)

			if tt.wantErr == nil {
				orgs, err := reader.NewRelationOrgUser(fb.GetClient()).ListOrgIDsByUserID(ctx, tt.give.UserID)
				assert.NoError(t, err)
				assert.NotContains(t, orgs, tt.give.OrgID)

				users, err := reader.NewRelationOrgUser(fb.GetClient()).ListUserIDsByOrgID(ctx, tt.give.OrgID)
				assert.NoError(t, err)
				assert.NotContains(t, users, tt.give.UserID)
			}
		})
	}
}
