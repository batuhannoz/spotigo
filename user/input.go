package spotify

type GetCurrentUsersProfileInput struct {
	Token string
}

type FollowPlaylistInput struct {
	Token      string
	PlaylistId string
	Public     bool
}

type GetUsersProfileInput struct {
	Token  string
	UserId string
}

type GetUsersTopItemsInput struct {
	Token        string
	TopItemsType string
	PlaylistId   string
	Limit        int
	Offset       int
}

type UnfollowPlaylistInput struct {
	Token      string
	PlaylistId string
}

type GetFollowedArtistsInput struct {
	Token        string
	FollowedType string
	After        int
}
