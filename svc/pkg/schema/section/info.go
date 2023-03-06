package section

type SectionInfoResponse struct {
	ID        string     `json:"id"`
	Questions []Question `json:"questions"`
}

type Question struct {
	ID             string          `json:"id"`
	Type           string          `json:"type"`
	Text           string          `json:"text"`
	Options        *[]Option       `json:"options,omitempty"`
	TextConstraint *TextConstraint `json:"text_constraint,omitempty"`
	FileConstraint *FileConstraint `json:"file_constraint,omitempty"`
}

type Option struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type TextConstraint struct {
	MinLength int    `json:"min_length"`
	MaxLength int    `json:"max_length"`
	Regex     string `json:"regex"`
}

type FileConstraint struct {
	Format string `json:"format"`
}
