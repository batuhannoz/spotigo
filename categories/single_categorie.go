package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type SingleCategoriesResponse struct {
	Href  string `json:"href"`
	Icons []struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"icons"`
	Id   string `json:"id"`
	Name string `json:"name"`
}

func SingleBrowseCategories(input spotify.UserInfo) (SingleCategoriesResponse, error) {
	var output spotify.OutputToSpotify
	var response SingleCategoriesResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = 200
	output.Url = spotify.BaseUrl + "browse/categories/" + input.CategoryId

	output.Query.QueryName[0], output.Query.QueryValue[0] = "country", input.Country
	output.Query.QueryName[1], output.Query.QueryValue[1] = "locale", input.Locale

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
