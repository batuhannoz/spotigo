package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type GenreSeedsResponse struct {
	Genres []string `json:"genres"`
}

func AvailableGenreSeeds(input spotify.UserInfo) (GenreSeedsResponse, error) {
	var output spotify.OutputToSpotify
	var response GenreSeedsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "recommendations/available-genre-seeds"

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
