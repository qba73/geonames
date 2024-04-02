package geonames

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Client holds data required for communicating with the Geonames Web Services.
type Client struct {
	UserName   string
	UserAgent  string
	BaseURL    string
	HTTPClient *http.Client

	// Optional HTTP headers to set for each API request.
	Headers map[string][]string
}

const (
	libraryVersion = "0.1"
	userAgent      = "geonames/" + libraryVersion
)

// NewClient creates a new client for GeoNames Web service.
//
// The username has to be registered at the GeoNames.org website.
// HTTP requests without a valid username will return 403 HTTP errors.
func NewClient(username string) *Client {
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
	return &c
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

var DemoClient = &Client{
	UserName:   "demo",
	BaseURL:    "http://api.geonames.org",
	HTTPClient: http.DefaultClient,
	Headers: map[string][]string{
		"Content-Type": {"application/json"},
	},
}

var ClientFromEnv = &Client{
	UserName: os.Getenv("GEONAMES_USER"),
	BaseURL:  "http://api.geonames.org",
	HTTPClient: &http.Client{
		Timeout: 10 * time.Second,
	},
	Headers: map[string][]string{
		"Content-Type": {"application/json"},
	},
}
