package main

import (
	"fmt"
	"net/http"
	"net/url"
)

// TrelloManager automates task creation based on system events
type TrelloManager struct {
	APIKey    string
	Token     string
	ListID    string
}

func (tm *TrelloManager) CreateIncidentCard(title string, desc string) error {
	baseURL := "https://api.trello.com/1/cards"
	params := url.Values{}
	params.Add("key", tm.APIKey)
	params.Add("token", tm.Token)
	params.Add("idList", tm.ListID)
	params.Add("name", title)
	params.Add("desc", desc)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	resp, err := http.Post(fullURL, "application/json", nil)
	
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("failed to sync incident to Trello")
	}
	
	fmt.Println("Incident card created in Trello successfully.")
	return nil
}

func main() {
	manager := TrelloManager{
		APIKey: "YOUR_KEY", 
		Token:  "YOUR_TOKEN", 
		ListID: "INCIDENT_LIST_ID",
	}
	
	// Called by your monitoring system when a system fails
	manager.CreateIncidentCard("CRITICAL: Database Latency", "P99 latency > 2s on AuthService")
}
