package spotify

import (
	"net/http"
)

// Pause playback on the user's account.
func PausePlayback(token string, deviceID string) error {

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, PausePlaybackURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("device_id", deviceID)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}