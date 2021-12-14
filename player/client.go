package spotify

import (
	"bytes"
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func GetAvailableDevices(input spotify.UserInfo) (GetAvailableDevicesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetAvailableDevicesResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/player/devices"

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

func GetPlaybackState(input spotify.UserInfo) (GetPlaybackStateResponse, error) {
	var output spotify.OutputToSpotify
	var response GetPlaybackStateResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/player"
	output.Query.QueryName[0], output.Query.QueryValue[0] = "additional_types", input.AdditionalTypes
	output.Query.QueryName[1], output.Query.QueryValue[1] = "market", input.Market
	output.Query.QueryNumber = 2

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

func GetCurrentlyPlayingTrack(input spotify.UserInfo) (GetPlaybackStateResponse, error) {
	var output spotify.OutputToSpotify
	var response GetPlaybackStateResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/player/currently-playing"
	output.Query.QueryName[0], output.Query.QueryValue[0] = "additional_types", input.AdditionalTypes
	output.Query.QueryName[1], output.Query.QueryValue[1] = "market", input.Market
	output.Query.QueryNumber = 2

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

func StartResumePlayback(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	var request StartResumePlaybackRequest
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/play"
	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryNumber = 1

	request.ContextUri = input.Uri
	request.PositionMs = input.PositionMS
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return response, err
	}
	output.Body = bytes.NewBuffer(jsonBody)

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func PausePlayback(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/pause"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryNumber = 1

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func SkipToNext(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodPost
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/next"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryNumber = 1

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func SkipToPrevious(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodPost
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/next"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryNumber = 1

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func SeekToPosition(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/seek"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "position_ms", string(input.PositionMS)
	output.Query.QueryNumber = 2

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func SetRepeatMode(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/repeat"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "state", input.RepeatModeState
	output.Query.QueryNumber = 2

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func SetPlaybackVolume(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/volume"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "volume_percent", string(input.VolumePercent)
	output.Query.QueryNumber = 2

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func TogglePlaybackShuffle(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/shuffle"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "state", strconv.FormatBool(input.ShuffleState)
	output.Query.QueryNumber = 2

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}

func GetRecentlyPlayedTracks(input spotify.UserInfo) (GetRecentlyPlayedTracksResponse, error) {
	var output spotify.OutputToSpotify
	var response GetRecentlyPlayedTracksResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/player/recently-played"
	output.Query.QueryName[0], output.Query.QueryValue[0] = "after", string(input.After)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "before", string(input.Before)
	output.Query.QueryName[2], output.Query.QueryValue[2] = "limit", string(input.Limit)
	output.Query.QueryNumber = 3

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

func AddItemPLaybackQueue(input spotify.UserInfo) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/queue"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "uri", input.Uri
	output.Query.QueryNumber = 2

	res, err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	response = res.Status

	return response, nil
}
