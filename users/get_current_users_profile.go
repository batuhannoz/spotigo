package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/get-current-users-profile
//
// Get detailed profile information about the current user (including the current user's username).
func GetCurrentUsersProfile(token string) (CurrentUsersProfileResponse, error) {
	response := CurrentUsersProfileResponse{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetCurrentUsersProfileURL, nil)
	if err != nil {
		return response, err
	}
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
