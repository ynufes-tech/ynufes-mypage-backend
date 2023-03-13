package section

type CreateRequest struct {
	FormID string `json:"form_id"`
}

type CreateResponse struct {
	SectionID string `json:"section_id"`
}
