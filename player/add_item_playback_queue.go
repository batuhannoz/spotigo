package spotify

import "net/http"

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/add-to-queue
//
// Add an item to the end of the user's current playback queue.
func AddItemPlaybackQueue(token string, deviceID string, trackUri string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, AddİtemPlaybackQueueURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("uri", trackUri)
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
