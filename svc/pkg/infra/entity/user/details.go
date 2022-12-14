package entity

type (
	// UserDetail StudentID 学部生は数字7桁, 大学院生は数字2桁+英数字2桁+数字3桁
	UserDetail struct {
		NameFirst     string `firestore:"detail-name_first"`
		NameFirstKana string `firestore:"detail-name_first_kana"`
		NameLast      string `firestore:"detail-name_last"`
		NameLastKana  string `firestore:"detail-name_last_kana"`
		Gender        int    `firestore:"detail-gender"`
		StudentID     string `firestore:"detail-student_id"`
		Email         string `firestore:"detail-email"`
		Type          int    `firestore:"detail-type"`
	}
)
