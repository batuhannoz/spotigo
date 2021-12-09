package spotify

import (
	"net/http"
)

// See here for more information => https://developer.spotify.com/documentation/web-api/reference/#/operations/follow-artists-users
//
// Add the current user as a follower of one or more artists or other Spotify users.
//
// idType can be one of: "user", "artist".
func FollowArtistsOrUser(token string, ids string, idType string) error {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, FollowArtistOrUserURL, nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("ids", ids)
	q.Add("type", idType)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return nil
}
