package dexcom

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type interceptor struct {
	c http.Client
	accessToken string
}

// Do will add the access token to every request in the Authorization header
func (c *interceptor) Do(req *http.Request) (*http.Response, error) {
	auth := fmt.Sprintf("Bearer %s", c.accessToken)
	req.Header.Set("Authorization", auth)

	resp, err := c.c.Do(req)
	if err != nil {
		return nil, err
	}

	// Return the results if we have a good status code
	if resp.StatusCode < http.StatusBadRequest {
		return resp, nil
	}

	var e Error
	err = json.NewDecoder(resp.Body).Decode(&e)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return nil, e
}