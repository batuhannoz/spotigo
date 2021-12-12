package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type GetEpisodesResponse struct {
	AudioPreviewUrl string `json:"audio_preview_url"`
	Description     string `json:"description"`
	HtmlDescription string `json:"html_description"`
	DurationMs      int    `json:"duration_ms"`
	Explicit        bool   `json:"explicit"`
	ExternalUrls    struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Href   string `json:"href"`
	Id     string `json:"id"`
	Images []struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"images"`
	IsExternallyHosted   bool     `json:"is_externally_hosted"`
	IsPlayable           bool     `json:"is_playable"`
	Language             string   `json:"language"`
	Languages            []string `json:"languages"`
	Name                 string   `json:"name"`
	ReleaseDate          string   `json:"release_date"`
	ReleaseDatePrecision string   `json:"release_date_precision"`
	ResumePoint          struct {
		FullyPlayed      bool `json:"fully_played"`
		ResumePositionMs int  `json:"resume_position_ms"`
	} `json:"resume_point"`
	Type         string `json:"type"`
	Uri          string `json:"uri"`
	Restrictions struct {
		Reason string `json:"reason"`
	} `json:"restrictions"`
	Show struct {
		AvailableMarkets []string `json:"available_markets"`
		Copyrights       []struct {
			Text string `json:"text"`
			Type string `json:"type"`
		} `json:"copyrights"`
		Description     string `json:"description"`
		HtmlDescription string `json:"html_description"`
		Explicit        bool   `json:"explicit"`
		ExternalUrls    struct {
			Spotify string `json:"spotify"`
		} `json:"external_urls"`
		Href   string `json:"href"`
		Id     string `json:"id"`
		Images []struct {
			Url    string `json:"url"`
			Height int    `json:"height"`
			Width  int    `json:"width"`
		} `json:"images"`
		IsExternallyHosted bool     `json:"is_externally_hosted"`
		Languages          []string `json:"languages"`
		MediaType          string   `json:"media_type"`
		Name               string   `json:"name"`
		Publisher          string   `json:"publisher"`
		Type               string   `json:"type"`
		Uri                string   `json:"uri"`
	} `json:"show"`
}

func GetEpisodes(input spotify.UserInfo) (GetEpisodesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetEpisodesResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "episodes/" + input.Id

	output.Query.QueryName[0], output.Query.QueryValue[0] = "market", input.Market
	output.Query.QueryNumber = 1

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	return response, nil
}
