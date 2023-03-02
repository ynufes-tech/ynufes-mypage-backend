package entity

type (
	// UserDetail StudentID 学部生は数字7桁, 大学院生は数字2桁+英数字2桁+数字3桁
	// all fields have omitempty tag
	// in order to easily handle incomplete request from frontend
	UserDetail struct {
		NameFirst     string `json:"detail-name_first,omitempty"`
		NameFirstKana string `json:"detail-name_first_kana,omitempty"`
		NameLast      string `json:"detail-name_last,omitempty"`
		NameLastKana  string `json:"detail-name_last_kana,omitempty"`
		Gender        int    `json:"detail-gender,omitempty"`
		StudentID     string `json:"detail-student_id,omitempty"`
		Email         string `json:"detail-email,omitempty"`
		Type          int    `json:"detail-type,omitempty"`
	}
)
