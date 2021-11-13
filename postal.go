package geonames

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
)

type PostalCodes struct {
	Codes []PostalCode `json:"postalCodes"`
}

type PostalCode struct {
	PlaceName   string  `json:"placeName"`
	AdminName1  string  `json:"adminName1"`
	AdminName2  string  `json:"adminName2"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	CountryCode string  `json:"countryCode"`
	PostalCode  string  `json:"postalCode"`
	AdminCode1  string  `json:"adminCode1"`
	AdminCode2  string  `json:"adminCode2"`
}

type PostalCodesService struct {
	cl *Client
}

// Get knows how to retrieve postal codes
// for the given place name and country code.
func (ps PostalCodesService) Get(place, country string) (PostalCodes, error) {
	u, err := ps.makePostalURL(place, country)
	if err != nil {
		return PostalCodes{}, err
	}

	req, err := PrepareGETRequest(u)
	if err != nil {
		return PostalCodes{}, err
	}

	res, err := ps.cl.HTTPClient.Do(req)
	if err != nil {
		return PostalCodes{}, err
	}
	defer res.Body.Close()

	var pc PostalCodes
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return PostalCodes{}, fmt.Errorf("reading response body %w", err)
	}
	if err := json.Unmarshal(data, &pc); err != nil {
		return PostalCodes{}, fmt.Errorf("unmarshalling data, %w", err)
	}
	return pc, nil
}

func (ps PostalCodesService) makePostalURL(placeName, countryCode string) (string, error) {
	prms := url.Values{
		"placename": {placeName},
		"country":   {countryCode},
		"username":  {ps.cl.UserName},
	}
	basePostal := fmt.Sprintf("%s/%s", ps.cl.BaseURL, "postalCodeSearchJSON")
	return MakeURL(basePostal, prms)
}
