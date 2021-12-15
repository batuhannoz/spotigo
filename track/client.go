package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func GetUsersSavedTracks(input GetUsersSavedTracksInput) (GetUsersSavedTracksResponse, error) {
	var output spotify.OutputToSpotify
	var response GetUsersSavedTracksResponse
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

func GetTrack(input GetTrackInput) (GetTracksResponse, error) {
	var output spotify.OutputToSpotify
	var response GetTracksResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "track/" + input.Id

	output.Query.QueryName[0], output.Query.QueryValue[0] = "market", input.Market
	output.Query.QueryNumber = 1

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetTracksAudioAnalysis(input GetTracksAudioAnalysisInput) (GetTracksAudioAnalysisResponse, error) {
	var output spotify.OutputToSpotify
	var response GetTracksAudioAnalysisResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "audio-analysis/" + input.Id

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}
