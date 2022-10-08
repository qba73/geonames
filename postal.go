package geonames

import (
	"fmt"
	"net/url"
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
	Position    Position
	CountryCode string
	PostalCode  string
	AdminCode1  string
}

// Get knows how to retrieve postal codes for the given place name and country code.
func (c Client) GetPostCode(place, country string) ([]PostalCode, error) {
	url, err := c.buildPostalURL(place, country)
	if err != nil {
		return nil, err
	}
	var pr postalResponse
	if err := c.get(url, &pr); err != nil {
		return nil, err
	}

	var postalCodes []PostalCode
	for _, pc := range pr.PostalCodes {
		p := PostalCode{
			PlaceName:  pc.PlaceName,
			AdminName1: pc.AdminName1,
			Position: Position{
				Lat:  pc.Lat,
				Long: pc.Lng,
			},
			PostalCode:  pc.PostalCode,
			CountryCode: pc.CountryCode,
			AdminCode1:  pc.AdminCode1,
		}
		postalCodes = append(postalCodes, p)
	}
	return postalCodes, nil
}

func (c Client) buildPostalURL(placeName, countryCode string) (string, error) {
	params := url.Values{
		"placename": {placeName},
		"country":   {countryCode},
		"username":  {c.userName},
	}
	basePostal := fmt.Sprintf("%s/postalCodeSearchJSON", c.baseURL)

	u, err := url.Parse(basePostal)
	if err != nil {
		return "", fmt.Errorf("parsing base url for postal, %w", err)
	}
	u.RawQuery = params.Encode()
	return u.String(), nil
}
