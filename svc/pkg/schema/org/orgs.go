package org

type (
	OrgsResponse struct {
		Orgs []Org `json:"orgs"`
	}
	Org struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		EventName string `json:"event_name"`
		EventID   string `json:"event_id"`
		IsOpen    bool   `json:"is_open"`
	}
)
