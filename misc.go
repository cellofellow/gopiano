package gopiano

import (
	"bytes"
	"encoding/json"

	"github.com/cellofellow/gopiano/requests"
	"github.com/cellofellow/gopiano/responses"
)

// ExplainTrack retrieves an incomplete list of attributes assigned to a specified song by the
// Music Genome Project.
// Calls API method "track.explainTrack".
func (c *Client) ExplainTrack(trackToken string) (*responses.ExplainTrack, error) {
	requestData := requests.ExplainTrack{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.ExplainTrack
	err = c.BlowfishCall("http://", "track.explainTrack", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MusicSearch searches for music, which can be used to create a new or add seeds to a station.
// Calls API method "music.search".
func (c *Client) MusicSearch(searchText string) (*responses.MusicSearch, error) {
	requestData := requests.MusicSearch{
		SearchText:    searchText,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.MusicSearch
	err = c.BlowfishCall("http://", "music.search", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BookmarkAddArtistBookmark bookmarks an artist.
// Argument trackToken is a token of a specific artist.
// Calls API method "bookmark.addArtistBookmark".
func (c *Client) BookmarkAddArtistBookmark(trackToken string) (*responses.BookmarkAddArtistBookmark, error) {
	requestData := requests.BookmarkAddArtistBookmark{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.BookmarkAddArtistBookmark
	err = c.BlowfishCall("http://", "bookmark.addArtistBookmark", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BookmarkAddSongBookmark bookmarks a song.
// Argument trackToken is a token of a specific song.
// Calls API method "bookmark.addSongBookmark".
func (c *Client) BookmarkAddSongBookmark(trackToken string) (*responses.BookmarkAddSongBookmark, error) {
	requestData := requests.BookmarkAddSongBookmark{
		TrackToken:    trackToken,
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.BookmarkAddSongBookmark
	err = c.BlowfishCall("http://", "bookmark.addSongBookmark", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
