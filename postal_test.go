package geonames_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/qba73/geonames"
)

func TestGetPostalCodes_ReturnsSingleValueOnValidInput(t *testing.T) {
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
	got, err := client.GetPostCode(context.Background(), place, country)
	if err != nil {
		t.Fatal(err)
	}

	want := []geonames.PostalCode{
		{
			PlaceName:  "Castlebar",
			AdminName1: "Connacht",
			Position: geonames.Position{
				Lat: 53.85,
				Lng: -9.3,
			},
			CountryCode: "IE",
			PostalCode:  "F23",
			AdminCode1:  "C",
		},
	}

	if len(got) != 1 {
		t.Fatalf("want one postal code, got %v", got)
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreFields(geonames.PostalCode{}, "Position")) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetPostalCodes_ReturnsMultipleValuesOnValidInput(t *testing.T) {
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
	got, err := client.GetPostCode(context.Background(), place, country)
	if err != nil {
		t.Fatal(err)
	}

	want := []geonames.PostalCode{
		{
			PlaceName:  "Dublin 1",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D01",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 2",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D02",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 3",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D03",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 4",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D04",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 5",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D05",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 6",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D06",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 7",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D07",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 8",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D08",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 9",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D09",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 10",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D10",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 11",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D11",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 12",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D12",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 13",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D13",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 14",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D14",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 15",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D15",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 16",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D16",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 17",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D17",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 18",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D18",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 20",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D20",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 22",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D22",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 24",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D24",
			AdminCode1:  "L",
		},
		{
			PlaceName:  "Dublin 6W",
			AdminName1: "Leinster",
			Position: geonames.Position{
				Lat: 53.354,
				Lng: -6.2545,
			},
			CountryCode: "IE",
			PostalCode:  "D6W",
			AdminCode1:  "L",
		},
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreFields(geonames.PostalCode{}, "Position")) {
		t.Errorf(cmp.Diff(want, got))
	}
}
