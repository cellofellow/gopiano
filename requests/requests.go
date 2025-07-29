/*
Structs for use with json.Marshal when sending requests to the Pandora API.
*/
package requests

type AuthPartnerLogin struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DeviceModel string `json:"deviceModel"`
	Version     string `json:"version"`
	IncludeURLs bool   `json:"includeUrls,omitempty"`
}

type AuthUserLogin struct {
	PartnerAuthToken              string `json:"partnerAuthToken"`
	Username                      string `json:"username"`
	Password                      string `json:"password"`
	LoginType                     string `json:"loginType"` // Should always be "user"
	SyncTime                      int    `json:"syncTime"`
	IncludeAdAttributes           bool   `json:"includeAdAttributes,omitempty"`
	IncludeDemographics           bool   `json:"IncludeDemographics,omitempty"`
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
	UserGetBookmarks           userTokenGeneric
	UserGetStationListChecksum userTokenGeneric
	UserCanSubscribe           userTokenGeneric
)

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

type UserEmailPassword struct {
	PartnerAuthToken string `json:"partnerAuthToken"`
	SyncTime         int    `json:"syncTime"`
	Username         string `json:"username"`
}

type UserGetStationList struct {
	IncludeStationArtURL bool   `json:"includeStationArtUrl,omitempty"`
	SyncTime             int    `json:"syncTime"`
	UserAuthToken        string `json:"userAuthToken"`
}

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
	UserSleepSong             trackAction
	BookmarkAddArtistBookmark trackAction
	BookmarkAddSongBookmark   trackAction
)

type MusicSearch struct {
	SearchText    string `json:"searchText"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type StationCreateStation struct {
	MusicToken    string `json:"musicToken,omitempty"`
	TrackToken    string `json:"trackToken,omitempty"`
	MusicType     string `json:"musicType,omitempty"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type StationDeleteStation struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type StationAddFeedback struct {
	TrackToken    string `json:"trackToken"`
	IsPositive    bool   `json:"isPositive"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type StationDeleteFeedback struct {
	FeedbackID    string `json:"feedbackId"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type StationAddMusic struct {
	MusicToken    string `json:"musicToken"`
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type StationDeleteMusic struct {
	SeedID        string `json:"seedId"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type (
	StationGetGenreStations         userTokenGeneric
	StationGetGenreStationsChecksum userTokenGeneric
)

type StationGetPlaylist struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type StationGetStation struct {
	StationToken              string `json:"stationToken"`
	IncludeExtendedAttributes bool   `json:"includeExtendedAttributes,omitempty"`
	SyncTime                  int    `json:"syncTime"`
	UserAuthToken             string `json:"userAuthToken"`
}

type StationShareStation struct {
	StationID     string   `json:"stationId"`
	StationToken  string   `json:"stationToken"`
	Emails        []string `json:"emails"`
	SyncTime      int      `json:"syncTime"`
	UserAuthToken string   `json:"userAuthToken"`
}

type StationRenameStation struct {
	StationToken  string `json:"stationToken"`
	StationName   string `json:"stationName"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type StationTransformSharedStation struct {
	StationToken  string `json:"stationToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}

type ExplainTrack struct {
	TrackToken    string `json:"trackToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}
