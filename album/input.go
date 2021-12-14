package spotify

type GetAlbumInput struct {
	Token  string
	Id     string
	Market string
}

type GetAlbumTracksInput struct {
	Token  string
	Id     string
	Limit  int
	Market string
	Offset int
}

type GetNewReleasesInput struct {
	Token   string
	Country string
	Limit   int
	Offset  int
}

type GetSavedAlbumsInput struct {
	Token  string
	Market string
	Limit  int
	Offset int
}
