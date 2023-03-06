package form

type FormInfoResponse struct {
	ID          string   `json:"id"`
	Title       string   `json:"name"`
	Summary     string   `json:"summary"`
	Description string   `json:"description"`
	Deadline    string   `json:"deadline"`
	Status      int      `json:"status"`
	IsOpen      bool     `json:"is_open"`
	Sections    []string `json:"sections"`
}
