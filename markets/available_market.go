package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type AvailableMarketsResponse struct {
	Markets []string `json:"markets"`
}

func GeneralSpotifyFunc(input spotify.UserInfo) (AvailableMarketsResponse, error) {
	var output spotify.OutputToSpotify
	var response AvailableMarketsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "markets/"

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
