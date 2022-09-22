package geonames_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/geonames"
)

func TestGetPostalCodesReturnsSingleValueOnValidInput(t *testing.T) {
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
		t.Fatal(err)
	}

	want := []geonames.PostalCode{
		{
			PlaceName:   "Castlebar",
			AdminName1:  "Connacht",
			Lat:         "53.85",
			Long:        "-9.3",
			CountryCode: "IE",
			PostalCode:  "F23",
			AdminCode1:  "C",
		},
	}

	if len(got) != 1 {
		t.Fatalf("want one postal code, got %#v", got)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetPostalCodesReturnsMultipleValuesOnValidInput(t *testing.T) {
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

	want := []geonames.PostalCode{
		{
			PlaceName:   "Dublin 1",
			AdminName1:  "Leinster",
			Lat:         "53.354",
			Long:        "-6.2545",
			CountryCode: "IE",
			PostalCode:  "D01",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 2",
			AdminName1:  "Leinster",
			Lat:         "53.34",
			Long:        "-6.2543",
			CountryCode: "IE",
			PostalCode:  "D02",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 3",
			AdminName1:  "Leinster",
			Lat:         "53.3645",
			Long:        "-6.2378",
			CountryCode: "IE",
			PostalCode:  "D03",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 4",
			AdminName1:  "Leinster",
			Lat:         "53.3334",
			Long:        "-6.2335",
			CountryCode: "IE",
			PostalCode:  "D04",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 5",
			AdminName1:  "Leinster",
			Lat:         "53.3842",
			Long:        "-6.1921",
			CountryCode: "IE",
			PostalCode:  "D05",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 6",
			AdminName1:  "Leinster",
			Lat:         "53.3088",
			Long:        "-6.2631",
			CountryCode: "IE",
			PostalCode:  "D06",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 7",
			AdminName1:  "Leinster",
			Lat:         "53.3615",
			Long:        "-6.2918",
			CountryCode: "IE",
			PostalCode:  "D07",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 8",
			AdminName1:  "Leinster",
			Lat:         "53.3346",
			Long:        "-6.2733",
			CountryCode: "IE",
			PostalCode:  "D08",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 9",
			AdminName1:  "Leinster",
			Lat:         "53.3818",
			Long:        "-6.2465",
			CountryCode: "IE",
			PostalCode:  "D09",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 10",
			AdminName1:  "Leinster",
			Lat:         "53.3409",
			Long:        "-6.3545",
			CountryCode: "IE",
			PostalCode:  "D10",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 11",
			AdminName1:  "Leinster",
			Lat:         "53.3899",
			Long:        "-6.293",
			CountryCode: "IE",
			PostalCode:  "D11",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 12",
			AdminName1:  "Leinster",
			Lat:         "53.322",
			Long:        "-6.3165",
			CountryCode: "IE",
			PostalCode:  "D12",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 13",
			AdminName1:  "Leinster",
			Lat:         "53.3946",
			Long:        "-6.1495",
			CountryCode: "IE",
			PostalCode:  "D13",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 14",
			AdminName1:  "Leinster",
			Lat:         "53.296",
			Long:        "-6.2593",
			CountryCode: "IE",
			PostalCode:  "D14",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 15",
			AdminName1:  "Leinster",
			Lat:         "53.3832",
			Long:        "-6.4165",
			CountryCode: "IE",
			PostalCode:  "D15",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 16",
			AdminName1:  "Leinster",
			Lat:         "53.3419",
			Long:        "-6.279",
			CountryCode: "IE",
			PostalCode:  "D16",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 17",
			AdminName1:  "Leinster",
			Lat:         "53.4006",
			Long:        "-6.2058",
			CountryCode: "IE",
			PostalCode:  "D17",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 18",
			AdminName1:  "Leinster",
			Lat:         "53.2469",
			Long:        "-6.1774",
			CountryCode: "IE",
			PostalCode:  "D18",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 20",
			AdminName1:  "Leinster",
			Lat:         "53.3518",
			Long:        "-6.3693",
			CountryCode: "IE",
			PostalCode:  "D20",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 22",
			AdminName1:  "Leinster",
			Lat:         "53.3275",
			Long:        "-6.4006",
			CountryCode: "IE",
			PostalCode:  "D22",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 24",
			AdminName1:  "Leinster",
			Lat:         "53.2851",
			Long:        "-6.3713",
			CountryCode: "IE",
			PostalCode:  "D24",
			AdminCode1:  "L",
		},
		{
			PlaceName:   "Dublin 6W",
			AdminName1:  "Leinster",
			Lat:         "53.3087",
			Long:        "-6.3012",
			CountryCode: "IE",
			PostalCode:  "D6W",
			AdminCode1:  "L",
		},
	}

	if !cmp.Equal(want, got) {
		t.Errorf(cmp.Diff(want, got))
	}
}
