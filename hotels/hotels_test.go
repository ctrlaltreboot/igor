package hotels

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCheapest(t *testing.T) {
	// Create fake hotels API endpoint that returns response from fixture.
	fakeResponse, err := ioutil.ReadFile("fixtures/sample-response.json")
	if err != nil {
		t.Fatalf("error reading fixture file: %v", err)
	}
	hotelsTS := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(fakeResponse)
	}))
	defer hotelsTS.Close()

	handler := &CheapestHandler{HotelsAPIEndpoint: hotelsTS.URL}
	ts := httptest.NewServer(handler)
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		t.Fatalf("error creating http.Request: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error making HTTP request: %v", err)
	}
	defer res.Body.Close()

	if want, got := http.StatusOK, res.StatusCode; want != got {
		t.Errorf("expected status code to be %v but got %v", want, got)
	}

	type property struct {
		ID string `json:"id"`
	}

	var decoded struct {
		Properties []property `json:"properties"`
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("error reading response body: %v", err)
	}

	if err := json.Unmarshal(body, &decoded); err != nil {
		t.Fatalf("error decoding response body as JSON: %v, body:\n%s", err, body)
	}

	// TODO Tests from this point should verify that we got the expected N
	// cheapest hotels.
	if want, got := 6, len(decoded.Properties); want != got {
		t.Errorf("expected response to contain %d hotels but got %d, body:\n%s", want, got, body)
	}
}
