package spotify

type GetEpisodesInput struct {
	Token  string
	Id     string
	Market string
}

type GetUsersSavedEpisodesInput struct {
	Token  string
	Limit  int
	Market string
	Offset int
}
