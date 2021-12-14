package spotify

type StartResumePlaybackRequest struct {
	ContextUri string `json:"context_uri"`
	Offset     struct {
		Position int `json:"position"`
	} `json:"offset"`
	PositionMs int `json:"position_ms"`
}
