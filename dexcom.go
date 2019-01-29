package dexcom

import (
	"net/http"
	"time"
)

// Client represents the information needed to make requests to Dexcom. The Endpoint field can be overwritten to test
// using the sandbox. The Logger can be overwritten when you wish to see debugging type output.
type Client struct {
	Endpoint string
	Logger Logger

	c httpClient
}

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New will create a new Dexcom client ready to make requests to the API. The default request timeout is 5 seconds.
func New(accessToken string) *Client {
	c := Client{
		Endpoint: "https://api.dexcom.com",
		Logger: noOpLogger{},
		c: &authClient{
			c: http.Client{
				Timeout: time.Second * 5,
			},
			accessToken: accessToken,
		},
	}

	return &c
}