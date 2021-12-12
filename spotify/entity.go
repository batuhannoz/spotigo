package spotify

import "io"

const (
	SpotifyAuthName = "Authorization"
	BaseUrl         = "https://api.spotify.com/v1/"
)

type UserInfo struct {
	Token           string
	Country         string
	Market          string
	Limit           int
	Offset          int
	Ids             string
	Id              string
	Uri             string
	DeviceId        string
	After           int
	Before          int
	RepeatModeState string
	ShuffleState    bool
	VolumePercent   int
	PositionMS      int
	contextURI      string
	Uris            []string
	AdditionalTypes string
	Play            bool
	CategoryId      string
	Locale          string
	PlaylistId      string
	categoryId      string
	timestamp       string
	UserId          string
	Name            string
	Public          bool
	Collaborative   bool
	Description     string
	SnapshotId      string
	DeviceIds       []string
	IncludeGroups   string
	Fields          string
	IncludeExternal string
	SearchType      []string
	TopItemsType    string
}

type OutputToSpotify struct {
	Token          string
	UserInfo       UserInfo
	MethodType     string
	TrueStatusCode int
	Url            string
	Body           io.Reader
	Query          Query
}

type Query struct {
	QueryName   [15]string
	QueryValue  [15]string
	QueryNumber int
}

type SpotifyErrorResponse struct {
	Error struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	} `json:"error"`
}
