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
type UserGetBookmarks userTokenGeneric
type UserGetStationListChecksum userTokenGeneric
type UserCanSubscribe userTokenGeneric

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
type UserSleepSong trackAction
type BookmarkAddArtistBookmark trackAction
type BookmarkAddSongBookmark trackAction

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
	synctime      int    `json:"synctime"`
	userauthtoken string `json:"userauthtoken"`
}

type StationDeleteFeedback struct {
	FeedbackID    string `json:"feedbackId"`
	synctime      int    `json:"synctime"`
	userauthtoken string `json:"userauthtoken"`
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

type StationGetGenreStations userTokenGeneric
type StationGetGenreStationsChecksum userTokenGeneric

type audioFormat string

const AAC_40 audioFormat = "HTTP_40_AAC_MONO"
const AAC_64 audioFormat = "HTTP_64_AAC"
const AACPLUS_32 audioFormat = "HTTP_32_AACPLUS"
const AACPLUS_64 audioFormat = "HTTP_64_AACPLUS"
const AACPLUS_ADTS_24 audioFormat = "HTTP_24_AACPLUS_ADTS"
const AACPLUS_ADTS_32 audioFormat = "HTTP_32_AACPLUS_ADTS"
const AACPLUS_ADTS_64 audioFormat = "HTTP_64_AACPLUS_ADTS"
const MP3_128 audioFormat = "HTTP_128_MP3"
const WMA_32 audioFormat = "WMA_32"

type StationGetPlaylist struct {
	StationToken       string        `json:"stationToken"`
	AdditionalAudioURL []audioFormat `json:"additionalAudioUrl,omitempty"`
	SyncTime           int           `json:"syncTime"`
	UserAuthToken      string        `json:"userAuthToken"`
}

type StationGetStation struct {
	StationToken              string `json:"stationToken"`
	IncludeExtendedAttributes string `json:"includeExtendedAttributes,omitempty"`
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

type TrackExplainTrack struct {
	TrackToken    string `json:"trackToken"`
	SyncTime      int    `json:"syncTime"`
	UserAuthToken string `json:"userAuthToken"`
}
