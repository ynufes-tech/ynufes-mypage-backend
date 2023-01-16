package entity

type (
	Form struct {
		ID          string              `firestore:"-"`
		Title       string              `firestore:"title"`
		Summary     string              `firestore:"summary"`
		Description string              `firestore:"description"`
		Questions   map[string]Question `firestore:"questions"`
	}
	Question struct {
		ID           string      `firestore:"question_id"`
		QuestionText string      `firestore:"question"`
		QuestionType int         `firestore:"question_type"`
		Properties   interface{} `firestore:"properties"`
		Order        int         `firestore:"order"`
	}
)
