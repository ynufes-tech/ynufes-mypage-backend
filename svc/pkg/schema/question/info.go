package question

type (
	CheckboxQuestionInfo struct {
		Options []CheckboxOptionInfo `json:"options"`
	}

	CheckboxOptionInfo struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}

	RadioQuestionInfo struct {
		Options []RadioOptionInfo `json:"options"`
	}

	RadioOptionInfo struct {
		ID   string `json:"id"`
		Text string `json:"text"`
	}
)
