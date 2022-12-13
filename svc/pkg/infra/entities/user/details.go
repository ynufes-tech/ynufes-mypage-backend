package entities

type (
	UserDetail struct {
		NameFirst     string `firestore:"name_first"`
		NameFirstKana string `firestore:"name_first_kana"`
		NameLast      string `firestore:"name_last"`
		NameLastKana  string `firestore:"name_last_kana"`
		Gender        int    `firestore:"gender"`
		AdmissionYear int    `firestore:"admission_year"`
	}
)
