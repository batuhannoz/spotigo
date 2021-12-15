package spotify

import (
	"bytes"
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func GetCurrentUsersProfile(input GetCurrentUsersProfileInput) (GetCurrentUsersProfileResponse, error) {
	var output spotify.OutputToSpotify
	var response GetCurrentUsersProfileResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me"

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func FollowPlaylist(input FollowPlaylistInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	var request FollowPlaylistRequest
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "playlists/" + input.PlaylistId + "/followers"

	request.Public = input.Public
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return response, err
	}
	output.Body = bytes.NewBuffer(jsonBody)

	err = spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetUsersProfile(input GetUsersProfileInput) (GetUsersProfileResponse, error) {
	var output spotify.OutputToSpotify
	var response GetUsersProfileResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "user/" + input.UserId

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetUsersTopItems(input GetUsersTopItemsInput) (GetUsersTopItemsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetUsersTopItemsResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/top/" + input.TopItemsType

	output.Query.QueryName[0], output.Query.QueryValue[0] = "limit", strconv.Itoa(input.Limit)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "offset", strconv.Itoa(input.Offset)
	output.Query.QueryName[2], output.Query.QueryValue[2] = "time_range", input.PlaylistId
	output.Query.QueryNumber = 3

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func UnfollowPlaylist(input UnfollowPlaylistInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodDelete
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "playlists/" + input.PlaylistId + "/followers"

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetFollowedArtists(input GetFollowedArtistsInput) (GetFollowedArtistsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetFollowedArtistsResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/following"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "type", input.FollowedType
	output.Query.QueryName[1], output.Query.QueryValue[1] = "after", strconv.Itoa(input.After)
	output.Query.QueryNumber = 1

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}
