package agent

type CreateEventRequest struct {
	EventName string `json:"event_name"`
}

type CreateEventResponse struct {
	EventID   string `json:"event_id"`
	EventName string `json:"event_name"`
}
