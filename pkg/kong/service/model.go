package service

const URL = "services"

type Response struct {
	Next interface{} `json:"next"`
	Data []Service   `json:"data"`
}

type Form struct {
	Host              *string             `json:"host,omitempty"`
	ConnectTimeout    *int                `json:"connect_timeout,omitempty"`
	Protocol          *string             `json:"protocol,omitempty"`
	Name              *string             `json:"name,omitempty"`
	ReadTimeout       *int                `json:"read_timeout,omitempty"`
	Port              *int                `json:"port,omitempty"`
	Path              *string             `json:"path,omitempty"`
	Retries           *int                `json:"retries,omitempty"`
	WriteTimeout      *int                `json:"write_timeout,omitempty"`
	Tags              []string           `json:"tags,omitempty"`
	ClientCertificate *ClientCertificate `json:"client_certificate,omitempty"`
}

type ClientCertificate struct {
	ID string `json:"id"`
}

type Service struct {
	Name              string             `json:"name"`
	TLSVerifyDepth    interface{}        `json:"tls_verify_depth"`
	Protocol          string             `json:"protocol"`
	Retries           int                `json:"retries"`
	Port              int                `json:"port"`
	Tags              []string           `json:"tags"`
	CaCertificates    interface{}        `json:"ca_certificates"`
	ID                string             `json:"id"`
	Path              string             `json:"path"`
	ConnectTimeout    int                `json:"connect_timeout"`
	WriteTimeout      int                `json:"write_timeout"`
	ReadTimeout       int                `json:"read_timeout"`
	TLSVerify         interface{}        `json:"tls_verify"`
	ClientCertificate *ClientCertificate `json:"client_certificate"`
	Host              string             `json:"host"`
	CreatedAt         int                `json:"created_at"`
	UpdatedAt         int                `json:"updated_at"`
}
