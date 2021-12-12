package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type GetArtistsRelatedArtistsResponse struct {
	Artists []struct {
		ExternalUrls struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Followers struct {
			Href  string `json:"href"`
			Total int    `json:"total"`
		} `json:"followers"`
		Genres []string `json:"genres"`
		Href   string   `json:"href"`
		Id     string   `json:"id"`
		Images []struct {
			Url    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		Name       string `json:"name"`
		Popularity int    `json:"popularity"`
		Type       string `json:"type"`
		Uri        string `json:"uri"`
	} `json:"artists"`
}

func GetArtistsRelatedArtists(input spotify.UserInfo) (GetArtistsRelatedArtistsResponse, error) {
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
