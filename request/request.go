package request

import (
	"net/http"
	"strconv"
)

var (
	apiURL  = "https://hacker-news.firebaseio.com/v0"
	itemURL = apiURL + "/item/"
	userURL = apiURL + "/user/"
)

// SetAPIURL set's the base and rest of the URLs
func SetAPIURL(url string) {
	apiURL = url
	itemURL = apiURL + "/item/"
	userURL = apiURL + "/user/"
}

// GetItemURL returns a request ready url for a specific item
func GetItemURL(id int) string {
	return itemURL + strconv.Itoa(id) + ".json"
}

// GetUserURL returns a request ready url for a specific user
func GetUserURL(id string) string {
	return userURL + id + ".json"
}

// Get data
func Get(url string) (*http.Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return response, err
	}

	return response, nil
}
