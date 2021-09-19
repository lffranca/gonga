package plugin

const URL = "plugins"

type Response struct {
	Offset *string  `json:"offset,omitempty"`
	Data   []Plugin `json:"data,omitempty"`
}

type ItemID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Form struct {
	Name      *string     `json:"name,omitempty"`
	Route     *ItemID     `json:"route,omitempty"`
	Service   *ItemID     `json:"service,omitempty"`
	Consumer  *ItemID     `json:"consumer,omitempty"`
	Config    interface{} `json:"config,omitempty"`
	Protocols []string    `json:"protocols,omitempty"`
	Enabled   *bool       `json:"enabled,omitempty"`
	Tags      []string    `json:"tags,omitempty"`
}

type Plugin struct {
	ID        *string     `json:"id,omitempty"`
	Name      *string     `json:"name,omitempty"`
	Route     *ItemID     `json:"route,omitempty"`
	Service   *ItemID     `json:"service,omitempty"`
	Consumer  *ItemID     `json:"consumer,omitempty"`
	Config    interface{} `json:"config,omitempty"`
	Protocols []string    `json:"protocols,omitempty"`
	Enabled   *bool       `json:"enabled,omitempty"`
	Tags      []string    `json:"tags,omitempty"`
	CreatedAt *int        `json:"created_at,omitempty"`
}
