package presenter

type GatewayURI struct {
	ID *string `json:"id,omitempty" yaml:"id,omitempty" uri:"gateway" binding:"required"`
}
