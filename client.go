// TODO: change the package name to something that makes sense for your service
// e.g. permissions - (the client this repo is inspired from)
package service_client

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// transportFunc is utility type for declaring a http.RoundTripper as a function literal
type transportFunc func(*http.Request) (*http.Response, error)

func (fn transportFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

type Client struct {
	http    *http.Client
	baseURL string
}

type clientOptions struct {
	client *http.Client
}

type ClientOption func(*clientOptions)

// An example method that hangs off the client. This should handle everything from creating the HTTP
// request to handling the errors and unmarshalling the response
//
// commented out you'll find an example of this flow - feel free to modify this to suit your needs
func (client Client) CallSomeAPIEndpoint(ctx context.Context, additionalArgs string) error {
	// req, err := http.NewRequestWithContext(ctx, "GET", client.baseURL+"/connections/"+providerName, nil)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to construct request: %w", err)
	// }

	// req.Header.Set("Accept", "application/json")

	// resp, err := client.http.Do(req)
	// if err != nil {
	// 	return "", fmt.Errorf("failed to get response: %w", err)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != 200 {
	// 	body, _ := io.ReadAll(resp.Body)
	// 	return "", fmt.Errorf("unexpected status-code %d: %q", resp.StatusCode, body)
	// }

	// var result struct {
	// 	ConnectionID string `json:"connection-id"`
	// }

	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	return "", fmt.Errorf("failed to decode response: %w", err)
	// }

	return errors.New("implement me")
}

// Returns a new client for your service
func NewClient(baseURL, userAgent string, options ...ClientOption) (*Client, error) {
	if userAgent == "" {
		return nil, errors.New("user-agent cannot be blank: required to identify client to external-credentials-service")
	}

	baseURL = strings.TrimRight(baseURL, "/")

	var opts clientOptions
	for _, apply := range options {
		apply(&opts)
	}

	if opts.client == nil {
		opts.client = http.DefaultClient
	}

	c := *opts.client

	transport := c.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	c.Transport = transportFunc(func(r *http.Request) (*http.Response, error) {
		r.Header.Set("User-Agent", userAgent)
		return transport.RoundTrip(r)
	})

	return &Client{
		http:    &c,
		baseURL: baseURL,
	}, nil
}

func WithClient(client *http.Client) ClientOption {
	return func(co *clientOptions) {
		co.client = client
	}
}
