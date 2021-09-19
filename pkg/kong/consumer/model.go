package consumer

const URL = "consumers"

type Form struct {
	Username *string  `json:"username,omitempty"`
	CustomID *string  `json:"custom_id,omitempty"`
	Tags     []string `json:"tags,omitempty"`
}

type Response struct {
	Offset *string    `json:"offset"`
	Data   []Consumer `json:"data"`
}

type Consumer struct {
	Tags      []string `json:"tags"`
	CustomID  *string  `json:"custom_id"`
	Username  *string  `json:"username"`
	ID        *string  `json:"id"`
	CreatedAt *int     `json:"created_at"`
}
