package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

func AvailableMarkets(input AvailableMarketsInput) (AvailableMarketsResponse, error) {
	var output spotify.OutputToSpotify
	var response AvailableMarketsResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "markets/"

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}
