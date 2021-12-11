package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type SeveralBrowseCategoriesResponse struct {
	Categories struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"categories"`
}

func SeveralBrowseCategories(input spotify.UserInfo) (SeveralBrowseCategoriesResponse, error) {
	var output spotify.OutputToSpotify
	var response SeveralBrowseCategoriesResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = 200
	output.Url = spotify.BaseUrl + "browse/categories"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "country", input.Country
	output.Query.QueryName[1], output.Query.QueryValue[1] = "limit", string(input.Limit)
	output.Query.QueryName[2], output.Query.QueryValue[2] = "locale", input.Locale
	output.Query.QueryName[3], output.Query.QueryValue[3] = "offset", string(input.Offset)

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
