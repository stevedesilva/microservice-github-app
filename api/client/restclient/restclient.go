package restclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	enabledMocks bool
	mocks        = make(map[string]*Mock, 0)
)

type Mock struct {
	Response   *http.Response
	Err        error
	Url        string
	HttpMethod string
}

func getMockID(httpMethod, URL string) string {
	return fmt.Sprintf("%s %s", httpMethod, URL)
}

// StartMockups func
func StartMockups() {
	enabledMocks = true
}

// StopMockups func
func StopMockups() {
	enabledMocks = false
}

// AddMockup func
func AddMockup(m Mock) {
	mocks[getMockID(m.HttpMethod, m.Url)] = &m
}

// FlushMockup func {
func FlushMockup() {
	mocks = make(map[string]*Mock, 0)
}

// Post func
func Post(url string, body interface{}, headers http.Header) (*http.Response, error) {
	if enabledMocks {
		if mock, ok := mocks[getMockID(http.MethodPost, url)]; ok {
			return mock.Response, mock.Err
		}
		return nil, fmt.Errorf("no error found for given request with url %s", url)
	}

	jsonBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(jsonBytes))
	request.Header = headers

	client := http.Client{}

	return client.Do(request)
}
