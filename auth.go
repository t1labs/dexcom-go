package dexcom

import (
	"fmt"
	"net/http"
)

type authClient struct {
	c http.Client
	accessToken string
}

// Do will add the access token to every request in the Authorization header
func (c *authClient) Do(req *http.Request) (*http.Response, error) {
	auth := fmt.Sprintf("Bearer %s", c.accessToken)
	req.Header.Set("Authorization", auth)
	return c.c.Do(req)
}