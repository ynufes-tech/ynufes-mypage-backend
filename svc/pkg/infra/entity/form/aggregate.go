package entity

type (
	Form struct {
		ID          int64               `firestore:"id"`
		Title       string              `firestore:"title"`
		Summary     string              `firestore:"summary"`
		Description string              `firestore:"description"`
		Questions   map[string]Question `firestore:"questions"`
	}
	Question struct {
		ID           string                 `firestore:"question_id"`
		QuestionText string                 `firestore:"question"`
		QuestionType int                    `firestore:"question_type"`
		QuestionData map[string]interface{} `firestore:"question_data"`
		Order        int                    `firestore:"order"`
	}
)
