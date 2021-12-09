package spotify

import (
	"net/http"
)

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
