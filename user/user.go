package user

import (
	"encoding/json"
	"errors"

	"github.com/luladjiev/hnews/request"
)

// Data - Hacker-news User structure
type Data struct {
	ID        string `json:"id"`
	Delay     int    `json:"delay"`
	Created   int    `json:"created"`
	Karma     int    `json:"karma"`
	About     string `json:"about"`
	Submitted []int  `json:"submitted"`
}

// Get a single item
func Get(id string) (Data, error) {
	var result Data

	url := request.GetUserURL(id)

	response, err := request.Get(url)

	if err != nil {
		return result, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return result, err
	}

	if result.ID == "" {
		return result, errors.New("invalid response data")
	}

	return result, nil
}
