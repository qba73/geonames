package geonames

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	libraryVersion = "0.1"
	userAgent      = "geonames/" + libraryVersion
)

type option func(*Client) error

// WithHTTPClient configures the Geonames HTTP Client.
func WithHTTPClient(h *http.Client) option {
	return func(c *Client) error {
		if h == nil {
			return errors.New("nil http Client")
		}
		c.HTTPClient = h
		return nil
	}
}

// WithBaseURL configures a base URL for the Geonames client.
func WithBaseURL(url string) option {
	return func(c *Client) error {
		if url == "" {
			return errors.New("nil baseURL")
		}
		c.BaseURL = url
		return nil
	}
}

// WithHTTPHeader configures custom HTTP Headers used by the Geonames client.
func WithHTTPHeaders(header http.Header) option {
	return func(c *Client) error {
		if header == nil {
			return errors.New("nil HTTP headers")
		}
		c.Headers = header
		return nil
	}
}

// Client holds data required for communicating with the Geonames Web Services.
type Client struct {
	UserName   string
	UserAgent  string
	BaseURL    string
	HTTPClient *http.Client

	// Optional HTTP headers to set for each API request.
	Headers map[string][]string
}

var DefaultClient = Client{
	UserName:   "demo",
	BaseURL:    "http://api.geonames.org",
	HTTPClient: http.DefaultClient,
	Headers:    map[string][]string{"Content-Type": {"application/json"}},
}

// NewDemoClient creates a demo client. Demo client
// does not require username registration at geonames.org website.
func NewDemoClient() *Client {
	return &DefaultClient
}

// NewClient creates a new client for GeoNames Web service.
//
// The username has to be registered at the GeoNames.org website.
// HTTP requests without a valid username will return 403 HTTP errors.
func NewClient(username string, options ...option) (*Client, error) {
	c := Client{
		UserName:  username,
		UserAgent: userAgent,
		BaseURL:   "http://api.geonames.org",
		HTTPClient: &http.Client{
			Timeout: time.Second * 5,
		},
		Headers: map[string][]string{
			"User-Agent":   {userAgent},
			"Content-Type": {"application/json"},
		},
	}
	for _, opt := range options {
		if err := opt(&c); err != nil {
			return nil, fmt.Errorf("creating geonames client: %w", err)
		}
	}
	return &c, nil
}

func (c Client) get(ctx context.Context, url string, data any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("sending GET request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("got response code: %v", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("reading response body: %w", err)
	}

	if err := json.Unmarshal(body, data); err != nil {
		return fmt.Errorf("unmarshaling response body: %w", err)
	}
	return nil
}

// Position holds information about Lat and Long.
type Position struct {
	Lat float64
	Lng float64
}
