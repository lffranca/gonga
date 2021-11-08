package presenter

type GatewayString struct {
	ID *string `form:"gateway" binding:"required"`
}
