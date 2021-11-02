package presenter

type OptionQuery struct {
	Size         *int      `form:"size"`
	Offset       *string   `form:"offset"`
	Tags         []*string `form:"tags"`
	MatchAllTags *bool     `form:"match_all_tags"`
}
