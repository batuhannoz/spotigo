package spotify

import (
	"bytes"
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

func GetCurrentUsersProfile(input spotify.UserInfo) (GetCurrentUsersProfileResponse, error) {
	var output spotify.OutputToSpotify
	var response GetCurrentUsersProfileResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me"

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

func FollowPlaylist(input spotify.UserInfo) (string, error) {
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

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func GetUsersProfile(input spotify.UserInfo) (GetUsersProfileResponse, error) {
	var output spotify.OutputToSpotify
	var response GetUsersProfileResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "user/" + input.UserId

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

func GetUsersTopItems(input spotify.UserInfo) (GetUsersTopItemsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetUsersTopItemsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/top/" + input.TopItemsType

	output.Query.QueryName[0], output.Query.QueryValue[0] = "limit", string(input.Limit)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "offset", string(input.Offset)
	output.Query.QueryName[2], output.Query.QueryValue[2] = "time_range", input.PlaylistId
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

func UnfollowPlaylist(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodDelete
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "playlists/" + input.PlaylistId + "/followers"

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func GetFollowedArtists(input spotify.UserInfo) (GetFollowedArtistsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetFollowedArtistsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/following"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "type", input.FollowedType
	output.Query.QueryName[1], output.Query.QueryValue[1] = "after", string(input.After)
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
