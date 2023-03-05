package user

type InfoResponse struct {
	NameFirst       string `json:"name_first"`
	NameLast        string `json:"name_last"`
	Type            int    `json:"type"`
	ProfileImageURL string `json:"profile_icon_url"`
}
