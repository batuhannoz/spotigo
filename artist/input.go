package spotify

type GetArtistInput struct {
	Token string
	Id    string
}

type SingleBrowseCategoriesInput struct {
	Token         string
	Id            string
	IncludeGroups string
	Limit         int
	Market        string
	Offset        int
}

type ArtistsTopTracksInput struct {
	Token  string
	Id     string
	Market string
}

type GetArtistsRelatedArtistsInput struct {
	Token string
	Id    string
}
