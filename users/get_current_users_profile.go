package spotify

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Get detailed profile information about the current user (including the current user's username).
func GetCurrentUsersProfile(token string) (CurrentUsersProfileResponse, error) {
	response := CurrentUsersProfileResponse{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, GetCurrentUsersProfileURL, nil)

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
