package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lffranca/gonga/pkg/kong/route"
	"path"
)

func (gateway *kong) DeleteRoute(routeNameOrID *string) error {
	if routeNameOrID == nil {
		return errors.New("invalid service name")
	}

	urlPath := path.Join(route.URL, *routeNameOrID)

	if _, err := request(gateway, urlPath, HTTPMethodDELETE, nil, nil, nil); err != nil {
		return err
	}

	return nil
}

func (gateway *kong) UpsertRoute(routeNameOrID *string, updateItem *route.Form) (*route.Route, error) {
	if routeNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(route.URL, *routeNameOrID)

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

	var item route.Route
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) UpdateRoute(routeNameOrID *string, updateItem *route.Form) (*route.Route, error) {
	if routeNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(route.URL, *routeNameOrID)

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

	var item route.Route
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CreateRoute(createItem *route.Form) (*route.Route, error) {
	body, errBody := json.Marshal(createItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, route.URL, HTTPMethodPOST, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item route.Route
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) RouteByNameOrID(routeNameOrID *string) (*route.Route, error) {
	if routeNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(route.URL, *routeNameOrID)

	response, errResponse := request(gateway, urlPath, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item route.Route
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ListAllRoutes(offset, tagNameFilter *string) (*route.Response, error) {
	query := map[string]string{}
	if offset != nil {
		query["offset"] = *offset
	}

	if tagNameFilter != nil {
		query["tags"] = *tagNameFilter
	}

	response, errResponse := request(gateway, route.URL, HTTPMethodGET, nil, query, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item route.Response
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
