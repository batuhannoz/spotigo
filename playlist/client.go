package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

func GetPlaylists(input GetPlaylistsInput) (GetPlaylistsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetPlaylistsResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "playlists/" + input.PlaylistId

	output.Query.QueryName[0], output.Query.QueryValue[0] = "additional_types", input.AdditionalTypes
	output.Query.QueryName[1], output.Query.QueryValue[1] = "fields", input.Fields
	output.Query.QueryName[2], output.Query.QueryValue[2] = "market", input.Market
	output.Query.QueryNumber = 3

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}
