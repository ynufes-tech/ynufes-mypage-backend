package entity

type (
	// UserDetail StudentID 学部生は数字7桁, 大学院生は数字2桁+英数字2桁+数字3桁
	UserDetail struct {
		NameFirst     string `json:"detail-name_first"`
		NameFirstKana string `json:"detail-name_first_kana"`
		NameLast      string `json:"detail-name_last"`
		NameLastKana  string `json:"detail-name_last_kana"`
		Gender        int    `json:"detail-gender"`
		StudentID     string `json:"detail-student_id"`
		Email         string `json:"detail-email"`
		Type          int    `json:"detail-type"`
	}
)
