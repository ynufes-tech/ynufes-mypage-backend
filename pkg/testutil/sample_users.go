package testutil

import (
	"ynufes-mypage-backend/pkg/identity"
	"ynufes-mypage-backend/svc/pkg/domain/model/id"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

func Users() []user.User {
	return []user.User{
		{
			ID: id.UserID(identity.NewID(1234)),
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
		},
		{
			ID: id.UserID(identity.NewID(12344)),
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
		},
	}
}
