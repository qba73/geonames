package geonames

import (
	"fmt"
	"net/url"
	"strconv"
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
	Position    Position
	CountryCode string
	Rank        int
	Language    string
	Title       string
	URL         string
}

// GetPlace retrives geo coordinates for given place name and country code.
func (c Client) GetPlace(name, country string, maxResults int) ([]Geoname, error) {
	if maxResults < 1 {
		return nil, fmt.Errorf("invalid max results: %d", maxResults)
	}
	url, err := c.buildWikiURL(name, country, maxResults)
	if err != nil {
		return nil, err
	}

	var wr wikipediaResponse
	if err := c.get(url, &wr); err != nil {
		return nil, err
	}

	var geonames []Geoname
	for _, g := range wr.Geonames {
		geoname := Geoname{
			Summary:   g.Summary,
			Elevation: g.Elevation,
			GeoNameID: g.GeoNameID,
			Feature:   g.Feature,
			Position: Position{
				Lat:  g.Lat,
				Long: g.Lng,
			},
			CountryCode: g.CountryCode,
			Rank:        g.Rank,
			Language:    g.Lang,
			Title:       g.Title,
			URL:         g.WikipediaURL,
		}
		geonames = append(geonames, geoname)
	}
	return geonames, nil
}

func (c Client) buildWikiURL(place, country string, maxResults int) (string, error) {
	params := url.Values{
		"q":           []string{place},
		"title":       []string{place},
		"countryCode": []string{country},
		"maxRows":     []string{strconv.Itoa(maxResults)},
		"username":    []string{c.userName},
	}
	baseWiki := fmt.Sprintf("%s/wikipediaSearchJSON", c.baseURL)
	u, err := url.Parse(baseWiki)
	if err != nil {
		return "", fmt.Errorf("parsing wikipedia base url: %w", err)
	}
	u.RawQuery = params.Encode()
	return u.String(), nil
}
