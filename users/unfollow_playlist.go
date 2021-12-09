package spotify

import (
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/unfollow-playlist
//
// Remove the current user as a follower of a playlist.
func UnfollowPlaylist(token string, playlistID string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodDelete, UnfollowPlaylistURL+playlistID+"/followers", nil)
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
