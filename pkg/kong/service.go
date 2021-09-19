package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lffranca/gonga/pkg/kong/service"
	"path"
)

func (gateway *kong) DeleteService(serviceNameOrID *string) error {
	if serviceNameOrID == nil {
		return errors.New("invalid service name")
	}

	urlPath := path.Join(service.URL, *serviceNameOrID)

	if _, err := request(gateway, urlPath, HTTPMethodDELETE, nil, nil, nil); err != nil {
		return err
	}

	return nil
}

func (gateway *kong) UpsertService(serviceNameOrID *string, updateItem *service.Form) (*service.Service, error) {
	if serviceNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(service.URL, *serviceNameOrID)

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

	var item service.Service
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) UpdateService(serviceNameOrID *string, updateItem *service.Form) (*service.Service, error) {
	if serviceNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(service.URL, *serviceNameOrID)

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

	var item service.Service
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CreateService(createItem *service.Form) (*service.Service, error) {
	body, errBody := json.Marshal(createItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, service.URL, HTTPMethodPOST, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item service.Service
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ServiceByNameOrID(serviceNameOrID *string) (*service.Service, error) {
	if serviceNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(service.URL, *serviceNameOrID)

	response, errResponse := request(gateway, urlPath, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item service.Service
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ListAllServices(offset, tagNameFilter *string) (*service.Response, error) {
	query := map[string]string{}
	if offset != nil {
		query["offset"] = *offset
	}

	if tagNameFilter != nil {
		query["tags"] = *tagNameFilter
	}

	response, errResponse := request(gateway, service.URL, HTTPMethodGET, nil, query, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item service.Response
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
