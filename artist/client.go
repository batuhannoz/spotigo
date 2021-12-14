package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func GetArtist(input GetArtistInput) (GetArtistResponse, error) {
	var output spotify.OutputToSpotify
	var response GetArtistResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "artist/" + input.Id

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

func SingleBrowseCategories(input SingleBrowseCategoriesInput) (GetArtistsAlbumsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetArtistsAlbumsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "artists/" + input.Id + "/albums"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "include_groups", input.IncludeGroups
	output.Query.QueryName[1], output.Query.QueryValue[1] = "limit", strconv.Itoa(input.Limit)
	output.Query.QueryName[2], output.Query.QueryValue[2] = "market", input.Market
	output.Query.QueryName[3], output.Query.QueryValue[3] = "offset", strconv.Itoa(input.Offset)
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

func ArtistsTopTracks(input ArtistsTopTracksInput) (ArtistsTopTracksResponse, error) {
	var output spotify.OutputToSpotify
	var response ArtistsTopTracksResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "artists/" + input.Id + "/top-track"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "market", input.Market
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

func GetArtistsRelatedArtists(input GetArtistsRelatedArtistsInput) (GetArtistsRelatedArtistsResponse, error) {
	var output spotify.OutputToSpotify
	var response GetArtistsRelatedArtistsResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "artists/" + input.Id + "/related-artists"

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
