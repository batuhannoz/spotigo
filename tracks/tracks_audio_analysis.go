package spotify

import (
	"encoding/json"
	"github.com/batuhannoz/spotigo/spotify"
	"net/http"
)

type GetTracksAudioAnalysisResponse struct {
	Meta struct {
		AnalyzerVersion string  `json:"analyzer_version"`
		Platform        string  `json:"platform"`
		DetailedStatus  string  `json:"detailed_status"`
		StatusCode      int     `json:"status_code"`
		Timestamp       int     `json:"timestamp"`
		AnalysisTime    float64 `json:"analysis_time"`
		InputProcess    string  `json:"input_process"`
	} `json:"meta"`
	Track struct {
		NumSamples              int     `json:"num_samples"`
		Duration                float64 `json:"duration"`
		SampleMd5               string  `json:"sample_md5"`
		OffsetSeconds           int     `json:"offset_seconds"`
		WindowSeconds           int     `json:"window_seconds"`
		AnalysisSampleRate      int     `json:"analysis_sample_rate"`
		AnalysisChannels        int     `json:"analysis_channels"`
		EndOfFadeIn             int     `json:"end_of_fade_in"`
		StartOfFadeOut          float64 `json:"start_of_fade_out"`
		Loudness                float64 `json:"loudness"`
		Tempo                   float64 `json:"tempo"`
		TempoConfidence         float64 `json:"tempo_confidence"`
		TimeSignature           int     `json:"time_signature"`
		TimeSignatureConfidence float64 `json:"time_signature_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float64 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float64 `json:"mode_confidence"`
		Codestring              string  `json:"codestring"`
		CodeVersion             float64 `json:"code_version"`
		Echoprintstring         string  `json:"echoprintstring"`
		EchoprintVersion        float64 `json:"echoprint_version"`
		Synchstring             string  `json:"synchstring"`
		SynchVersion            int     `json:"synch_version"`
		Rhythmstring            string  `json:"rhythmstring"`
		RhythmVersion           int     `json:"rhythm_version"`
	} `json:"track"`
	Bars []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"bars"`
	Beats []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"beats"`
	Sections []struct {
		Start                   int     `json:"start"`
		Duration                float64 `json:"duration"`
		Confidence              int     `json:"confidence"`
		Loudness                float64 `json:"loudness"`
		Tempo                   float64 `json:"tempo"`
		TempoConfidence         float64 `json:"tempo_confidence"`
		Key                     int     `json:"key"`
		KeyConfidence           float64 `json:"key_confidence"`
		Mode                    int     `json:"mode"`
		ModeConfidence          float64 `json:"mode_confidence"`
		TimeSignature           int     `json:"time_signature"`
		TimeSignatureConfidence int     `json:"time_signature_confidence"`
	} `json:"sections"`
	Segments []struct {
		Start           float64   `json:"start"`
		Duration        float64   `json:"duration"`
		Confidence      float64   `json:"confidence"`
		LoudnessStart   float64   `json:"loudness_start"`
		LoudnessMax     float64   `json:"loudness_max"`
		LoudnessMaxTime float64   `json:"loudness_max_time"`
		LoudnessEnd     int       `json:"loudness_end"`
		Pitches         []float64 `json:"pitches"`
		Timbre          []float64 `json:"timbre"`
	} `json:"segments"`
	Tatums []struct {
		Start      float64 `json:"start"`
		Duration   float64 `json:"duration"`
		Confidence float64 `json:"confidence"`
	} `json:"tatums"`
}

func GetTracksAudioAnalysis(input spotify.UserInfo) (GetTracksAudioAnalysisResponse, error) {
	var output spotify.OutputToSpotify
	var response GetTracksAudioAnalysisResponse
	output.Token = input.Token
	output.MethodType = http.MethodGet
	output.TrueStatusCode = http.StatusOK
	output.Url = spotify.BaseUrl + "audio-analysis/" + input.Id

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
