package spotify

type GetUsersSavedShowsInput struct {
	Token  string
	Limit  int
	Offset int
}

type GetShowsInput struct {
	Token  string
	Id     string
	Market string
}

type GetShowEpisodesInput struct {
	Token  string
	Limit  int
	Market string
	Offset int
	Id     string
}
