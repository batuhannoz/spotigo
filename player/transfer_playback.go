package spotify

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/transfer-a-users-playback
//
// Transfer playback to a new device and determine if it should start playing.
func TransferPLayback(spotifyToken string, deviceIDs []string, play bool) error {
	client := &http.Client{}
	var requestBody TransferPlaybackRequest
	requestBody.DeviceIds = deviceIDs
	requestBody.Play = play
	jsonBody, err := json.Marshal(requestBody)
	req, err := http.NewRequest(http.MethodPut, TransferPlaybackURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", spotifyToken)
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return err
}
