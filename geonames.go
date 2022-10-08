package geonames

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	libraryVersion = "0.1"
	userAgent      = "geonames/" + libraryVersion
)

type option func(*Client) error

func WithHTTPClient(h *http.Client) option {
	return func(c *Client) error {
		if h == nil {
			return errors.New("nil http Client")
		}
		c.httpClient = h
		return nil
	}
}

func WithBaseURL(url string) option {
	return func(c *Client) error {
		if url == "" {
			return errors.New("nil baseURL")
		}
		c.baseURL = url
		return nil
	}
}

func WithHTTPHeaders(header http.Header) option {
	return func(c *Client) error {
		if header == nil {
			return errors.New("nil HTTP headers")
		}
		c.headers = header
		return nil
	}
}

func WithUserName(username string) option {
	return func(c *Client) error {
		if username == "" {
			return errors.New("nil username")
		}
		c.userName = username
		return nil
	}
}

// Client is a client used for communicating with GeoNames web service.
type Client struct {
	userName   string
	userAgent  string
	baseURL    string
	httpClient *http.Client

	// Optional HTTP headers to set for each API request.
	headers map[string][]string
}

// NewClient knows how to create a client for GeoNames Web service.
// The username has to be registered at the GeoNames.org website.
// HTTP requests without a valid username will return 403 HTTP errors.
func NewClient(options ...option) (*Client, error) {
	c := Client{
		userName:  os.Getenv("GEONAMES_USER"),
		userAgent: userAgent,
		baseURL:   "http://api.geonames.org",
		httpClient: &http.Client{
			Timeout: time.Second * 5,
		},
		headers: map[string][]string{
			"User-Agent":   {userAgent},
			"Content-Type": {"application/json"}},
	}

	for _, opt := range options {
		if err := opt(&c); err != nil {
			return nil, fmt.Errorf("creating geonames client: %w", err)
		}
	}
	return &c, nil
}

func (c Client) get(url string, data interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("creating request: %w", err)
	}
	res, err := c.httpClient.Do(req)
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

type Position struct {
	Lat  float64
	Long float64
}
