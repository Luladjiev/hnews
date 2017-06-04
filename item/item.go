package item

import (
	"encoding/json"

	"github.com/luladjiev/hnews/request"
)

// Data - Hacker-news Item structure
type Data struct {
	ID          int    `json:"id"`
	Deleted     bool   `json:"deleted"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Time        int    `json:"time"`
	Text        string `json:"text"`
	Dead        bool   `json:"dead"`
	Parent      int    `json:"parent"`
	Poll        int    `json:"poll"`
	Kids        []int  `json:"kids"`
	URL         string `json:"url"`
	Score       int    `json:"score"`
	Title       string `json:"title"`
	Parts       []int  `json:"parts"`
	Descendants int    `json:"descendants"`
}

// Get a single item
func Get(id int) (Data, error) {
	var result Data

	url := request.GetItemURL(id)

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
