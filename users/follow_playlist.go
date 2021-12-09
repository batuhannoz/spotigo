package spotify

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/follow-playlist
//
// Add the current user as a follower of a playlist.
func FollowPlaylist(token string, playlistID string, public bool) error {
	var reqBody FollowPlaylistRequest
	reqBody.Public = public
	client := &http.Client{}
	jsonBody, err := json.Marshal(reqBody)
	req, err := http.NewRequest(http.MethodPut, FollowPlaylistURL+playlistID+"/followers", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return nil
}
