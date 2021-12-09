package spotify

import (
	"fmt"
	"net/http"
	"strconv"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/set-volume-for-users-playback
//
// Set the volume for the user’s current playback device.
func SetPlaybackVolume(token string, deviceID string, volume int) error {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodPut, SetPlaybackVolumeURL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", token)
	q := req.URL.Query()
	q.Add("device_id", deviceID)
	q.Add("volume_percent", strconv.Itoa(volume))
	req.URL.RawQuery = q.Encode()
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	return err
}
