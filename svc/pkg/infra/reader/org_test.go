package reader

import (
	"context"
	"errors"
	"testing"
	"ynufes-mypage-backend/pkg/firebase"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/org"
	"ynufes-mypage-backend/svc/pkg/exception"
)

func TestOrg_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		query   id.OrgID
		want    *org.Org
		wantErr error
	}{
		{
			name:    "NotFound",
			want:    nil,
			query:   identity.IssueID(),
			wantErr: exception.ErrNotFound,
		},
	}
	fb := firebase.New()
	r := NewOrg(&fb)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := r.GetByID(ctx, tt.query)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
