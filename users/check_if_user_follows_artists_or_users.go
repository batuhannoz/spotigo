package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/check-current-user-follows
//
// Check to see if the current user is following one or more artists or other Spotify users.
//
// idType can be one of: "user", "artist",
func CheckIfUserFollowsArtistsOrUsers(token string, ids string, idType string) (CheckIfUserFollowsArtistsOrUsersResponse, error) {
	response := CheckIfUserFollowsArtistsOrUsersResponse{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, CheckIfUserFollowsArtistsOrUsersURL, nil)
	if err != nil {
		return response, err
	}

	q := req.URL.Query()
	q.Add("ids", ids)
	q.Add("type", idType)
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
