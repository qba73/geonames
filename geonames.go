package geonames

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	libraryVersion = "0.1"
	userAgent      = "geonames/" + libraryVersion
	baseURL        = "http://api.geonames.org"
)

// Client is a client used for communicating with GeoNames web service.
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
		return "", fmt.Errorf("parsing base url, %w", err)
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

const (
	// 111km 100km
	// 11km 10km
	// 111m 100m
	// 11m 10m
	// 11cm 10cm

	Precision100km = 0
	Precision10km  = 1
	Precision1km   = 2
	Precision100m  = 3
	Precision10m   = 4
	Precision1m    = 5
	Precision10cm  = 6
	Precision1cm   = 7
	Precision1mm   = 8
)

type Position struct {
	Lat  float64
	Long float64
}

func LatLongWithPrecision(p Position, pr int) (string, string, error) {
	lat, err := coordinateWithPrecision(p.Lat, pr)
	if err != nil {
		return "", "", err
	}
	long, err := coordinateWithPrecision(p.Long, pr)
	if err != nil {
		return "", "", err
	}
	return lat, long, nil
}

func coordinateWithPrecision(c float64, precision int) (string, error) {
	if precision < 0 || precision > 8 {
		return "", fmt.Errorf("invalid precision %d, expected one of 0,1,2,3,4,5,6,7,8", precision)
	}
	coord := strconv.FormatFloat(c, 'f', precision, 64)
	return coord, nil
}
