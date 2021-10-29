package domain

type Gateway struct {
	ID   *string `json:"id,omitempty" yaml:"id,omitempty"`
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
	Host *string `json:"host,omitempty" yaml:"host,omitempty"`
}
