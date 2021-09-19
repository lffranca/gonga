package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lffranca/gonga/pkg/kong/certificate"
	"path"
)

func (gateway *kong) DeleteCertificate(certificateID *string) error {
	if certificateID == nil {
		return errors.New("invalid service name")
	}

	urlPath := path.Join(certificate.URL, *certificateID)

	if _, err := request(gateway, urlPath, HTTPMethodDELETE, nil, nil, nil); err != nil {
		return err
	}

	return nil
}

func (gateway *kong) UpsertCertificate(certificateID *string, updateItem *certificate.Form) (*certificate.Certificate, error) {
	if certificateID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(certificate.URL, *certificateID)

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

	var item certificate.Certificate
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) UpdateCertificate(certificateID *string, updateItem *certificate.Form) (*certificate.Certificate, error) {
	if certificateID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(certificate.URL, *certificateID)

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

	var item certificate.Certificate
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CreateCertificate(createItem *certificate.Form) (*certificate.Certificate, error) {
	body, errBody := json.Marshal(createItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, certificate.URL, HTTPMethodPOST, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item certificate.Certificate
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CertificateByNameOrID(certificateID *string) (*certificate.Certificate, error) {
	if certificateID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(certificate.URL, *certificateID)

	response, errResponse := request(gateway, urlPath, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item certificate.Certificate
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ListAllCertificates(offset, tagNameFilter *string) (*certificate.Response, error) {
	query := map[string]string{}
	if offset != nil {
		query["offset"] = *offset
	}

	if tagNameFilter != nil {
		query["tags"] = *tagNameFilter
	}

	response, errResponse := request(gateway, certificate.URL, HTTPMethodGET, nil, query, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item certificate.Response
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

