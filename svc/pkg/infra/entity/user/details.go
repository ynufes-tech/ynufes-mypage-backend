package entity

type (
	// UserDetail StudentID 学部生は数字7桁, 大学院生は数字2桁+英数字2桁+数字3桁
	// all fields have omitempty tag
	// in order to easily handle incomplete request from frontend
	UserDetail struct {
		NameFirst     string `json:"name_first,omitempty"`
		NameFirstKana string `json:"name_first_kana,omitempty"`
		NameLast      string `json:"name_last,omitempty"`
		NameLastKana  string `json:"name_last_kana,omitempty"`
		Gender        int    `json:"gender,omitempty"`
		StudentID     string `json:"student_id,omitempty"`
		Email         string `json:"email,omitempty"`
		Type          int    `json:"type,omitempty"`
		PictureURL    string `json:"picture_url,omitempty"`
	}
)
