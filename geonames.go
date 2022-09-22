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
	baseURL        = "http://api.geonames.org"
)

// Client is a client used for communicating
// with GeoNames web service.
type Client struct {
	// UserName is a user name chosen when registered for GeoNames.org
	UserName string

	UserAgent  string
	BaseURL    string
	HTTPClient *http.Client

	// Optional HTTP headers to set for each API request.
	Headers map[string]string

	// Geonames services
	Wikipedia   WikipediaService
	PostalCodes PostalCodesService
}

// NewClient knows how to create a client for GeoNames Web service.
// The user name has to be registered at the GeoNames.org website.
// HTTP requests without a valid username param will return 403 HTTP errors.
func NewClient(username string) *Client {
	c := Client{
		UserAgent: userAgent,
		UserName:  username,
		BaseURL:   baseURL,
		HTTPClient: &http.Client{
			Timeout: time.Second * 5,
		},
	}
	c.Wikipedia = WikipediaService{cl: &c}
	c.PostalCodes = PostalCodesService{cl: &c}
	return &c
}

// makeURL knows how to create encoded URL with provided query parameters.
func makeURL(base string, params url.Values) (string, error) {
	b, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("parsing base url, %v", err)
	}
	b.RawQuery = params.Encode()
	return b.String(), nil
}

// prepareGETRequest takes URL string and prepares HTTP Get request.
func prepareGETRequest(u string) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", userAgent)
	return req, nil
}
