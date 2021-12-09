package spotify

import "net/http"

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/toggle-shuffle-for-users-playback
//
// Toggle shuffle on or off for user’s playback.
//
// Can be one of the states: true, false
func TogglePlaybackShuffle(token string, deviceID string, state string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, TogglePlaybackShuffleURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("state", state)
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
