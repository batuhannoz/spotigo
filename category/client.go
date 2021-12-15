package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func SeveralBrowseCategories(input SeveralBrowseCategoriesInput) (SeveralBrowseCategoriesResponse, error) {
	var output spotify.OutputToSpotify
	var response SeveralBrowseCategoriesResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "browse/categories"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "country", input.Country
	output.Query.QueryName[1], output.Query.QueryValue[1] = "limit", strconv.Itoa(input.Limit)
	output.Query.QueryName[2], output.Query.QueryValue[2] = "locale", input.Locale
	output.Query.QueryName[3], output.Query.QueryValue[3] = "offset", strconv.Itoa(input.Offset)
	output.Query.QueryNumber = 4

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func SingleBrowseCategories(input SingleBrowseCategoriesInput) (SingleCategoriesResponse, error) {
	var output spotify.OutputToSpotify
	var response SingleCategoriesResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "browse/categories/" + input.CategoryId

	output.Query.QueryName[0], output.Query.QueryValue[0] = "country", input.Country
	output.Query.QueryName[1], output.Query.QueryValue[1] = "locale", input.Locale
	output.Query.QueryNumber = 2
	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}
