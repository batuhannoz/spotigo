package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/get-users-profile
//
// Get public profile information about a Spotify user.
func GetUsersProfile(token string) (GetUsersProfileResponse, error) {
	response := GetUsersProfileResponse{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetUsersProfileURL, nil)

	if err != nil {
		fmt.Println(err)
		return response, err
	}

	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	bodyString := string(bodyBytes)
	json.Unmarshal([]byte(bodyString), &response)

	return response, nil
}