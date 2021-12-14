package spotify

type GetUsersSavedShowsResponse struct {
	Href  string `json:"href"`
	Items []struct {
	} `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}

type GetShowsResponse struct {
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
	Episodes           struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"episodes"`
}

type GetShowEpisodesResponse struct {
	Href  string `json:"href"`
	Items []struct {
	} `json:"items"`
	Limit    int    `json:"limit"`
	Next     string `json:"next"`
	Offset   int    `json:"offset"`
	Previous string `json:"previous"`
	Total    int    `json:"total"`
}
