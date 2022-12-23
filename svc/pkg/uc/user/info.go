package uc

import (
	"encoding/json"
	"ynufes-mypage-backend/svc/pkg/domain/model/user"
)

type InfoUseCase struct {
}

func NewInfoUseCase() InfoUseCase {
	return InfoUseCase{}
}

type InfoInput struct {
	User user.User
}

type InfoOutput struct {
	Response string
}

func (uc InfoUseCase) Do(input InfoInput) *InfoOutput {
	resp := struct {
		NameFirst       string `json:"name_first"`
		NameLast        string `json:"name_last"`
		Type            int    `json:"type"`
		ProfileImageURL string `json:"profile_icon_url"`
		Status          int    `json:"status"`
	}{
		NameFirst:       input.User.Detail.Name.FirstName,
		NameLast:        input.User.Detail.Name.LastName,
		Type:            input.User.Detail.Type,
		ProfileImageURL: string(input.User.Line.LineProfilePictureURL),
		Status:          int(input.User.Status),
	}
	j, _ := json.Marshal(resp)
	return &InfoOutput{Response: string(j)}
}
