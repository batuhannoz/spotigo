package spotify

import "net/http"

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/set-repeat-mode-on-users-playback
//
// Set the repeat mode for the user's playback. Options are repeat-track, repeat-context, and off.
//
// Can be one of the states: track, context, off
func SetRepeatMode(token string, deviceID string, state string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, SetRepeatModeURL, nil)
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
