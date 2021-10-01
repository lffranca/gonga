package entity

type Service struct {
	ID                  *string
	Name                *string
	Protocol            *string
	Retries             *int
	Port                *int
	Tags                []*string
	CACertificates      []*string
	Path                *string
	ConnectTimeout      *int
	WriteTimeout        *int
	ReadTimeout         *int
	TLSVerify           *bool
	ClientCertificateID *string
	Host                *string
	TLSVerifyDepth      *int
	CreatedAt           *int
	UpdatedAt           *int
}
