package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lffranca/gonga/pkg/kong/sni"
	"path"
)

func (gateway *kong) DeleteSNI(sniNameOrID *string) error {
	if sniNameOrID == nil {
		return errors.New("invalid service name")
	}

	urlPath := path.Join(sni.URL, *sniNameOrID)

	if _, err := request(gateway, urlPath, HTTPMethodDELETE, nil, nil, nil); err != nil {
		return err
	}

	return nil
}

func (gateway *kong) UpsertSNI(sniNameOrID *string, updateItem *sni.Form) (*sni.SNI, error) {
	if sniNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(sni.URL, *sniNameOrID)

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

	var item sni.SNI
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) UpdateSNI(sniNameOrID *string, updateItem *sni.Form) (*sni.SNI, error) {
	if sniNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(sni.URL, *sniNameOrID)

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

	var item sni.SNI
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CreateSNI(createItem *sni.Form) (*sni.SNI, error) {
	body, errBody := json.Marshal(createItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, sni.URL, HTTPMethodPOST, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item sni.SNI
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) SNIByNameOrID(sniNameOrID *string) (*sni.SNI, error) {
	if sniNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(sni.URL, *sniNameOrID)

	response, errResponse := request(gateway, urlPath, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item sni.SNI
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ListAllSNIs(offset, tagNameFilter *string) (*sni.Response, error) {
	query := map[string]string{}
	if offset != nil {
		query["offset"] = *offset
	}

	if tagNameFilter != nil {
		query["tags"] = *tagNameFilter
	}

	response, errResponse := request(gateway, sni.URL, HTTPMethodGET, nil, query, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item sni.Response
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
