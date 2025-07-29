// Package requests provides structs for use with json.Marshal when sending requests to the Pandora API.
package requests

// AuthPartnerLogin represents the request data for auth.partnerLogin.
type AuthPartnerLogin struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DeviceModel string `json:"deviceModel"`
	Version     string `json:"version"`
	IncludeURLs bool   `json:"includeUrls,omitempty"`
}

// AuthUserLogin represents the request data for auth.userLogin.
type AuthUserLogin struct {
	PartnerAuthToken              string `json:"partnerAuthToken"`
	Username                      string `json:"username"`
	Password                      string `json:"password"`
	LoginType                     string `json:"loginType"` // Should always be "user"
	SyncTime                      int    `json:"syncTime"`
	IncludeAdAttributes           bool   `json:"includeAdAttributes,omitempty"`
	IncludeDemographics           bool   `json:"IncludeDemographics,omitempty"`   //nolint:tagliatelle // matches Pandora API
	IncludePandoraOneInfo         bool   `json:"includePandoraOneInfo,omitempty"` // Appears to do nothing.
	IncludeStationArtURL          bool   `json:"includeStationArtUrl,omitempty"`
	IncludeSubscriptionExpiration bool   `json:"includeSubscriptionExpiration,omitempty"`
	ReturnCapped                  bool   `json:"returnCapped,omitempty"`
	ReturnGenreStations           bool   `json:"returnGenreStations,omitempty"`
	ReturnStationList             bool   `json:"returnStationList,omitempty"`
}

type userTokenGeneric struct {
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}
type (
	// UserGetBookmarks represents the request data for user.getBookmarks.
	UserGetBookmarks userTokenGeneric
	// UserGetStationListChecksum represents the request data for user.getStationListChecksum.
	UserGetStationListChecksum userTokenGeneric
	// UserCanSubscribe represents the request data for user.canSubscribe.
	UserCanSubscribe userTokenGeneric
)

// UserCreateUser represents the request data for user.createUser.
type UserCreateUser struct {
	AccountType      string `json:"accountType"`
	BirthYear        int    `json:"birthYear"`
	CountryCode      string `json:"countryCode"`
	EmailOptin       bool   `json:"emailOptin"`
	Gender           string `json:"gender"`
	PartnerAuthToken string `json:"partnerAuthToken"`
	Password         string `json:"password"`
	RegisteredType   string `json:"registeredType"`
	SyncTime         int    `json:"syncTime"`
	Username         string `json:"username"`
	ZipCode          int    `json:"zip"`
}

// UserEmailPassword represents the request data for user.emailPassword.
type UserEmailPassword struct {
	PartnerAuthToken string `json:"partnerAuthToken"`
	SyncTime         int    `json:"syncTime"`
	Username         string `json:"username"`
}

// UserGetStationList represents the request data for user.getStationList.
type UserGetStationList struct {
	IncludeStationArtURL bool   `json:"includeStationArtUrl,omitempty"`
	SyncTime             int    `json:"syncTime"`
	UserAuthToken        string `json:"userAuthToken"`
}

// UserSetQuickMix represents the request data for user.setQuickMix.
type UserSetQuickMix struct {
	QuickMixStationIDs []string `json:"quickMixStationIds"`
	SyncTime           int      `json:"syncTime"`
	UserAuthToken      string   `json:"userAuthToken"`
}

type trackAction struct {
	TrackToken    string `json:"trackToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}
type (
	// UserSleepSong represents the request data for user.sleepSong.
	UserSleepSong trackAction
	// BookmarkAddArtistBookmark represents the request data for bookmark.addArtistBookmark.
	BookmarkAddArtistBookmark trackAction
	// BookmarkAddSongBookmark represents the request data for bookmark.addSongBookmark.
	BookmarkAddSongBookmark trackAction
)

// MusicSearch represents the request data for music.search.
type MusicSearch struct {
	SearchText    string `json:"searchText"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationCreateStation represents the request data for station.createStation.
type StationCreateStation struct {
	MusicToken    string `json:"musicToken,omitempty"`
	TrackToken    string `json:"trackToken,omitempty"`
	MusicType     string `json:"musicType,omitempty"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationDeleteStation represents the request data for station.deleteStation.
type StationDeleteStation struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationAddFeedback represents the request data for station.addFeedback.
type StationAddFeedback struct {
	TrackToken    string `json:"trackToken"`
	IsPositive    bool   `json:"isPositive"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationDeleteFeedback represents the request data for station.deleteFeedback.
type StationDeleteFeedback struct {
	FeedbackID    string `json:"feedbackId"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationAddMusic represents the request data for station.addMusic.
type StationAddMusic struct {
	MusicToken    string `json:"musicToken"`
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationDeleteMusic represents the request data for station.deleteMusic.
type StationDeleteMusic struct {
	SeedID        string `json:"seedId"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type (
	// StationGetGenreStations represents the request data for station.getGenreStations.
	StationGetGenreStations userTokenGeneric
	// StationGetGenreStationsChecksum represents the request data for station.getGenreStationsChecksum.
	StationGetGenreStationsChecksum userTokenGeneric
)

// StationGetPlaylist represents the request data for station.getPlaylist.
type StationGetPlaylist struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationGetStation represents the request data for station.getStation.
type StationGetStation struct {
	StationToken              string `json:"stationToken"`
	IncludeExtendedAttributes bool   `json:"includeExtendedAttributes,omitempty"`
	SyncTime                  int    `json:"syncTime"`
	UserAuthToken             string `json:"userAuthToken"`
}

// StationShareStation represents the request data for station.shareStation.
type StationShareStation struct {
	StationID     string   `json:"stationId"`
	StationToken  string   `json:"stationToken"`
	Emails        []string `json:"emails"`
	SyncTime      int      `json:"syncTime"`
	UserAuthToken string   `json:"userAuthToken"`
}

// StationRenameStation represents the request data for station.renameStation.
type StationRenameStation struct {
	StationToken  string `json:"stationToken"`
	StationName   string `json:"stationName"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// StationTransformSharedStation represents the request data for station.transformSharedStation.
type StationTransformSharedStation struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

// ExplainTrack represents the request data for track.explainTrack.
type ExplainTrack struct {
	TrackToken    string `json:"trackToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}
