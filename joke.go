package concurrencystuff

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JokeService contains joke implementations
type JokeService interface {
	GetJoke() (interface{}, error)
	FuckOffAsshole(name string) (interface{}, error)
}

// Svc contains service url
type svc struct {
	yomamaURL string
	foaasURL  string
}

// New initializes new service
func New(yomama, foaas string) JokeService {
	return &svc{
		yomamaURL: yomama,
		foaasURL:  foaas,
	}
}

// GetJoke gets a joke from the yo mama api
func (s *svc) GetJoke() (interface{}, error) {
	resp, err := http.Get(s.yomamaURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	j := struct {
		Joke string `json:"joke,omitempty"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return nil, err
	}

	return j, nil
}

// FuckOff gets content from FOAAS (Fuck off as a service)
func (s *svc) FuckOffAsshole(name string) (interface{}, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/%s", s.foaasURL, name)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	m := struct {
		Message  string `json:"message,omitempty"`
		Subtitle string `json:"subtitle,omitempty"`
	}{}

	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
