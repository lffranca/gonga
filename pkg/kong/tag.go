package kong

import (
	"encoding/json"
	"github.com/lffranca/gonga/pkg/kong/tag"
	"path"
)

func (gateway *kong) ListAllTags(offset, tagNameFilter *string) (*tag.Response, error) {
	var query map[string]string
	if offset != nil {
		query = map[string]string{
			"offset": *offset,
		}
	}

	urlPath := tag.URL
	if tagNameFilter != nil {
		urlPath = path.Join(urlPath, *tagNameFilter)
	}

	response, errResponse := request(gateway, urlPath, HTTPMethodGET, nil, query, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item tag.Response
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
