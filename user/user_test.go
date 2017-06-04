package user

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/luladjiev/hnews/request"
)

func TestGetWrongResponseData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "null")
	}))

	defer ts.Close()

	request.SetAPIURL(ts.URL)

	_, err := Get("luladjiev")

	if err == nil {
		t.Errorf("Get() didn't check server response")
	}
}

func TestGetCorrectParsing(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "{\"created\": 1468307419, \"id\": \"luladjiev\", \"karma\": 1, \"submitted\": [13349223]}")
	}))

	defer ts.Close()

	request.SetAPIURL(ts.URL)

	data, err := Get("luladjiev")

	if err != nil {
		t.Errorf("Get() got an error parsing server response")
	}

	if data.ID != "luladjiev" {
		t.Errorf("Get() couldn't parse ID")
	}

	if data.Created != 1468307419 {
		t.Errorf("Get() couldn't parse Created")
	}

	if data.Karma != 1 {
		t.Errorf("Get() couldn't parse Karma")
	}

	if len(data.Submitted) != 1 || data.Submitted[0] != 13349223 {
		t.Errorf("Get() couldn't parse Submitted")
	}
}
