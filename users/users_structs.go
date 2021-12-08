package spotify

type CurrentUsersProfileResponse struct {
	Country         string `json:"country"`
	DisplayName     string `json:"display_name"`
	Email           string `json:"email"`
	ExplicitContent struct {
		FilterEnabled bool `json:"filter_enabled"`
		FilterLocked  bool `json:"filter_locked"`
	} `json:"explicit_content"`
	ExternalUrls struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Href   string `json:"href"`
	Id     string `json:"id"`
	Images []struct {
		Height interface{} `json:"height"`
		Url    string      `json:"url"`
		Width  interface{} `json:"width"`
	} `json:"images"`
	Product string `json:"product"`
	Type    string `json:"type"`
	Uri     string `json:"uri"`
}
