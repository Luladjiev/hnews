package user

import (
	"encoding/json"

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

	defer response.Body.Close()

	if err != nil {
		return result, err
	}

	err = json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}
