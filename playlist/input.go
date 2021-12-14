package spotify

type GetPlaylistsInput struct {
	Token           string
	AdditionalTypes string
	Fields          string
	Market          string
	PlaylistId      string
}
