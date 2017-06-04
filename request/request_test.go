package request

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetWrongResponseStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))

	defer ts.Close()

	_, err := Get(ts.URL)

	if err == nil {
		t.Errorf("Get() didn't check server status code")
	}
}

func TestSetAPIURL(t *testing.T) {
	SetAPIURL("test.com")

	itemURL := GetItemURL(1)
	userURL := GetUserURL("luladjiev")

	allURLs := [2]string{itemURL, userURL}

	for _, url := range allURLs {
		if strings.HasPrefix(url, "test.com") != true {
			t.Errorf("SetAPIURL() didn't set URL correctly")
		}
	}

	if itemURL != "test.com/item/1.json" {
		t.Errorf("GetItemURL() didn't return correct URL")
	}

	if userURL != "test.com/user/luladjiev.json" {
		t.Errorf("GetUserURL() didn't return correct URL")
	}
}
