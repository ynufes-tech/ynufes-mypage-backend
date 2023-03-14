package reader

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/event"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/question"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

func TestQuestion_GetByID(t *testing.T) {
	fbt := testutil.NewFirebaseTest()
	defer fbt.Reset()
	event1 := event.Event{
		Name: "TestEvent1",
	}
	eventW := writer.NewEvent(fbt.GetClient())
	assert.NoError(t, eventW.Create(context.Background(), &event1))
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
