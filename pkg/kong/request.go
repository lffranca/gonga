package kong

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path"
)

func request(gateway *kong, pathService string, method HTTPMethod, body io.Reader, query, headers map[string]string) ([]byte, error) {
	gateway.url.Path = path.Join(gateway.url.Path, pathService)

	q := gateway.url.Query()
	for key, value := range query {
		q.Set(key, value)
	}
	gateway.url.RawQuery = q.Encode()

	req, errReq := http.NewRequest(string(method), gateway.url.String(), body)
	if errReq != nil {
		return nil, errReq
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Println("resp.Body.Close(): ", err)
		}
	}()

	responseBody, errResponse := ioutil.ReadAll(resp.Body)
	if errResponse != nil {
		return nil, errResponse
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		var errMessage MessageError
		if err := json.Unmarshal(responseBody, &errMessage); err != nil {
			return nil, fmt.Errorf("is not object \"MessageError\" - status code: %d - error: %s",
				resp.StatusCode, err)
		}

		return nil, errors.New(errMessage.Message)
	}

	return responseBody, nil
}
