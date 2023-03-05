package writer

import (
	"context"
	"testing"
	"time"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
)

func TestRelation_CreateOrgUser(t *testing.T) {
	fb := firebase.New()
	w := NewRelation(&fb)
	tests := []struct {
		name string
		give struct {
			OrgID  id.OrgID
			UserID id.UserID
		}
		hasError bool
	}{
		{
			name: "normal create",
			give: struct {
				OrgID  id.OrgID
				UserID id.UserID
			}{
				OrgID:  id.OrgID(identity.IssueID()),
				UserID: id.UserID(identity.IssueID()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if err := w.CreateOrgUser(ctx, tt.give.OrgID, tt.give.UserID); (err != nil) != tt.hasError {
				t.Errorf("CreateOrgUser() error = %v, hasError %v", err, tt.hasError)
			}
			t.Logf("user: %v, org: %v", tt.give.UserID.GetValue(), tt.give.OrgID.GetValue())
		})
	}
}

func TestRelation_DeleteOrgUser(t *testing.T) {
	fb := firebase.New()
	w := NewRelation(&fb)
	tests := []struct {
		name string
		give struct {
			OrgID  id.OrgID
			UserID id.UserID
		}
		hasError bool
	}{
		{
			name: "normal create",
			give: struct {
				OrgID  id.OrgID
				UserID id.UserID
			}{
				OrgID:  id.OrgID(identity.IssueID()),
				UserID: id.UserID(identity.IssueID()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			if err := w.CreateOrgUser(ctx, tt.give.OrgID, tt.give.UserID); (err != nil) != tt.hasError {
				t.Errorf("CreateOrgUser() error = %v, hasError %v", err, tt.hasError)
			}
			time.Sleep(5 * time.Second)
			t.Logf("user: %v, org: %v", tt.give.UserID.GetValue(), tt.give.OrgID.GetValue())
			err := w.DeleteOrgUser(ctx, tt.give.OrgID, tt.give.UserID)
			if err != nil {
				t.Error(err)
				return
			}
		})
	}
}
