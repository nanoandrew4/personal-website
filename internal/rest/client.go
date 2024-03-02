package rest

import (
	"errors"
	"github.com/go-resty/resty/v2"
)

type RestClient struct {
	client *resty.Client
}

func (r *RestClient) R() *resty.Request {
	return r.client.R()
}

func NewRestClient(c *resty.Client) *RestClient {
	c.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if !response.IsSuccess() {
			return errors.New(string(response.Body()))
		}
		return nil
	})
	return &RestClient{client: c}
}
