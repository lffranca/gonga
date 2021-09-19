package kong

type MessageError struct {
	Name    string                 `json:"name"`
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Fields  map[string]interface{} `json:"fields"`
}
