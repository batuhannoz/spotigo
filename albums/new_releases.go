package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type GetNewReleasesResponse struct {
	Albums struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"albums"`
}

func GetNewReleases(input spotify.UserInfo) (GetNewReleasesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetNewReleasesResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "browse/new-releases"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "limit", string(input.Limit)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "country", input.Country
	output.Query.QueryName[2], output.Query.QueryValue[2] = "offset", string(input.Offset)
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
