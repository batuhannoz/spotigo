package spotify

type SeveralBrowseCategoriesResponse struct {
	Categories struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"categories"`
}

type SingleCategoriesResponse struct {
	Href  string `json:"href"`
	Icons []struct {
		Url    string `json:"url"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"icons"`
	Id   string `json:"id"`
	Name string `json:"name"`
}
