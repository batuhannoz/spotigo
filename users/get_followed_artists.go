package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/get-followed
//
// Get the current user's followed artists.
func GetFollowedArtists(token string, followedType string, after string, limit int) (GetFollowedArtistsResponse, error) {
	response := GetFollowedArtistsResponse{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetFollowedArtistsURL, nil)
	if err != nil {
		return response, err
	}
	q := req.URL.Query()
	q.Add("type", followedType)
	q.Add("after", after)
	q.Add("limit", strconv.Itoa(limit))
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
