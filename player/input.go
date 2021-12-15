package spotify

type GetAvailableDevicesInput struct {
	Token string
}

type GetPlaybackStateInput struct {
	Token           string
	AdditionalTypes string
	Market          string
}

type GetCurrentlyPlayingTrackInput struct {
	Token           string
	AdditionalTypes string
	Market          string
}

type StartResumePlaybackInput struct {
	Token    string
	DeviceId string
}

type PausePlaybackInput struct {
	Token    string
	DeviceId string
}

type SkipToNextInput struct {
	Token    string
	DeviceId string
}

type SkipToPreviousInput struct {
	Token    string
	DeviceId string
}

type SeekToPositionInput struct {
	Token      string
	DeviceId   string
	PositionMS int
}

type SetRepeatModeInput struct {
	Token           string
	DeviceId        string
	RepeatModeState string
}

type SetPlaybackVolumeInput struct {
	Token         string
	DeviceId      string
	VolumePercent int
}

type TogglePlaybackShuffleInput struct {
	Token        string
	DeviceId     string
	ShuffleState bool
}

type GetRecentlyPlayedTracksInput struct {
	Token  string
	After  int
	Before int
	Limit  int
}

type AddItemPLaybackQueueInput struct {
	Token    string
	Uri      string
	DeviceId string
}
