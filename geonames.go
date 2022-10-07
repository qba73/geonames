package geonames

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	libraryVersion = "0.1"
	userAgent      = "geonames/" + libraryVersion
)

// Client is a client used for communicating with GeoNames web service.
type Client struct {
	// UserName is a user name chosen when registered for GeoNames.org
	UserName   string
	UserAgent  string
	BaseURL    string
	HTTPClient *http.Client

	// Optional HTTP headers to set for each API request.
	Headers map[string][]string
}

// NewClient knows how to create a client for GeoNames Web service.
// The user name has to be registered at the GeoNames.org website.
// HTTP requests without a valid username param will return 403 HTTP errors.
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
			"Content-Type": {"application/json"}},
	}
	return &c
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
	req.Header = c.Headers
	return req, nil
}

type Position struct {
	Lat  float64
	Long float64
}
