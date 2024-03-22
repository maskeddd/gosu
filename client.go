package gosu

import (
	"context"
	"errors"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2/clientcredentials"
	"strconv"
)

type Client struct {
	httpClient *resty.Client
}

func NewClient(clientID int, clientSecret string) (*Client, error) {
	client := &Client{}

	ctx := context.Background()
	oauthConfig := clientcredentials.Config{
		ClientID:     strconv.Itoa(clientID),
		ClientSecret: clientSecret,
		TokenURL:     "https://osu.ppy.sh/oauth/token",
		Scopes:       []string{"public"},
	}

	client.httpClient = resty.NewWithClient(oauthConfig.Client(ctx)).SetBaseURL("https://osu.ppy.sh/api/v2")

	client.httpClient.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		switch resp.StatusCode() {
		case 404:
			return errors.New("not found")
		}

		return nil
	})

	return client, nil
}
