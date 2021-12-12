package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type GetArtistsAlbumsResponse struct {
	Href  string `json:"href"`
	Items []struct {
	} `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

func SingleBrowseCategories(input spotify.UserInfo) (GetArtistsAlbumsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetArtistsAlbumsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "artists/" + input.Id + "/albums"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "include_groups", input.IncludeGroups
	output.Query.QueryName[1], output.Query.QueryValue[1] = "limit", string(input.Limit)
	output.Query.QueryName[2], output.Query.QueryValue[2] = "market", input.Market
	output.Query.QueryName[3], output.Query.QueryValue[3] = "offset", string(input.Offset)
	output.Query.QueryNumber = 4

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
