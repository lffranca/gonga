package mapper

import (
	"github.com/kong/go-kong/kong"
	"github.com/lffranca/gonga/domain"
)

func OptionModelToEntity(item *kong.ListOpt) *domain.Option {
	if item == nil {
		return nil
	}

	return &domain.Option{
		Size:         &item.Size,
		Offset:       &item.Offset,
		Tags:         item.Tags,
		MatchAllTags: &item.MatchAllTags,
	}
}
