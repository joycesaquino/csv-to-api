package internal

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

type Config struct {
	Url string `env:"SERVICE_URL,notEmpty"`
}

type Client struct {
	restyClient *resty.Client
}

const (
	contentTypeHeader = "Content-Type"
	contentTypeJson   = "application/json"
)

func (client Client) Post(ctx context.Context, body VisitorEventBody) error {

	resp, err := client.restyClient.
		R().
		SetContext(ctx).
		SetHeader(contentTypeHeader, contentTypeJson).
		SetHeader("5AJWT_AUTH", "eyJraWQiOm51bGwsImFsZyI6IlJTMjU2In0.eyJqdGkiOiJJWUZ6WnFuQUU3VWVDRmtBVUVXSWtRIiwiaWF0IjoxNjYzNjIwNjIxLCJpZCI6MTYwMDEyMSwicGVyc29uVVVJRCI6bnVsbCwibmFtZSI6IkpveWNlIEQgUyBBIiwiZmlyc3RuYW1lIjoiSm95Y2UiLCJ0aXR1bG8iOiJKb3ljZSIsImVtYWlsIjoiaW50ZWdyYWNhbythNzljNmIxNDg3YjljZjBhZjk3YzliODk0OGEwM2Y5MkBxdWludG9hbmRhci5jb20uYnIiLCJ0ZWxlZm9uZSI6Iis1NTAwNzk5ODYxNTcwIiwicm9sZXMiOiJhZG1pbixzdWRvIiwiY3JlYXRpb25EYXRlIjoiMTkvMDkvMjAyMiAyMDo1MCIsInVzZXJDcmVhdGVkQXQiOiIyMDIxLTA5LTE0VDEyOjQxOjQ2WiIsImV4cCI6MTY2NDIyNTQyMX0.byoEQQstUxBCv5ySBMB-_wvEVVxj7tNj0nW-i7xt2APElLUrjexbXDKRckLA69TtIzapclRmryszPszVh-ExlavRlVHnvJFOHMEbZmshYWYDqg7apeHKuu7d9_FXyH6BhQmfO75VRDWKoZJIC6e7oxlX-cF8eJOCi5T4pxKkZ5INqfyUOUtJjf1jQTUD6j2C9ZKgTX9OmZ-3v8bz9V5s1kEWjPdxCrwoIxhfV69n16z1sm_xWaAMEwfaiuPC--OY79iQuRmrZLpqA_2avusVu31ceCdVR1XCyNYlY5Gps-0UOqRk_TQSM5rDdiPxEvTj7efGoRXiHxzp4ve68pLOWA").
		SetBody(body).
		Post("/visitors/event")

	if err != nil {
		return err
	}

	if resp.IsError() {
		return fmt.Errorf("[ERROR] - StatusCode : %d , Error n: %s", resp.StatusCode(), err)
	}

	return nil
}

func NewClient() *Client {

	client := resty.New()
	client.RetryCount = 3
	client.RetryWaitTime = 10 * time.Second
	client.RetryConditions = []resty.RetryConditionFunc{
		func(response *resty.Response, err error) bool {
			return response.StatusCode() == 500
		},
	}
	client.SetBaseURL("https://apigw.forno.quintoandar.com.br/hubs-api")

	return &Client{restyClient: client}
}
