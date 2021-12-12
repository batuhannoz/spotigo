package spotify

import (
	"bytes"
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type FollowPlaylistRequest struct {
	Public bool `json:"public"`
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
