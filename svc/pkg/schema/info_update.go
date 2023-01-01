package schema

type InfoUpdateRequest struct {
	NameFirst     string `json:"name_first"`
	NameLast      string `json:"name_last"`
	NameFirstKana string `json:"name_first_kana"`
	NameLastKana  string `json:"name_last_kana"`
	Email         string `json:"email"`
	Gender        int    `json:"gender"`
	StudentID     string `json:"student_id"`
}
