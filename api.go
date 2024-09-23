package api

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"net/http"
)

// FetchJobs fetches the jobs from the API and returns a slice of jobs as a string
func FetchJobs() (string, error) {
	url := "https://cat-fact.herokuapp.com/facts"

	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch jobs: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	var facts Facts
	err = json.Unmarshal(body, &facts)
	if err != nil {
		return "", fmt.Errorf("failed to parse jobs: %v", err)
	}

	factsString, err := json.MarshalIndent(facts, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to convert jobs to string: %v", err)
	}

	return string(factsString), nil
}

type Facts []struct {
	Status struct {
		Verified  bool `json:"verified"`
		SentCount int  `json:"sentCount"`
	} `json:"status,omitempty"`
	ID        string    `json:"_id"`
	User      string    `json:"user"`
	Text      string    `json:"text"`
	V         int       `json:"__v"`
	Source    string    `json:"source"`
	UpdatedAt time.Time `json:"updatedAt"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	Deleted   bool      `json:"deleted"`
	Used      bool      `json:"used"`
	Status0   struct {
		Verified  bool   `json:"verified"`
		Feedback  string `json:"feedback"`
		SentCount int    `json:"sentCount"`
	} `json:"status,omitempty"`
}

// func main() {
// 	fmt.Println(Hello())
// }go
