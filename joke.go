package concurrencystuff

import (
	"encoding/json"
	"net/http"
)

// GetJoke gets a joke from the yo mama api
func GetJoke() (interface{}, error) {
	req, err := http.Get("https://api.yomomma.info/")
	if err != nil {
		return nil, err
	}

	j := struct {
		Joke string `json:"joke,omitempty"`
	}{}

	err = json.NewDecoder(req.Body).Decode(&j)
	if err != nil {
		return nil, err
	}

	return j, nil
}
