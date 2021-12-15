package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func GetAlbum(input GetAlbumInput) (GetAlbumResponse, error) {
	var output spotify.OutputToSpotify
	var response GetAlbumResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "albums/" + input.Id

	output.Query.QueryName[0], output.Query.QueryValue[0] = "market", input.Market
	output.Query.QueryNumber = 1

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetAlbumTracks(input GetAlbumTracksInput) (GetAlbumTracksResponse, error) {
	var output spotify.OutputToSpotify
	var response GetAlbumTracksResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "albums/" + input.Id + "/track"

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

func GetNewReleases(input GetNewReleasesInput) (GetNewReleasesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetNewReleasesResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "browse/new-releases"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "limit", strconv.Itoa(input.Limit)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "country", input.Country
	output.Query.QueryName[2], output.Query.QueryValue[2] = "offset", strconv.Itoa(input.Offset)
	output.Query.QueryNumber = 3

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetSavedAlbums(input GetSavedAlbumsInput) (GetSavedAlbumsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetSavedAlbumsResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/albums"

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
