package kong

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/lffranca/gonga/pkg/kong/consumer"
	"path"
)

func (gateway *kong) DeleteConsumer(consumerNameOrID *string) error {
	if consumerNameOrID == nil {
		return errors.New("invalid service name")
	}

	urlPath := path.Join(consumer.URL, *consumerNameOrID)

	if _, err := request(gateway, urlPath, HTTPMethodDELETE, nil, nil, nil); err != nil {
		return err
	}

	return nil
}

func (gateway *kong) UpsertConsumer(consumerNameOrID *string, updateItem *consumer.Form) (*consumer.Consumer, error) {
	if consumerNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(consumer.URL, *consumerNameOrID)

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

	var item consumer.Consumer
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) UpdateConsumer(consumerNameOrID *string, updateItem *consumer.Form) (*consumer.Consumer, error) {
	if consumerNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(consumer.URL, *consumerNameOrID)

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

	var item consumer.Consumer
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) CreateConsumer(createItem *consumer.Form) (*consumer.Consumer, error) {
	body, errBody := json.Marshal(createItem)
	if errBody != nil {
		return nil, errBody
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	}

	response, errResponse := request(gateway, consumer.URL, HTTPMethodPOST, bytes.NewReader(body), nil, headers)
	if errResponse != nil {
		return nil, errResponse
	}

	var item consumer.Consumer
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ConsumerByNameOrID(consumerNameOrID *string) (*consumer.Consumer, error) {
	if consumerNameOrID == nil {
		return nil, errors.New("invalid service name")
	}

	urlPath := path.Join(consumer.URL, *consumerNameOrID)

	response, errResponse := request(gateway, urlPath, HTTPMethodGET, nil, nil, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item consumer.Consumer
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}

func (gateway *kong) ListAllConsumers(offset, tagNameFilter *string) (*consumer.Response, error) {
	query := map[string]string{}
	if offset != nil {
		query["offset"] = *offset
	}

	if tagNameFilter != nil {
		query["tags"] = *tagNameFilter
	}

	response, errResponse := request(gateway, consumer.URL, HTTPMethodGET, nil, query, nil)
	if errResponse != nil {
		return nil, errResponse
	}

	var item consumer.Response
	if err := json.Unmarshal(response, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
