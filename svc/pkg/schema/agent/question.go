package agent

import (
	schema "ynufes-mypage-backend/svc/pkg/schema/question"
)

type (
	CreateQuestionRequest struct {
		FormID    string           `json:"form_id"`
		SectionID string           `json:"section_id"`
		Type      string           `json:"type"`
		Text      string           `json:"text"`
		AfterID   string           `json:"qid_after"`
		PosAt     *int             `json:"position_at"`
		Checkbox  *NewCheckboxInfo `json:"checkbox"`
		Radio     *NewRadioInfo    `json:"radio"`
	}

	NewCheckboxInfo struct {
		Options []string `json:"options"`
	}

	NewRadioInfo struct {
		Options []string `json:"options"`
	}

	CreateQuestionResponse struct {
		QuestionID string                       `json:"question_id"`
		Radio      *schema.RadioQuestionInfo    `json:"radio,omitempty"`
		Checkbox   *schema.CheckboxQuestionInfo `json:"checkbox,omitempty"`
	}
)
