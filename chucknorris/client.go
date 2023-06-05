package chucknorris

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const RequestURI = "https://api.chucknorris.io/jokes/random"

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout <= 0 {
		return nil, fmt.Errorf("timeout must be greater than zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &RoundTripLogger{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c *Client) GetRandomeJoke() (*Joke, error) {
	resp, err := c.client.Get(RequestURI)
	if err != nil {
		return nil, err
	}

	isSucces, err := c.checkCode(resp)
	if !isSucces {
		return nil, err
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	joke := Joke{}
	err = json.Unmarshal(data, &joke)
	if err != nil {
		return nil, err
	}

	return &joke, nil
}

func (c *Client) checkCode(r *http.Response) (bool, error) {
	if r.StatusCode < 200 || r.StatusCode >= 300 {
		return false, fmt.Errorf("resource %s is unavalible", RequestURI)
	}
	return true, nil
}
