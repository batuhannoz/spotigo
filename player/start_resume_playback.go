package spotify

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type StartResumePlaybackRequest struct {
	ContextUri string `json:"context_uri,omitempty"`
	Offset     struct {
		Position int `json:"position,omitempty"`
	} `json:"offset,omitempty"`
	PositionMs int `json:"position_ms,omitempty"`
}

// Start a new context or resume current playback on the user's active device.
func StartResumePlayback(token string, deviceID string, contextUri string, position int, positionms int) error {
	client := http.Client{}
	var reqBody StartResumePlaybackRequest
	reqBody.ContextUri = contextUri
	reqBody.Offset.Position = position
	reqBody.PositionMs = positionms

	jsonBody, err := json.Marshal(reqBody)
	req, err := http.NewRequest(http.MethodPut, StartResumePlaybackURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("device_id", deviceID)
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", token)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
