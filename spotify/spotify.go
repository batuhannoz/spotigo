package spotify

import (
	"encoding/json"
	"errors"
	"net/http"
)

func SpotigoToSpotify(request OutputToSpotify) (*http.Response, error) {
	var response *http.Response
	client := &http.Client{}

	// Create request
	spotifyReq, err := http.NewRequest(request.MethodType, request.Url, request.Body)
	if err != nil {
		return response, err
	}

	// Add query.
	q := spotifyReq.URL.Query()
	if request.Query.QueryName[0] != "" {
		for i := 0; i <= request.Query.QueryNumber-1; i++ {
			if request.Query.QueryValue[i] != "" {
				q.Add(request.Query.QueryName[i], request.Query.QueryValue[i])
			}
		}
	}
	spotifyReq.URL.RawQuery = q.Encode()

	// Add header.
	if request.Token != "" {
		spotifyReq.Header.Add(SpotifyAuthName, request.Token)
	}

	// Do Request.
	response, err = client.Do(spotifyReq)
	if err != nil {
		return response, err
	}

	// Check spotify response status code.
	if response.StatusCode != request.TrueStatusCode {
		if response.StatusCode == http.StatusUnauthorized || response.StatusCode == http.StatusForbidden || response.StatusCode == http.StatusTooManyRequests {
			errorResponse := SpotifyErrorResponse{}
			err = json.NewDecoder(response.Body).Decode(&errorResponse)
			if err != nil {
				return response, err
			}
			err = errors.New(errorResponse.Error.Message)
			return response, err
		}
		return response, errors.New(response.Status)
	}

	return response, nil
}

/*
func GeneralSpotifyFunc(input spotify.UserInfo) (GeneralSpotifyFuncResponse, error) {
	// Type to be given to the main spotify function
	var output spotify.OutputToSpotify

	// If GeneralSpotifyFunc nothing's return => response type = string,
	var response GeneralSpotifyFuncResponse

	// If making json request => use requestBody.
	var request GeneralSpotifyFuncRequest

	// Assign given token to variable.
	output.Token = input.Token

	// Determine the method type according to the request to be made.
	output.MethodType = http.MethodPost

	// Write the desired status code.
	output.TrueStatusCode = http.StatusOK


	// Create the url of the request to be made.
	output.Url = spotify.BaseUrl + "me/player/queue"

	// In some cases, variable type should be changed in query and body.

	// Prepare the queries that will go to the main spotify function.
	output.Query.QueryName[0], output.Query.QueryValue[0] = "playlist_id", input.PlaylistId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "uri", input.Uri
	// Number of queries added
	output.Query.QueryNumber = 2

	// Prepare json body.
	requestBody.ContextUri = input.contextURI
	requestBody.PositionMS = string(input.PositionMS)
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return response, err
	}
	output.Body = bytes.NewBuffer(jsonBody)

	// Pass information to main function.
	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	// Don't use 1 and 2 at the same time.
	// 1) use this if it will do struct return
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return response, err
	}

	// 2) Use this if nothing will return
	response = res.Status

	return response, nil
}
*/
