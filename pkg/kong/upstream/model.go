package upstream

const URL = "upstreams"

type Response struct {
	Offset *string `json:"offset,omitempty"`
	Data   []Upstream  `json:"data,omitempty"`
}

type Healthy struct {
	Successes    *int  `json:"successes,omitempty"`
	HTTPStatuses []int `json:"http_statuses,omitempty"`
	Interval     *int  `json:"interval,omitempty"`
}

type Unhealthy struct {
	Timeouts     *int  `json:"timeouts,omitempty"`
	Interval     *int  `json:"interval,omitempty"`
	HTTPFailures *int  `json:"http_failures,omitempty"`
	HTTPStatuses []int `json:"http_statuses,omitempty"`
	TCPFailures  *int  `json:"tcp_failures,omitempty"`
}

type Active struct {
	HTTPSVerifyCertificate *bool      `json:"https_verify_certificate,omitempty"`
	Timeout                *int       `json:"timeout,omitempty"`
	Type                   *string    `json:"type,omitempty"`
	HTTPSSni               *string    `json:"https_sni,omitempty"`
	HTTPPath               *string    `json:"http_path,omitempty"`
	Healthy                *Healthy   `json:"healthy,omitempty"`
	Unhealthy              *Unhealthy `json:"unhealthy,omitempty"`
	Concurrency            *int       `json:"concurrency,omitempty"`
}

type Passive struct {
	Healthy   *Healthy   `json:"healthy,omitempty"`
	Type      *string    `json:"type,omitempty"`
	Unhealthy *Unhealthy `json:"unhealthy,omitempty"`
}

type HealthChecks struct {
	Active    *Active  `json:"active,omitempty"`
	Passive   *Passive `json:"passive,omitempty"`
	Threshold *int     `json:"threshold,omitempty"`
}

type ItemID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Form struct {
	Name              *string       `json:"name,omitempty"`
	HostHeader        *string       `json:"host_header,omitempty"`
	HashOnCookiePath  *string       `json:"hash_on_cookie_path,omitempty"`
	HashFallback      *string       `json:"hash_fallback,omitempty"`
	Algorithm         *string       `json:"algorithm,omitempty"`
	Slots             *int          `json:"slots,omitempty"`
	HashOn            *string       `json:"hash_on,omitempty"`
	HealthChecks      *HealthChecks `json:"healthchecks,omitempty"`
	ClientCertificate *ItemID       `json:"client_certificate,omitempty"`
	Tags              []string      `json:"tags,omitempty"`
}

type Upstream struct {
	Name               *string       `json:"name,omitempty"`
	HostHeader         *string       `json:"host_header,omitempty"`
	HashOnCookiePath   *string       `json:"hash_on_cookie_path,omitempty"`
	HashFallback       *string       `json:"hash_fallback,omitempty"`
	Algorithm          *string       `json:"algorithm,omitempty"`
	Slots              *int          `json:"slots,omitempty"`
	HashOn             *string       `json:"hash_on,omitempty"`
	HealthChecks       *HealthChecks `json:"healthchecks,omitempty"`
	ClientCertificate  *ItemID       `json:"client_certificate,omitempty"`
	ID                 *string       `json:"id,omitempty"`
	CreatedAt          *int          `json:"created_at,omitempty"`
	Tags               []string      `json:"tags,omitempty"`
	HashFallbackHeader interface{}   `json:"hash_fallback_header,omitempty"`
	HashOnCookie       interface{}   `json:"hash_on_cookie,omitempty"`
	HashOnHeader       interface{}   `json:"hash_on_header,omitempty"`
}
