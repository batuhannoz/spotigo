package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func GetUsersSavedShows(input GetUsersSavedShowsInput) (GetUsersSavedShowsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetUsersSavedShowsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/shows"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "limit", strconv.Itoa(input.Limit)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "offset", strconv.Itoa(input.Offset)
	output.Query.QueryNumber = 2

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

func GetShows(input GetShowsInput) (GetShowsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetShowsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "shows/" + input.Id

	output.Query.QueryName[0], output.Query.QueryValue[0] = "market", input.Market
	output.Query.QueryNumber = 1

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

func GetShowEpisodes(input GetShowEpisodesInput) (GetShowEpisodesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetShowEpisodesResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "shows/" + input.Id + "/episodes"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "limit", strconv.Itoa(input.Limit)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "market", input.Market
	output.Query.QueryName[2], output.Query.QueryValue[2] = "offset", strconv.Itoa(input.Offset)
	output.Query.QueryNumber = 3

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
