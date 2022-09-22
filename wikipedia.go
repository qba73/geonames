package geonames

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"

	"github.com/shopspring/decimal"
)

type wikipediaResponse struct {
	Geonames []struct {
		Summary      string  `json:"summary"`
		Elevation    int     `json:"elevation"`
		GeoNameID    int     `json:"geoNameId,omitempty"`
		Lng          float64 `json:"lng"`
		CountryCode  string  `json:"countryCode"`
		Rank         int     `json:"rank"`
		Lang         string  `json:"lang"`
		Title        string  `json:"title"`
		Lat          float64 `json:"lat"`
		WikipediaURL string  `json:"wikipediaUrl"`
		Feature      string  `json:"feature,omitempty"`
	} `json:"geonames"`
}

type Geoname struct {
	Summary     string
	Elevation   int
	GeoNameID   int
	Feature     string
	Lat         string
	Long        string
	CountryCode string
	Rank        int
	Lang        string
	Title       string
	URL         string
}

type WikipediaService struct {
	cl *Client
}

// Get knows how to retrive geo coordinates for
// the given place name and country code.
func (ws WikipediaService) Get(place, country string, maxResults int) ([]Geoname, error) {
	u, err := ws.makeWikiURL(place, country, maxResults)
	if err != nil {
		return nil, err
	}
	req, err := prepareGETRequest(u)
	if err != nil {
		return nil, err
	}

	res, err := ws.cl.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var wr wikipediaResponse
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body %w", err)
	}
	if err := json.Unmarshal(data, &wr); err != nil {
		return nil, fmt.Errorf("unmarshalling data, %w", err)
	}

	var gx []Geoname

	for _, g := range wr.Geonames {
		geoname := Geoname{
			Summary:     g.Summary,
			Elevation:   g.Elevation,
			GeoNameID:   g.GeoNameID,
			Feature:     g.Feature,
			Lat:         decimal.NewFromFloatWithExponent(g.Lat, -4).String(),
			Long:        decimal.NewFromFloatWithExponent(g.Lng, -4).String(),
			CountryCode: g.CountryCode,
			Rank:        g.Rank,
			Lang:        g.Lang,
			Title:       g.Title,
			URL:         g.WikipediaURL,
		}
		gx = append(gx, geoname)
	}

	return gx, nil
}

func (ws WikipediaService) makeWikiURL(place, country string, maxResults int) (string, error) {
	if maxResults < 1 {
		return "", fmt.Errorf("incorrect results limit: %q", maxResults)
	}
	prms := url.Values{
		"q":           []string{place},
		"title":       []string{place},
		"countryCode": []string{country},
		"maxRows":     []string{strconv.Itoa(maxResults)},
		"username":    []string{ws.cl.UserName},
	}
	base := fmt.Sprintf("%s/%s", ws.cl.BaseURL, "wikipediaSearchJSON")
	return makeURL(base, prms)
}
