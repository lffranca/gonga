package route

const URL = "routes"

type Response struct {
	Offset *string `json:"offset"`
	Data   []Route `json:"data"`
}

type ItemID struct {
	ID   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Form struct {
	Name                    *string             `json:"name,omitempty"`
	Protocols               []string            `json:"protocols,omitempty"`
	Methods                 []string            `json:"methods,omitempty"`
	Hosts                   []string            `json:"hosts,omitempty"`
	Paths                   []string            `json:"paths,omitempty"`
	Headers                 map[string][]string `json:"headers,omitempty"`
	HTTPSRedirectStatusCode *int                `json:"https_redirect_status_code,omitempty"`
	RegexPriority           *int                `json:"regex_priority,omitempty"`
	StripPath               *bool               `json:"strip_path,omitempty"`
	PathHandling            *string             `json:"path_handling,omitempty"`
	PreserveHost            *bool               `json:"preserve_host,omitempty"`
	RequestBuffering        *bool               `json:"request_buffering,omitempty"`
	ResponseBuffering       *bool               `json:"response_buffering,omitempty"`
	SNIs                    []string            `json:"snis,omitempty"`
	Sources                 []string            `json:"sources,omitempty"`
	Destinations            []string            `json:"destinations,omitempty"`
	Tags                    []string            `json:"tags,omitempty"`
	Service                 *ItemID             `json:"service,omitempty"`
}

type Route struct {
	Name                    *string             `json:"name,omitempty"`
	Protocols               []string            `json:"protocols,omitempty"`
	PathHandling            *string             `json:"path_handling,omitempty"`
	Paths                   []string            `json:"paths,omitempty"`
	Methods                 []string            `json:"methods,omitempty"`
	Sources                 []string            `json:"sources,omitempty"`
	Destinations            []string            `json:"destinations,omitempty"`
	SNIs                    []string            `json:"snis,omitempty"`
	StripPath               *bool               `json:"strip_path,omitempty"`
	Tags                    []string            `json:"tags,omitempty"`
	RequestBuffering        *bool               `json:"request_buffering,omitempty"`
	ResponseBuffering       *bool               `json:"response_buffering,omitempty"`
	Headers                 map[string][]string `json:"headers,omitempty"`
	HTTPSRedirectStatusCode *int                `json:"https_redirect_status_code,omitempty"`
	PreserveHost            *bool               `json:"preserve_host,omitempty"`
	Service                 *ItemID             `json:"service,omitempty"`
	Hosts                   []string            `json:"hosts,omitempty"`
	ID                      *string             `json:"id,omitempty"`
	RegexPriority           *int                `json:"regex_priority,omitempty"`
	CreatedAt               *int                `json:"created_at,omitempty"`
	UpdatedAt               *int                `json:"updated_at,omitempty"`
}
