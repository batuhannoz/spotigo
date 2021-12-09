package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/get-recently-played
//
// Get tracks from the current user's recently played tracks. Note: Currently doesn't support podcast episodes.
//
// beforeorafter can only be one of: before, after
func GetRecentlyPlayedTracks(token string, beforeOrAfter string, unixTimestap int, limit int) (GetRecentlyPlayedTracksResponse, error) {
	response := GetRecentlyPlayedTracksResponse{}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetRecentlyPlayedTracksURL, nil)
	if err != nil {
		return response, err
	}

	q := req.URL.Query()
	q.Add(beforeOrAfter, strconv.Itoa(unixTimestap))
	q.Add("limit", strconv.Itoa(limit))
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return response, err
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return response, err
	}

	bodyString := string(bodyBytes)
	err = json.Unmarshal([]byte(bodyString), &response)
	if err != nil {
		return response, err
	}

	defer res.Body.Close()

	return response, nil
}
