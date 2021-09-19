package sni

const URL = "snis"

type Response struct {
	Offset *string `json:"offset"`
	Data   []SNI   `json:"data"`
}

type ItemID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Form struct {
	Name        *string  `json:"name,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Certificate *ItemID  `json:"certificate,omitempty"`
}

type SNI struct {
	ID          *string  `json:"id,omitempty"`
	CreatedAt   *int     `json:"created_at,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Certificate *ItemID  `json:"certificate,omitempty"`
}
