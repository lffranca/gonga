package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lffranca/gonga/pkg/kong/cacertificate"
	"path"
)

func (gateway *kong) DeleteCACertificate(caCertificateID *string) error {
	if caCertificateID == nil {
		return errors.New("invalid service name")
	}

	urlPath := path.Join(cacertificate.URL, *caCertificateID)

	if _, err := request(gateway, urlPath, HTTPMethodDELETE, nil, nil, nil); err != nil {
		return err
	}

	return nil
}

func (gateway *kong) UpsertCACertificate(caCertificateID *string, updateItem *cacertificate.Form) (*cacertificate.CACertificate, error) {
	if caCertificateID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(cacertificate.URL, *caCertificateID)

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

	var item cacertificate.CACertificate
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) UpdateCACertificate(caCertificateID *string, updateItem *cacertificate.Form) (*cacertificate.CACertificate, error) {
	if caCertificateID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(cacertificate.URL, *caCertificateID)

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

	var item cacertificate.CACertificate
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CreateCACertificate(createItem *cacertificate.Form) (*cacertificate.CACertificate, error) {
	body, errBody := json.Marshal(createItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, cacertificate.URL, HTTPMethodPOST, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item cacertificate.CACertificate
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CACertificateByNameOrID(caCertificateID *string) (*cacertificate.CACertificate, error) {
	if caCertificateID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(cacertificate.URL, *caCertificateID)

	response, errResponse := request(gateway, urlPath, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item cacertificate.CACertificate
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ListAllCACertificates(offset, tagNameFilter *string) (*cacertificate.Response, error) {
	query := map[string]string{}
	if offset != nil {
		query["offset"] = *offset
	}

	if tagNameFilter != nil {
		query["tags"] = *tagNameFilter
	}

	response, errResponse := request(gateway, cacertificate.URL, HTTPMethodGET, nil, query, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item cacertificate.Response
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

