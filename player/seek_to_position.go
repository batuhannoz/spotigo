package spotify

import "net/http"

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/seek-to-position-in-currently-playing-track
//
// Seeks to the given position in the user’s currently playing track.
func SeekToPosition(token string, deviceID string, positionms string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, SeekToPositionURL, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("position_ms", positionms)
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
