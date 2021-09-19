package cacertificate

const URL = "ca_certificates"

type Response struct {
	Offset *string         `json:"offset,omitempty"`
	Data   []CACertificate `json:"data,omitempty"`
}

type Form struct {
	Tags       []string `json:"tags,omitempty"`
	CertDigest *string  `json:"cert_digest,omitempty"`
	Cert       *string  `json:"cert,omitempty"`
}

type CACertificate struct {
	Tags       []string `json:"tags,omitempty"`
	CertDigest *string  `json:"cert_digest,omitempty"`
	Cert       *string  `json:"cert,omitempty"`
	CreatedAt  *int     `json:"created_at,omitempty"`
	ID         *string  `json:"id,omitempty"`
}
