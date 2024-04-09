package linepkg

// WebHook types
type EventMessage struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

type WebhookEvent struct {
	Type    string       `json:"type"`
	Message EventMessage `json:"message"`
}

type WebhookBody struct {
	Destination string         `json:"destination"`
	Events      []WebhookEvent `json:"events"`
}
