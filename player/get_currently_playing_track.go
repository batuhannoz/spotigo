package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get the object currently being played on the user's Spotify account.
func GetCurrentlyPlayingTrack(token string, market string) (GetCurrentlyPlayingTrackResponse, error) {
	response := GetCurrentlyPlayingTrackResponse{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetCurrentlyPlayingTrackURL, nil)

	if err != nil {
		fmt.Println(err)
		return response, err
	}
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	bodyString := string(bodyBytes)
	err = json.Unmarshal([]byte(bodyString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
