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
)

func TestRelationRoleStaff_Create(t *testing.T) {
	ctx := context.Background()
	fbt := testutil.NewFirebaseTest()
	defer fbt.Reset()
	roleID := id.RoleID(identity.IssueID())
	userID := id.UserID(identity.IssueID())
	fmt.Println("firebase test created")
	w := NewRelationRoleStaff(fbt.GetClient())
	err := w.CreateRoleStaff(ctx, roleID, userID)
	assert.NoError(t, err)

	// TODO: create reader and check if the relation is created
	//relationR := reader.NewRelationRoleStaff(fbt.GetClient())
	//orgs, err := relationR.ListOrgIDsByUserID(ctx, userID)
	//assert.NoError(t, err)
	//assert.Equal(t, []id.RoleID{roleID}, orgs)
	//users, err := relationR.ListUserIDsByOrgID(ctx, roleID)
	//assert.NoError(t, err)
	//assert.Equal(t, []id.UserID{userID}, users)
}

func TestRelationRoleStaff_Delete(t *testing.T) {
	type relationRoleStaff struct {
		UserID id.UserID
		RoleID id.RoleID
	}

	rSimple := relationRoleStaff{RoleID: identity.IssueID(), UserID: id.UserID(identity.IssueID())}
	rMultiple := relationRoleStaff{RoleID: identity.IssueID(), UserID: id.UserID(identity.IssueID())}
	relations := []relationRoleStaff{rSimple, rMultiple, rMultiple, rMultiple}
	tests := []struct {
		name    string
		give    relationRoleStaff
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
			give:    relationRoleStaff{RoleID: identity.IssueID(), UserID: id.UserID(identity.IssueID())},
			wantErr: exception.ErrNotFound,
		},
		{
			name:    "not exist - 2",
			give:    relationRoleStaff{RoleID: relations[0].RoleID, UserID: id.UserID(identity.IssueID())},
			wantErr: exception.ErrNotFound,
		},
		{
			name:    "not exist - 3",
			give:    relationRoleStaff{RoleID: identity.IssueID(), UserID: relations[0].UserID},
			wantErr: exception.ErrNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fb := testutil.NewFirebaseTest()
			defer fb.Reset()

			ctx := context.Background()
			w := NewRelationRoleStaff(fb.GetClient())

			for _, r := range relations {
				assert.NoError(t, w.CreateRoleStaff(ctx, r.RoleID, r.UserID))
			}

			err := w.DeleteRoleStaff(ctx, tt.give.RoleID, tt.give.UserID)
			assert.ErrorIs(t, err, tt.wantErr)

			// TODO: create reader and check if the relation is deleted
			//if tt.wantErr == nil {
			//	roles, err := reader.NewRelationRoleStaff(fb.GetClient()).ListRoleIDsByStaffID(ctx, tt.give.UserID)
			//	assert.NoError(t, err)
			//	assert.NotContains(t, roles, tt.give.RoleID)
			//
			//	staffs, err := reader.NewRelationRoleStaff(fb.GetClient()).ListStaffIDsByRoleID(ctx, tt.give.RoleID)
			//	assert.NoError(t, err)
			//	assert.NotContains(t, staffs, tt.give.UserID)
			//}
		})
	}
}
