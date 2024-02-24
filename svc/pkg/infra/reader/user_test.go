package reader

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/pkg/testutil"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
	"ynufes-mypage-backend/svc/pkg/exception"
	"ynufes-mypage-backend/svc/pkg/infra/writer"
)

func TestUser_GetByID(t *testing.T) {
	users := []user.User{
		{
			Detail: user.Detail{
				Name: user.Name{
					FirstName:     "詩恩",
					LastName:      "市川",
					FirstNameKana: "シオン",
					LastNameKana:  "イチカワ",
				},
				Email:      "shion1305@gmail.com",
				Gender:     user.GenderMan,
				StudentID:  "2164027",
				Type:       user.TypeNormal,
				PictureURL: "https://shion1305.com/picture.png",
			},
			Agent: user.Agent{
				Roles: []user.Role{},
			},
		},
	}

	ctx := context.Background()
	fbt := testutil.NewFirebaseTest()
	defer fbt.Reset()
	w := writer.NewUser(fbt.GetClient())
	for i := range users {
		assert.NoError(t, w.Create(ctx, &users[i]))
	}

	tests := []struct {
		name    string
		query   id.UserID
		want    *user.User
		wantErr error
	}{
		{
			name:    "Success",
			query:   users[0].ID,
			want:    &users[0],
			wantErr: nil,
		},
		{
			name:    "NotFound",
			query:   identity.IssueID(),
			want:    nil,
			wantErr: exception.ErrNotFound,
		},
	}
	r := NewUser(fbt.GetClient())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			got, err := r.GetByID(ctx, tt.query)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
