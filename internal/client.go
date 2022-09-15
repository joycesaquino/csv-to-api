package internal

import (
	"context"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

type Config struct {
	Url   string `env:"SERVICE_URL,notEmpty"`
	Token string `env:"SERVICE_TOKEN,notEmpty"`
}

type Client struct {
	restyClient *resty.Client
}

const (
	contentTypeHeader = "Content-Type"
	contentTypeJson   = "application/json"
)

func (client Client) Post(ctx context.Context, body VisitorEvent) error {

	resp, err := client.restyClient.
		R().
		SetContext(ctx).
		SetHeader(contentTypeHeader, contentTypeJson).
		SetBody(body).
		Post("/visitor-events")

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("[ERROR] - StatusCode : %d , Error n: %s", resp.StatusCode(), err)
	}

	return nil
}

func NewClient() *Client {
	var config Config
	if err := env.Parse(&config); err != nil {
		log.Fatalf("[ERROR] - Error on configure env on http client: %s", err)
	}

	client := resty.New()
	client.RetryCount = 3
	client.RetryWaitTime = 10 * time.Second
	client.RetryConditions = []resty.RetryConditionFunc{
		func(response *resty.Response, err error) bool {
			return response.StatusCode() == 500
		},
	}
	client.SetBaseURL(config.Url)
	client.SetAuthScheme("Basic")
	client.SetAuthToken(config.Token)

	return &Client{restyClient: client}
}
