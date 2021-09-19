package certificate

const URL = "certificates"

type Response struct {
	Next interface{}   `json:"next"`
	Data []Certificate `json:"data"`
}

type ItemID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Form struct {
	KeyAlt         *string  `json:"key_alt,omitempty"`
	CaCertificates *ItemID  `json:"ca_certificates,omitempty"`
	TLSVerifyDepth *int     `json:"tls_verify_depth,omitempty"`
	TLSVerify      *bool    `json:"tls_verify,omitempty"`
	Key            *string  `json:"key,omitempty"`
	Cert           *string  `json:"cert,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	SNIs           []string `json:"snis,omitempty"`
	CertAlt        *string  `json:"cert_alt,omitempty"`
}

type Certificate struct {
	KeyAlt         *string  `json:"key_alt,omitempty"`
	CaCertificates *ItemID  `json:"ca_certificates,omitempty"`
	TLSVerifyDepth *int     `json:"tls_verify_depth,omitempty"`
	ID             *string  `json:"id,omitempty"`
	TLSVerify      *bool    `json:"tls_verify,omitempty"`
	Key            *string  `json:"key,omitempty"`
	Cert           *string  `json:"cert,omitempty"`
	Tags           []string `json:"tags,omitempty"`
	SNIs           []string `json:"snis,omitempty"`
	CreatedAt      *int     `json:"created_at,omitempty"`
	CertAlt        *string  `json:"cert_alt,omitempty"`
}
