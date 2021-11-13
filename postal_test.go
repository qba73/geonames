package geonames_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/geonames"
)

func TestGetPostalCodesSingle(t *testing.T) {
	t.Parallel()

	ts := newTestServer(
		"testdata/response-geoname-postal-single.json",
		"/postalCodeSearchJSON?country=IE&placename=Castlebar&username=DummyUser",
		t,
	)
	defer ts.Close()

	client := geonames.NewClient("DummyUser")
	client.BaseURL = ts.URL

	place, country := "Castlebar", "IE"
	got, err := client.PostalCodes.Get(place, country)
	if err != nil {
		t.Fatalf("Get(%q, %q) got err %v", place, country, err)
	}
	want := geonames.PostalCode{
		PlaceName:   "Castlebar",
		AdminName1:  "Connacht",
		Lat:         53.85,
		Lng:         -9.3,
		CountryCode: "IE",
		PostalCode:  "F23",
		AdminCode1:  "C",
	}

	if !cmp.Equal(want, got.Codes[0]) {
		t.Errorf("Get(%q, %q) \n%s", place, country, cmp.Diff(want, got))
	}
}

func TestGetPostalCodesMultiple(t *testing.T) {
	t.Parallel()

	ts := newTestServer(
		"testdata/response-geoname-postal-multiple.json",
		"/postalCodeSearchJSON?country=IE&placename=Dublin&username=DummyUser",
		t,
	)
	defer ts.Close()

	client := geonames.NewClient("DummyUser")
	client.BaseURL = ts.URL

	place, country := "Dublin", "IE"
	got, err := client.PostalCodes.Get(place, country)
	if err != nil {
		t.Fatalf("GetCodes(%q, %q) got err %v", place, country, err)
	}

	want := geonames.PostalCodes{
		Codes: []geonames.PostalCode{
			{
				PlaceName:   "Dublin 1",
				AdminName1:  "Leinster",
				Lat:         53.353976,
				Lng:         -6.254537,
				CountryCode: "IE",
				PostalCode:  "D01",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 2",
				AdminName1:  "Leinster",
				Lat:         53.339971,
				Lng:         -6.254295,
				CountryCode: "IE",
				PostalCode:  "D02",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 3",
				AdminName1:  "Leinster",
				Lat:         53.364465,
				Lng:         -6.23776,
				CountryCode: "IE",
				PostalCode:  "D03",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 4",
				AdminName1:  "Leinster",
				Lat:         53.333435,
				Lng:         -6.233526,
				CountryCode: "IE",
				PostalCode:  "D04",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 5",
				AdminName1:  "Leinster",
				Lat:         53.384222,
				Lng:         -6.192128,
				CountryCode: "IE",
				PostalCode:  "D05",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 6",
				AdminName1:  "Leinster",
				Lat:         53.308787,
				Lng:         -6.263126,
				CountryCode: "IE",
				PostalCode:  "D06",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 7",
				AdminName1:  "Leinster",
				Lat:         53.361507,
				Lng:         -6.291792,
				CountryCode: "IE",
				PostalCode:  "D07",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 8",
				AdminName1:  "Leinster",
				Lat:         53.33455,
				Lng:         -6.273257,
				CountryCode: "IE",
				PostalCode:  "D08",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 9",
				AdminName1:  "Leinster",
				Lat:         53.381763,
				Lng:         -6.246501,
				CountryCode: "IE",
				PostalCode:  "D09",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 10",
				AdminName1:  "Leinster",
				Lat:         53.340906,
				Lng:         -6.354476,
				CountryCode: "IE",
				PostalCode:  "D10",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 11",
				AdminName1:  "Leinster",
				Lat:         53.389903,
				Lng:         -6.292976,
				CountryCode: "IE",
				PostalCode:  "D11",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 12",
				AdminName1:  "Leinster",
				Lat:         53.32203,
				Lng:         -6.316477,
				CountryCode: "IE",
				PostalCode:  "D12",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 13",
				AdminName1:  "Leinster",
				Lat:         53.394577,
				Lng:         -6.1495,
				CountryCode: "IE",
				PostalCode:  "D13",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 14",
				AdminName1:  "Leinster",
				Lat:         53.295987,
				Lng:         -6.259331,
				CountryCode: "IE",
				PostalCode:  "D14",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 15",
				AdminName1:  "Leinster",
				Lat:         53.383156,
				Lng:         -6.416518,
				CountryCode: "IE",
				PostalCode:  "D15",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 16",
				AdminName1:  "Leinster",
				Lat:         53.341884,
				Lng:         -6.278967,
				CountryCode: "IE",
				PostalCode:  "D16",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 17",
				AdminName1:  "Leinster",
				Lat:         53.400646,
				Lng:         -6.205763,
				CountryCode: "IE",
				PostalCode:  "D17",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 18",
				AdminName1:  "Leinster",
				Lat:         53.246902,
				Lng:         -6.177386,
				CountryCode: "IE",
				PostalCode:  "D18",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 20",
				AdminName1:  "Leinster",
				Lat:         53.35177,
				Lng:         -6.369332,
				CountryCode: "IE",
				PostalCode:  "D20",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 22",
				AdminName1:  "Leinster",
				Lat:         53.32751,
				Lng:         -6.400591,
				CountryCode: "IE",
				PostalCode:  "D22",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 24",
				AdminName1:  "Leinster",
				Lat:         53.285119,
				Lng:         -6.371327,
				CountryCode: "IE",
				PostalCode:  "D24",
				AdminCode1:  "L",
			},
			{
				PlaceName:   "Dublin 6W",
				AdminName1:  "Leinster",
				Lat:         53.308651,
				Lng:         -6.30119,
				CountryCode: "IE",
				PostalCode:  "D6W",
				AdminCode1:  "L",
			},
		},
	}
	if !cmp.Equal(want, got) {
		t.Errorf("Get(%q, %q) \n%s\n", place, country, cmp.Diff(want, got))
	}
}
