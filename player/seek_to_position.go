package spotify

import "net/http"

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
