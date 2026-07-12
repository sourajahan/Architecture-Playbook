package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// SlackBlock represents a professional, actionable alert structure
type SlackBlock struct {
	Type string `json:"type"`
	Text struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"text,omitempty"`
}

type SlackPayload struct {
	Blocks []SlackBlock `json:"blocks"`
}

// AlertManager processes system events and formats them for the team
type AlertManager struct {
	WebhookURL string
}

func (am *AlertManager) NotifyCritical(service, traceID, message string) error {
	// Constructing Block Kit payload for immediate actionability
	payload := SlackPayload{
		Blocks: []SlackBlock{
			{Type: "section", Text: struct {
				Type string `json:"type"`
				Text string `json:"text"`
			}{Type: "mrkdwn", Text: fmt.Sprintf("*CRITICAL FAILURE in %s*", service)}},
			{Type: "section", Text: struct {
				Type string `json:"type"`
				Text string `json:"text"`
			}{Type: "mrkdwn", Text: fmt.Sprintf("Trace: `%s`\nError: %s", traceID, message)}},
		},
	}

	data, _ := json.Marshal(payload)
	resp, err := http.Post(am.WebhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil || resp.StatusCode != 200 {
		return fmt.Errorf("failed to notify slack")
	}
	return nil
}

func main() {
	notifier := AlertManager{WebhookURL: "https://hooks.slack.com/services/..."}
	notifier.NotifyCritical("AuthService", "tr-12345", "Database connection pool exhausted")
}
