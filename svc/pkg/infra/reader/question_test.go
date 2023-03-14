package reader

import (
	"context"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/exception"
)

func TestQuestion_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		query   id.QuestionID
		want    *question.Question
		wantErr error
	}{
		{
			name:    "NotFound",
			query:   identity.IssueID(),
			want:    nil,
			wantErr: exception.ErrNotFound,
		},
	}
	fbt := testutil.NewFirebaseTest()
	defer fbt.Reset()
	r := NewQuestion(fbt.GetClient())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := r.GetByID(ctx, tt.query)
			if err != tt.wantErr {
				t.Errorf("GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
