package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

func AvailableGenreSeeds(input AvailableGenreSeedsInput) (GenreSeedsResponse, error) {
	var output spotify.OutputToSpotify
	var response GenreSeedsResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "recommendations/available-genre-seeds"

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}
