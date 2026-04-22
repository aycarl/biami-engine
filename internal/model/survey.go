package model

import "encoding/json"

// SurveyResponse stores channel metadata and dynamic survey answers.
type SurveyResponse struct {
	SessionID string          `json:"session_id"`
	Mobile    string          `json:"mobile"`
	Channel   string          `json:"channel"`
	Payload   json.RawMessage `json:"payload"`
}

// MobileHookRequest is the raw payload received from a mobile gateway.
type MobileHookRequest struct {
	From    string `json:"from"`
	Text    string `json:"text"`
	Session string `json:"session_id"`
}
