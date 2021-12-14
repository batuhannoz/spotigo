package spotify

type GetUsersSavedTracksInput struct {
	Token  string
	Limit  int
	Market string
	Offset int
}

type GetTrackInput struct {
	Token  string
	Id     string
	Market string
}

type GetTracksAudioAnalysisInput struct {
	Token string
	Id    string
}
