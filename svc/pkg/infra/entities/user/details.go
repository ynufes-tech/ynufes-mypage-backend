package entities

type (
	// UserDetail StudentID 学部生は数字7桁, 大学院生は数字2桁+英数字2桁+数字3桁
	UserDetail struct {
		NameFirst     string `firestore:"name_first"`
		NameFirstKana string `firestore:"name_first_kana"`
		NameLast      string `firestore:"name_last"`
		NameLastKana  string `firestore:"name_last_kana"`
		Gender        int    `firestore:"gender"`
		StudentID     string `firestore:"student_id"`
	}
)
