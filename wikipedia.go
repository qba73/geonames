package geonames

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strconv"
)

type WikiResponse struct {
	Geonames []Geoname `json:"geonames"`
}

type Geoname struct {
	Summary      string  `json:"summary"`
	Elevation    int     `json:"elevation"`
	GeoNameID    int     `json:"geoNameId"`
	Feature      string  `json:"feature"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	CountryCode  string  `json:"countryCode"`
	Rank         int     `json:"rank"`
	ThumbnailImg string  `json:"thumbnailImg"`
	Lang         string  `json:"lang"`
	Title        string  `json:"title"`
	WikipediaURL string  `json:"wikipediaUrl"`
}

type WikipediaService struct {
	cl *Client
}

// Get knows how to retrive geo coordinates for
// the given place name and country code.
func (ws WikipediaService) Get(place, country string, maxResults int) (WikiResponse, error) {
	u, err := ws.makeWikiURL(place, country, maxResults)
	if err != nil {
		return WikiResponse{}, err
	}
	req, err := PrepareGETRequest(u)
	if err != nil {
		return WikiResponse{}, err
	}

	res, err := ws.cl.HTTPClient.Do(req)
	if err != nil {
		return WikiResponse{}, err
	}
	defer res.Body.Close()

	var wr WikiResponse
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return WikiResponse{}, fmt.Errorf("reading response body %w", err)
	}

	if err := json.Unmarshal(data, &wr); err != nil {
		return WikiResponse{}, fmt.Errorf("unmarshalling data, %w", err)
	}
	return wr, nil
}

func (ws WikipediaService) makeWikiURL(place, country string, maxResults int) (string, error) {
	if maxResults < 1 {
		return "", fmt.Errorf("max number of results should be min 1, got %q", maxResults)
	}
	prms := url.Values{
		"q":           []string{place},
		"title":       []string{place},
		"countryCode": []string{country},
		"maxRows":     []string{strconv.Itoa(maxResults)},
		"username":    []string{ws.cl.UserName},
	}
	base := fmt.Sprintf("%s/%s", ws.cl.BaseURL, "wikipediaSearchJSON")
	return MakeURL(base, prms)
}
