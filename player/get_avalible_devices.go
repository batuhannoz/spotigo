package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type GetAvailableDevicesResponse struct {
	Devices []struct {
		Id               string `json:"id"`
		IsActive         bool   `json:"is_active"`
		IsPrivateSession bool   `json:"is_private_session"`
		IsRestricted     bool   `json:"is_restricted"`
		Name             string `json:"name"`
		Type             string `json:"type"`
		VolumePercent    int    `json:"volume_percent"`
	} `json:"devices"`
}

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
