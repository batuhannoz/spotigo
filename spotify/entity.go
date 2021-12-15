package spotify

import (
	"io"
)

const (
	SpotifyAuthName = "Authorization"
	BaseUrl         = "https://api.spotify.com/v1/"
)

type OutputToSpotify struct {
	Token          string
	MethodType     string
	TrueStatusCode int
	Url            string
	Body           io.Reader
	Query          Query
	ResponseType   interface{}
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
