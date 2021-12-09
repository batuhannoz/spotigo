package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/get-a-users-available-devices
//
// Get information about a user’s available devices
func GetAvailableDevices(spotifyToken string) (GetAvailableDevicesResponse, error) {
	response := GetAvailableDevicesResponse{}
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, GetAvailableDevicesURL, nil)

	if err != nil {
		return response, err
	}

	request.Header.Add("Authorization", spotifyToken)

	res, err := client.Do(request)
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
