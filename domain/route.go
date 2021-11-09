package domain

import "time"

type Route struct {
	ID                      *string             `json:"id,omitempty" yaml:"id,omitempty"`
	Name                    *string             `json:"name,omitempty" yaml:"name,omitempty"`
	Hosts                   []*string           `json:"hosts,omitempty" yaml:"hosts,omitempty"`
	Headers                 map[string][]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	Methods                 []*string           `json:"methods,omitempty" yaml:"methods,omitempty"`
	Paths                   []*string           `json:"paths,omitempty" yaml:"paths,omitempty"`
	PathHandling            *string             `json:"path_handling,omitempty" yaml:"path_handling,omitempty"`
	PreserveHost            *bool               `json:"preserve_host,omitempty" yaml:"preserve_host,omitempty"`
	Protocols               []*string           `json:"protocols,omitempty" yaml:"protocols,omitempty"`
	RegexPriority           *int                `json:"regex_priority,omitempty" yaml:"regex_priority,omitempty"`
	Service                 *Service            `json:"service,omitempty" yaml:"service,omitempty"`
	StripPath               *bool               `json:"strip_path,omitempty" yaml:"strip_path,omitempty"`
	CreatedAt               *time.Time          `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	UpdatedAt               *time.Time          `json:"updated_at,omitempty" yaml:"updated_at,omitempty"`
	SNIs                    []*string           `json:"snis,omitempty" yaml:"snis,omitempty"`
	Sources                 []*CIDRPort         `json:"sources,omitempty" yaml:"sources,omitempty"`
	Destinations            []*CIDRPort         `json:"destinations,omitempty" yaml:"destinations,omitempty"`
	Tags                    []*string           `json:"tags,omitempty" yaml:"tags,omitempty"`
	HTTPSRedirectStatusCode *int                `json:"https_redirect_status_code,omitempty" yaml:"https_redirect_status_code,omitempty"`
	RequestBuffering        *bool               `json:"request_buffering,omitempty" yaml:"request_buffering,omitempty"`
	ResponseBuffering       *bool               `json:"response_buffering,omitempty" yaml:"response_buffering,omitempty"`
}
