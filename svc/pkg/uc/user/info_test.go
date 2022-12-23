package uc

import (
	"github.com/go-playground/assert/v2"
	"testing"
	userDomain "ynufes-mypage-backend/svc/pkg/domain/model/user"
)

func TestInfoUseCase(t *testing.T) {
	tests := []struct {
		args InfoInput
		want InfoOutput
	}{
		{
			args: InfoInput{
				User: userDomain.User{
					ID:     12345,
					Status: 0,
					Detail: userDomain.Detail{
						Name: userDomain.Name{
							FirstName:     "名前",
							LastName:      "苗字",
							FirstNameKana: "ナマエ",
							LastNameKana:  "ミョウジ",
						},
						Email:     "testing@testing.co.jp",
						Gender:    userDomain.GenderWoman,
						StudentID: "2164022",
						Type:      userDomain.TypeNormal,
					},
					Line: userDomain.Line{
						LineServiceID:         "SERVICE_ID",
						LineProfilePictureURL: "https://testing.co.jp/test.png",
						LineDisplayName:       "みょーじねーむ",
						EncryptedAccessToken:  "TestAccessToken",
						EncryptedRefreshToken: "TestRefreshToken",
					},
				},
			},
			want: InfoOutput{
				Response: `{"name_first":"名前","name_last":"苗字","type":0,"profile_icon_url":"https://testing.co.jp/test.png","status":0}`,
			},
		},
	}
	uc := NewInfoUseCase()
	for _, t := range tests {
		out := uc.Do(t.args)
		assert.IsEqual(out, t.want)
	}
}
