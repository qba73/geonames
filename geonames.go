package geonames

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
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

// Client is a client used for communicating with GeoNames web service.
type Client struct {
	// UserName is a user name chosen when registered for GeoNames.org
	userName   string
	userAgent  string
	baseURL    string
	httpClient *http.Client

	// Optional HTTP headers to set for each API request.
	headers map[string][]string
}

// NewClient knows how to create a client for GeoNames Web service.
// The user name has to be registered at the GeoNames.org website.
// HTTP requests without a valid username param will return 403 HTTP errors.
func NewClient(username string, options ...option) (*Client, error) {
	c := Client{
		userName:  username,
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

// makeURL knows how to create encoded URL with provided query parameters.
func makeURL(base string, params url.Values) (string, error) {
	b, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("parsing base url, %w", err)
	}
	b.RawQuery = params.Encode()
	return b.String(), nil
}

// prepareGETRequest takes URL string and prepares HTTP Get request.
func (c Client) prepareGETRequest(u string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header = c.headers
	return req, nil
}

type Position struct {
	Lat  float64
	Long float64
}
