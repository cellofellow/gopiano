package gopiano

import (
	"bytes"
	"encoding/json"

	"github.com/cellofellow/gopiano/requests"
	"github.com/cellofellow/gopiano/responses"
)

func (c *Client) StationAddFeedback(trackToken string, isPositive bool) (*responses.StationAddFeedback, error) {
	requestData := requests.StationAddFeedback{
		TrackToken: trackToken,
		IsPositive: isPositive,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationAddFeedback
	err = c.BlowfishCall("http://", "station.addFeedback", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) StationAddMusic(musicToken, stationToken string) (*responses.StationAddMusic, error) {
	requestData := requests.StationAddMusic{
		MusicToken: musicToken,
		StationToke: stationToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationAddMusic
	err = c.BlowfishCall("http://", "station.addMusic", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationCreateStationTrack creates a new station from a specified track.
// trackToken is a token of a song or artist obtianed from Client.StationGetPlaylist.
// musicType is either "song" or "artist" specifying the type of track being used.
func (c *Client) StationCreateStationTrack (trackToken, musicType string) (*responses.StationCreateStation, error) {
	resquestData := requests.StationCreatStation{
		TrackToken: trackToken,
		MusicType: musicType,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationCreateStation
	err = c.BlowfishCall("http://", "station.createStation", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Client.StationCreateStationMusic creates a new station from a music search result.
// musicToken is  obtianed from Client.MusicSearch.
func (c *Client) StationCreateStationMusic (musicToken string) (*responses.StationCreateStation, error) {
	resquestData := requests.StationCreatStation{
		MusicToken: musicToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.StationCreateStation
	err = c.BlowfishCall("http://", "station.createStation", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
