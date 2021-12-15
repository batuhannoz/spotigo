package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func AvailableGenreSeeds(input AvailableGenreSeedsInput) (GenreSeedsResponse, error) {
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

	err = mapstructure.Decode(res, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
