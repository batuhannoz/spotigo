package spotify

/*
type SearchResponse struct {
	Tracks struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"tracks"`
	Artists struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"artists"`
	Albums struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"albums"`
	Playlists struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"playlists"`
	Shows struct {
		Href  string `json:"href"`
		Items []struct {
		} `json:"items"`
		Limit    int    `json:"limit"`
		Next     string `json:"next"`
		Offset   int    `json:"offset"`
		Previous string `json:"previous"`
		Total    int    `json:"total"`
	} `json:"shows"`
	Episodes struct {
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

func Search(input spotify.UserInfo) (SearchResponse, error) {
	var output spotify.OutputToSpotify
	var response SearchResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "search"

	// Prepare the queries that will go to the main spotify function.
	output.Query.QueryName[0], output.Query.QueryValue[0] = "offset", string(input.Offset)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "market", input.Market
	output.Query.QueryName[2], output.Query.QueryValue[2] = "limit", string(input.Limit)
	output.Query.QueryName[3], output.Query.QueryValue[3] = "include_external", input.IncludeExternal
	// TODO figure out how to type
	//output.Query.QueryName[4], output.Query.QueryValue[4] = "type", input.SearchType
	output.Query.QueryName[5], output.Query.QueryValue[5] = "limit", string(input.Limit)

	output.Query.QueryNumber = 6

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
*/
