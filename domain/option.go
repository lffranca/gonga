package domain

type Option struct {
	Size         *int      `json:"size,omitempty" yaml:"size,omitempty"`
	Offset       *string   `json:"offset,omitempty" yaml:"offset,omitempty"`
	Tags         []*string `json:"tags,omitempty" yaml:"tags,omitempty"`
	MatchAllTags *bool     `json:"match_all_tags,omitempty" yaml:"match_all_tags,omitempty"`
}
