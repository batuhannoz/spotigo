package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/get-users-top-artists-and-tracks
//
// Get the current user's top tracks based on calculated affinity.
//
// limit Default: 20. Minimum: 1. Maximum: 50. timeRange can only be one of: "short_term", "medium_term", "long_term".
func GetUsersTopTracks(token string, limit int, offset int, timeRange string) (GetUsersTopTracksResponse, error) {
	response := GetUsersTopTracksResponse{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetUsersTopTracksURL, nil)
	if err != nil {
		return response, err
	}
	q := req.URL.Query()
	q.Add("limit", strconv.Itoa(limit))
	q.Add("offset", strconv.Itoa(offset))
	q.Add("time_range", timeRange)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	bodyString := string(bodyBytes)
	err = json.Unmarshal([]byte(bodyString), &response)
	if err != nil {
		return response, err
	}

	return response, nil
}
