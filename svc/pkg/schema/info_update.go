package schema

import "ynufes-mypage-backend/svc/pkg/domain/model/user"

type InfoUpdateRequest struct {
	NameFirst     string `json:"name_first"`
	NameLast      string `json:"name_last"`
	NameFirstKana string `json:"name_first_kana"`
	NameLastKana  string `json:"name_last_kana"`
	Email         string `json:"email"`
	Gender        int    `json:"gender"`
	StudentID     string `json:"student_id"`
}

func (r InfoUpdateRequest) ToUserDetail() (*user.Detail, error) {
	email, err := user.NewEmail(r.Email)
	if err != nil {
		return nil, err
	}
	gender, err := user.NewGender(r.Gender)
	if err != nil {
		return nil, err
	}
	return &user.Detail{
		Name: user.Name{
			FirstName:     r.NameFirst,
			LastName:      r.NameLast,
			FirstNameKana: r.NameFirstKana,
			LastNameKana:  r.NameLastKana,
		},
		Email:     email,
		Gender:    gender,
		StudentID: user.StudentID(r.StudentID),
	}, nil
}
