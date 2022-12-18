package infra

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

func genTestCase() []user.User {
	return []user.User{
		{
			ID:     1234,
			Status: user.StatusNew,
			Detail: user.Detail{
				Name: user.Name{
					FirstName:     "詩恩",
					LastName:      "市川",
					FirstNameKana: "シオン",
					LastNameKana:  "イチカワ",
				},
				Email:     "shion1305@gmail.com",
				Gender:    user.GenderMan,
				StudentID: "2164027",
				Type:      user.TypeNormal,
			},
			Line: user.Line{
				LineServiceID:         "LineServiceID",
				EncryptedAccessToken:  "EncryptedAccessToken",
				EncryptedRefreshToken: "EncryptedRefreshToken",
			},
			Dashboard: user.Dashboard{
				Grants: []string{"grant1", "grant2"},
			},
		},
		{
			ID:     12344,
			Status: user.StatusRegistered,
			Detail: user.Detail{
				Name: user.Name{
					FirstName:     "友哉",
					LastName:      "廣江",
					FirstNameKana: "トモヤ",
					LastNameKana:  "ヒロエ",
				},
				Email:     "tomoya4creative@gmail.com",
				Gender:    user.GenderMan,
				StudentID: "2125178",
				Type:      user.TypeNormal,
			},
			Line: user.Line{
				LineServiceID:         "LineServiceID",
				EncryptedAccessToken:  "EncryptedAccessToken",
				EncryptedRefreshToken: "EncryptedRefreshToken",
			},
			Dashboard: user.Dashboard{
				Grants: []string{"grant1", "grant2"},
			},
		},
	}
}
