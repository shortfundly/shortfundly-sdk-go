package shortfundly

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// CRequest defines the request given by the client to Shortfundly
type CRequest struct {
	Method string
	Path   string
}

// sendRequest makes the request to Shortfundly's API
func (s *Shortfundly) sendRequest(r CRequest, data interface{}) error {
	req, err := http.NewRequest(r.Method, fmt.Sprintf("%s/%s", s.Host, r.Path), strings.NewReader(""))
	if err != nil {
		return Error{ErrMessage: fmt.Sprintf("Unable to create the request: %s", err)}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-API-KEY", s.APIKey)

	httpClient := &http.Client{Timeout: s.Timeout}

	resp, err := httpClient.Do(req)
	if err != nil {
		return Error{ErrMessage: fmt.Sprintf("Failed to make request: %s", err)}
	}

	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Error{ErrMessage: fmt.Sprintf("Could not read response: %s", err)}
	}

	if resp.StatusCode == http.StatusOK {
		return json.Unmarshal(contents, &data)
	}

	errg := Error{}
	json.Unmarshal(contents, &errg)
	return Error{ErrMessage: errg.ErrMessage}
}
