package geonames_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/qba73/geonames"
)

// ts is a helper func that creates a test server with embedded URI validation.
var testServer = func(reader io.Reader, wantURI string, t *testing.T) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		gotReqURI := r.RequestURI
		verifyURIs(wantURI, gotReqURI, t)
		_, err := io.Copy(rw, reader)
		if err != nil {
			t.Fatal(err)
		}
	}))
	return ts
}

func TestGetElevationReturnsDataOnValidInput(t *testing.T) {
	t.Parallel()

	resp := new(bytes.Buffer)
	resp.Read([]byte(`23`))

	ts := testServer(&resp)

}

func TestGetElevationValidInput(t *testing.T) {
	t.Parallel()

	wantReqURI := "/srtm1JSON?lat=50&lng=50&username=DummyUser"
	ts := newTestServer(testFile, wantReqURI, t)
	defer ts.Close()

	client, err := geonames.NewClient("DummyUser", geonames.WithBaseURL(ts.URL))
	if err != nil {
		t.Fatal(err)
	}

	client.GetElevation()
}
