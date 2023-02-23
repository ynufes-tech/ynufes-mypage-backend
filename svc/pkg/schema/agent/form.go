package agent

type CreateFormRequest struct {
	EventID     string `json:"event_id"`
	Title       string `json:"title"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Deadline    int64  `json:"deadline"`
}

type CreateFormResponse struct {
	FormID string `json:"form_id"`
}
