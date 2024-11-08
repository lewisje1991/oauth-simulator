package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	client := getClient()

	url := &url.URL{
		Scheme: "http",
		Host:   "localhost:8080",
		Path:   "/token-echo",
	}

	for {
		time.Sleep(1 * time.Second)
		resp, err := client.Do(&http.Request{
			Method: "POST",
			URL:    url,
		})

		fmt.Println(resp.StatusCode, err)
	}
}

func getClient() *http.Client {
	ctx := context.Background()

	conf := &clientcredentials.Config{
		ClientID:     "blah",
		ClientSecret: "blah",
		TokenURL:     "http://localhost:8080/token",
		AuthStyle:    oauth2.AuthStyleInParams,
	}

	return conf.Client(ctx)
}
