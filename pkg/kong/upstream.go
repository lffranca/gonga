package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lffranca/gonga/pkg/kong/upstream"
	"path"
)

func (gateway *kong) DeleteUpstream(upstreamNameOrID *string) error {
	if upstreamNameOrID == nil {
		return errors.New("invalid service name")
	}

	urlPath := path.Join(upstream.URL, *upstreamNameOrID)

	if _, err := request(gateway, urlPath, HTTPMethodDELETE, nil, nil, nil); err != nil {
		return err
	}

	return nil
}

func (gateway *kong) UpsertUpstream(upstreamNameOrID *string, updateItem *upstream.Form) (*upstream.Upstream, error) {
	if upstreamNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(upstream.URL, *upstreamNameOrID)

	body, errBody := json.Marshal(updateItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, urlPath, HTTPMethodPATCH, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item upstream.Upstream
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) UpdateUpstream(upstreamNameOrID *string, updateItem *upstream.Form) (*upstream.Upstream, error) {
	if upstreamNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(upstream.URL, *upstreamNameOrID)

	body, errBody := json.Marshal(updateItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, urlPath, HTTPMethodPATCH, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item upstream.Upstream
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CreateUpstream(createItem *upstream.Form) (*upstream.Upstream, error) {
	body, errBody := json.Marshal(createItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, upstream.URL, HTTPMethodPOST, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item upstream.Upstream
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) UpstreamByNameOrID(upstreamNameOrID *string) (*upstream.Upstream, error) {
	if upstreamNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(upstream.URL, *upstreamNameOrID)

	response, errResponse := request(gateway, urlPath, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item upstream.Upstream
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ListAllUpstreams(offset, tagNameFilter *string) (*upstream.Response, error) {
	query := map[string]string{}
	if offset != nil {
		query["offset"] = *offset
	}

	if tagNameFilter != nil {
		query["tags"] = *tagNameFilter
	}

	response, errResponse := request(gateway, upstream.URL, HTTPMethodGET, nil, query, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item upstream.Response
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
