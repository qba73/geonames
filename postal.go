package geonames

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"

	"github.com/shopspring/decimal"
)

type postalResponse struct {
	PostalCodes []struct {
		AdminCode1  string  `json:"adminCode1"`
		Lng         float64 `json:"lng"`
		CountryCode string  `json:"countryCode"`
		PostalCode  string  `json:"postalCode"`
		AdminName1  string  `json:"adminName1"`
		ISO31662    string  `json:"ISO3166-2"`
		PlaceName   string  `json:"placeName"`
		Lat         float64 `json:"lat"`
	} `json:"postalCodes"`
}

type PostalCode struct {
	PlaceName   string
	AdminName1  string
	Lat         string
	Long        string
	CountryCode string
	PostalCode  string
	AdminCode1  string
}

type PostalCodesService struct {
	cl *Client
}

// Get knows how to retrieve postal codes for the given place name and country code.
func (ps PostalCodesService) Get(place, country string) ([]PostalCode, error) {
	u, err := ps.makePostalURL(place, country)
	if err != nil {
		return nil, err
	}

	req, err := prepareGETRequest(u)
	if err != nil {
		return nil, err
	}

	res, err := ps.cl.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var pr postalResponse
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response body %w", err)
	}
	if err := json.Unmarshal(data, &pr); err != nil {
		return nil, fmt.Errorf("unmarshalling data, %w", err)
	}

	var postalCodes []PostalCode
	for _, pc := range pr.PostalCodes {
		p := PostalCode{
			PlaceName:   pc.PlaceName,
			AdminName1:  pc.AdminName1,
			Lat:         decimal.NewFromFloatWithExponent(pc.Lat, -4).String(),
			Long:        decimal.NewFromFloatWithExponent(pc.Lng, -4).String(),
			PostalCode:  pc.PostalCode,
			CountryCode: pc.CountryCode,
			AdminCode1:  pc.AdminCode1,
		}
		postalCodes = append(postalCodes, p)
	}
	return postalCodes, nil
}

func (ps PostalCodesService) makePostalURL(placeName, countryCode string) (string, error) {
	prms := url.Values{
		"placename": {placeName},
		"country":   {countryCode},
		"username":  {ps.cl.UserName},
	}
	basePostal := fmt.Sprintf("%s/%s", ps.cl.BaseURL, "postalCodeSearchJSON")
	return makeURL(basePostal, prms)
}
