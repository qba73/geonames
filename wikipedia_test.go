package geonames_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/qba73/geonames"
)

func TestGetPlace_RetrievesSingleGeoNameOnValidInput(t *testing.T) {
	t.Parallel()
	testFile := "testdata/response-geoname-wikipedia-single.json"
	wantReqURI := "/wikipediaSearchJSON?q=Castlebar&title=Castlebar&countryCode=IE&maxRows=1&username=DummyUser"
	ts := newTestServer(testFile, wantReqURI, t)
	defer ts.Close()

	client, err := geonames.NewClient(
		geonames.WithUserName("DummyUser"),
		geonames.WithBaseURL(ts.URL),
	)
	if err != nil {
		t.Fatal(err)
	}

	name := "Castlebar"
	country := "IE"
	resultLimit := 1

	got, err := client.GetPlace(name, country, resultLimit)
	if err != nil {
		t.Fatal(err)
	}

	want := []geonames.Geoname{
		{
			Summary:   "Castlebar is the county town of County Mayo, Ireland. It is in the middle of the county and is its largest town by population. A campus of Galway-Mayo Institute of Technology and the Country Life section of the National Museum of Ireland are two important local amenities (...)",
			Elevation: 41,
			GeoNameID: 2965654,
			Position: geonames.Position{
				Lat:  53.8608,
				Long: -9.2988,
			},
			CountryCode: "IE",
			Rank:        100,
			Language:    "en",
			Title:       "Castlebar",
			URL:         "en.wikipedia.org/wiki/Castlebar",
		},
	}

	if !cmp.Equal(want, got, cmpopts.IgnoreFields(geonames.Geoname{}, "Summary", "Position")) {
		t.Errorf(cmp.Diff(want, got))
	}
}
