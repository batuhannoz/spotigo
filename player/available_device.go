package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
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

func GetAvailableDevices(input spotify.UserInfo) (GetAvailableDevicesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetAvailableDevicesResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = 200
	output.Url = spotify.BaseUrl + "me/player/devices"

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
