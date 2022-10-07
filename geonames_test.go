package geonames_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/geonames"
)

// newTestServer is a helper func that creates a test server with embedded URI validation.
func newTestServer(testFile, wantURI string, t *testing.T) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		gotReqURI := r.RequestURI
		verifyURIs(wantURI, gotReqURI, t)
		f, err := os.Open(testFile)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		_, err = io.Copy(rw, f)
		if err != nil {
			t.Fatal(err)
		}
	}))
	return ts
}

// verifyURIs is a test helper function that verifies if provided URIs are the same.
func verifyURIs(wanturi, goturi string, t *testing.T) {
	wantU, err := url.Parse(wanturi)
	if err != nil {
		t.Fatalf("error parsing URL %q, %v", wanturi, err)
	}
	gotU, err := url.Parse(goturi)
	if err != nil {
		t.Fatalf("error parsing URL %q, %v", wanturi, err)
	}
	// Verify if paths of both URIs are the same.
	if wantU.Path != gotU.Path {
		t.Fatalf("want %q, got %q", wantU.Path, gotU.Path)
	}

	wantQuery, err := url.ParseQuery(wantU.RawQuery)
	if err != nil {
		t.Fatal(err)
	}
	gotQuery, err := url.ParseQuery(gotU.RawQuery)
	if err != nil {
		t.Fatal(err)
	}

	// Verify if query parameters match in both, got and want URIs.
	if !cmp.Equal(wantQuery, gotQuery) {
		t.Fatalf("URIs are not equal, \n%s\n", cmp.Diff(wantQuery, gotQuery))
	}
}

func TestWikipediaResolvesGeoNameOnValidInput(t *testing.T) {
	t.Parallel()

	testFile := "testdata/response-geoname-wikipedia-single.json"
	wantReqURI := "/wikipediaSearchJSON?q=Castlebar&title=Castlebar&countryCode=IE&maxRows=1&username=DummyUser"
	ts := newTestServer(testFile, wantReqURI, t)
	defer ts.Close()

	name, country := "Castlebar", "IE"
	client, err := geonames.NewClient(
		geonames.WithUserName("DummyUser"),
		geonames.WithBaseURL(ts.URL),
	)
	if err != nil {
		t.Fatal(err)
	}

	client.GetPlace(name, country, 1)
}
