package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func GetEpisodes(input GetEpisodesInput) (GetEpisodesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetEpisodesResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "episodes/" + input.Id

	output.Query.QueryName[0], output.Query.QueryValue[0] = "market", input.Market
	output.Query.QueryNumber = 1

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetUsersSavedEpisodes(input GetUsersSavedEpisodesInput) (GetUsersSavedEpisodesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetUsersSavedEpisodesResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/episodes"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "limit", strconv.Itoa(input.Limit)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "market", input.Market
	output.Query.QueryName[2], output.Query.QueryValue[2] = "offset", strconv.Itoa(input.Offset)
	output.Query.QueryNumber = 3

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}
