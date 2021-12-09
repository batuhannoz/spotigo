package spotify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/check-if-user-follows-playlist
//
// Check to see if one or more Spotify users are following a specified playlist.
func CheckIfUsersFollowPlaylist(token string, playlistID string, ids string) (CheckIfUsersFollowPlaylistResponse, error) {
	response := CheckIfUsersFollowPlaylistResponse{}
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, CheckIfUsersFollowPlaylistURL+playlistID+"/followers/contains", nil)
	if err != nil {
		return response, err
	}

	q := req.URL.Query()
	q.Add("ids", ids)
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
