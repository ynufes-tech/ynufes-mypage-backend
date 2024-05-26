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
	FileType        FileTypes        `json:"file_types"`
	Extensions      []string         `json:"extensions"`
	ImageConstraint *ImageConstraint `json:"img_constraint,omitempty"`
}

type FileTypes struct {
	AcceptAny   bool `json:"any"`
	AcceptImage bool `json:"image"`
	AcceptPDF   bool `json:"pdf"`
}

type ImageConstraint struct {
	Min    *int          `json:"min"`
	Max    *int          `json:"max"`
	Width  DimensionSpec `json:"width"`
	Height DimensionSpec `json:"height"`
	Ratio  RatioSpec     `json:"ratio"`
}

type DimensionSpec struct {
	Min *int `json:"min"`
	Max *int `json:"max"`
	Eq  *int `json:"eq"`
}

type RatioSpec struct {
	Min *float32 `json:"min"`
	Max *float32 `json:"max"`
	Eq  *float32 `json:"eq"`
}
