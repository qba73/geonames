package geonames_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/qba73/geonames"
)

// newElevationtestServer creates a test server with embedded URI validation.
func newElevationTestServer(data []byte, wantURI string, t *testing.T) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		gotReqURI := r.RequestURI
		verifyURIs(wantURI, gotReqURI, t)
		_, err := io.Copy(rw, bytes.NewBuffer(data))
		if err != nil {
			t.Fatal(err)
		}
	}))
	return ts
}

func TestGetElevationSRTM1ReturnsDataOnValidInput(t *testing.T) {
	t.Parallel()

	lat, lng := 50.0, 50.0
	wantReqURI := fmt.Sprintf("/srtm1JSON?lat=%.3f&lng=%.3f&username=DummyUser", lat, lng)

	ts := newElevationTestServer(srtm1, wantReqURI, t)
	defer ts.Close()

	client := geonames.NewClient("DummyUser")
	client.BaseURL = ts.URL

	want := geonames.Elevation{
		Type:  "srtm1",
		Lat:   54.166,
		Lng:   -6.083,
		Value: 375,
	}

	got, err := client.GetElevationSRTM1(context.Background(), lat, lng)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetElevationSRTM3ReturnsDataOnValidInput(t *testing.T) {
	t.Parallel()

	lat, lng := 55.166, -6.088
	wantReqURI := fmt.Sprintf("/srtm3JSON?lat=%.3f&lng=%.3f&username=DummyUser", lat, lng)

	ts := newElevationTestServer(srtm3, wantReqURI, t)
	defer ts.Close()

	client := geonames.NewClient("DummyUser")
	client.BaseURL = ts.URL

	want := geonames.Elevation{
		Type:  "srtm3",
		Lat:   55.166,
		Lng:   -6.088,
		Value: 263,
	}

	got, err := client.GetElevationSRTM3(context.Background(), lat, lng)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetElevationAstergdemReturnsDataOnValidInput(t *testing.T) {
	t.Parallel()

	lat, lng := 50.01, 10.20
	wantReqURI := fmt.Sprintf("/astergdemJSON?lat=%.3f&lng=%.3f&username=DummyUser", lat, lng)

	ts := newElevationTestServer(astergdem, wantReqURI, t)
	defer ts.Close()

	client := geonames.NewClient("DummyUser")
	client.BaseURL = ts.URL

	want := geonames.Elevation{
		Type:  "astergdem",
		Lat:   50.010,
		Lng:   10.200,
		Value: 206,
	}

	got, err := client.GetElevationAstergdem(context.Background(), lat, lng)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetElevationGTOPO30ReturnsDataOnValidInput(t *testing.T) {
	t.Parallel()

	lat, lng := 47.01, 10.2
	wantReqURI := fmt.Sprintf("/gtopo30JSON?lat=%.3f&lng=%.3f&username=DummyUser", lat, lng)

	ts := newElevationTestServer(gtopo30, wantReqURI, t)
	defer ts.Close()

	client := geonames.NewClient("DummyUser")
	client.BaseURL = ts.URL

	want := geonames.Elevation{
		Type:  "gtopo30",
		Lat:   47.01,
		Lng:   10.20,
		Value: 2632,
	}

	got, err := client.GetElevationGTOPO30(context.Background(), lat, lng)
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

var (
	srtm1     = []byte(`{"srtm1":375,"lng":-6.083,"lat":54.166}`)
	srtm3     = []byte(`{"srtm3":263,"lng":-6.088,"lat":55.166}`)
	astergdem = []byte(`{"lng":10.2,"astergdem":206,"lat":50.01}`)
	gtopo30   = []byte(`{"lng":10.2,"gtopo30":2632,"lat":47.01}`)
)
