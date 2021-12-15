package spotify

import (
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
	"strconv"
)

func GetAvailableDevices(input GetAvailableDevicesInput) (GetAvailableDevicesResponse, error) {
	var output spotify.OutputToSpotify
	var response GetAvailableDevicesResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/player/devices"

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetPlaybackState(input GetPlaybackStateInput) (GetPlaybackStateResponse, error) {
	var output spotify.OutputToSpotify
	var response GetPlaybackStateResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/player"
	output.Query.QueryName[0], output.Query.QueryValue[0] = "additional_types", input.AdditionalTypes
	output.Query.QueryName[1], output.Query.QueryValue[1] = "market", input.Market
	output.Query.QueryNumber = 2

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetCurrentlyPlayingTrack(input GetCurrentlyPlayingTrackInput) (GetCurrentlyPlayingResponse, error) {
	var output spotify.OutputToSpotify
	var response GetCurrentlyPlayingResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/player/currently-playing"
	output.Query.QueryName[0], output.Query.QueryValue[0] = "additional_types", input.AdditionalTypes
	output.Query.QueryName[1], output.Query.QueryValue[1] = "market", input.Market
	output.Query.QueryNumber = 2

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func StartResumePlayback(input StartResumePlaybackInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/play"
	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryNumber = 1

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func PausePlayback(input PausePlaybackInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/pause"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryNumber = 1

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func SkipToNext(input SkipToNextInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPost
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/next"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryNumber = 1

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func SkipToPrevious(input SkipToPreviousInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPost
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/next"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryNumber = 1

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func SeekToPosition(input SeekToPositionInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/seek"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "position_ms", strconv.Itoa(input.PositionMS)
	output.Query.QueryNumber = 2

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func SetRepeatMode(input SetRepeatModeInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/repeat"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "state", input.RepeatModeState
	output.Query.QueryNumber = 2

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func SetPlaybackVolume(input SetPlaybackVolumeInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/volume"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "volume_percent", strconv.Itoa(input.VolumePercent)
	output.Query.QueryNumber = 2

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func TogglePlaybackShuffle(input TogglePlaybackShuffleInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/shuffle"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "state", strconv.FormatBool(input.ShuffleState)
	output.Query.QueryNumber = 2

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func GetRecentlyPlayedTracks(input GetRecentlyPlayedTracksInput) (GetRecentlyPlayedTracksResponse, error) {
	var output spotify.OutputToSpotify
	var response GetRecentlyPlayedTracksResponse
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "me/player/recently-played"
	output.Query.QueryName[0], output.Query.QueryValue[0] = "after", strconv.Itoa(input.After)
	output.Query.QueryName[1], output.Query.QueryValue[1] = "before", strconv.Itoa(input.Before)
	output.Query.QueryName[2], output.Query.QueryValue[2] = "limit", strconv.Itoa(input.Limit)
	output.Query.QueryNumber = 3

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}

func AddItemPLaybackQueue(input AddItemPLaybackQueueInput) (string, error) {
	var output spotify.OutputToSpotify
	var response string
	output.ResponseType = &response
	output.Token = input.Token
	output.MethodType = http.MethodPut
	output.TrueStatusCode = http.StatusNoContent
	output.Url = spotify.BaseUrl + "me/player/queue"

	output.Query.QueryName[0], output.Query.QueryValue[0] = "device_id", input.DeviceId
	output.Query.QueryName[1], output.Query.QueryValue[1] = "uri", input.Uri
	output.Query.QueryNumber = 2

	err := spotify.SpotigoToSpotify(output)
	if err != nil {
		return response, err
	}

	return response, nil
}
