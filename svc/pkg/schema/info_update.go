package schema

import (
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type InfoUpdateRequest struct {
	NameFirst     string `json:"name_first"`
	NameLast      string `json:"name_last"`
	NameFirstKana string `json:"name_first_kana"`
	NameLastKana  string `json:"name_last_kana"`
	Email         string `json:"email"`
	Gender        int    `json:"gender"`
	StudentID     string `json:"student_id"`
}

func (r InfoUpdateRequest) ApplyToDetail(d *user.Detail) error {
	// get each field from InfoUpdateRequest, if each field has a value,
	// replace a field in user.Detail with matching json value
	cp := *d
	if r.Email != "" {
		email, err := user.NewEmail(r.Email)
		if err != nil {
			return err
		}
		cp.Email = email
	}
	if r.Gender != 0 {
		gender, err := user.NewGender(r.Gender)
		if err != nil {
			return err
		}
		cp.Gender = gender
	}
	if r.StudentID != "" {
		cp.StudentID = user.StudentID(r.StudentID)
	}
	if r.NameFirst != "" {
		cp.Name.FirstName = r.NameFirst
	}
	if r.NameLast != "" {
		cp.Name.LastName = r.NameLast
	}
	if r.NameFirstKana != "" {
		cp.Name.FirstNameKana = r.NameFirstKana
	}
	if r.NameLastKana != "" {
		cp.Name.LastNameKana = r.NameLastKana
	}
	*d = cp
	return nil
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
