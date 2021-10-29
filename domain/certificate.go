package domain

import "time"

type Certificate struct {
	ID        *string    `json:"id,omitempty" yaml:"id,omitempty"`
	Cert      *string    `json:"cert,omitempty" yaml:"cert,omitempty"`
	Key       *string    `json:"key,omitempty" yaml:"key,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty" yaml:"created_at,omitempty"`
	SNIs      []*string  `json:"snis,omitempty" yaml:"snis,omitempty"`
	Tags      []*string  `json:"tags,omitempty" yaml:"tags,omitempty"`
}
