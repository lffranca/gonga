package tag

const URL = "tags"

type Response struct {
	Offset interface{} `json:"offset"`
	Data []Tag      `json:"data"`
}

type Tag struct {
	EntityID   string `json:"entity_id"`
	Tag        string `json:"tag"`
	EntityName string `json:"entity_name"`
}
