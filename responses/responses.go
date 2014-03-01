package responses

import "fmt"
import "time"

type ErrorResponse struct {
	Stat    string `json:"stat"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Pandora Error: %d %s", e.Code, e.Message)
}

type dateResponse struct {
	nanos          int
	seconds        int
	year           int
	month          int
	hours          int
	time           int
	date           int
	minutes        int
	day            int
	timezoneOffset int
}

func (d dateResponse) Date() time.Time {
	return time.Date(d.year, time.Month(d.month), d.date, d.hours, d.minutes, d.seconds,
		d.nanos, time.FixedZone("Local Time", d.timezoneOffset*60))
}

type AuthPartnerLogin struct {
	Result struct {
		SyncTime         string `json:"syncTime"`
		StationSkipLimit int    `json:"stationSkipLimit"`
		PartnerAuthToken string `json:"partnerAuthToken"`
		PartnerID        string `json:"partnerId"`
		StationSkipUnit  string `json:"stationSkipUnit"`
		DeviceProperties struct {
			VideoAdRefreshInterval int `json:"videoAdRefreshInterval"`
			VideoAdUniqueInterval  int `json:"videoAdUniqueInterval"`
			AdRefreshInterval      int `json:"adRefreshInterval"`
			VideoAdStartInterval   int `json:"videoAdStartInterval"`
		} `json:"deviceProperties"`
		Urls struct {
			AutoComplete string `json:"autoComplete"`
		} `json:"urls"`
	}
}

type AuthUserLogin struct {
	Result struct {
		CanListen                   bool   `json:"canListen"`
		HasAudioAds                 bool   `json:"hasAudioAds"`
		IsCapped                    bool   `json:"isCapped,omitempty"`
		ListeningTimeoutAlertMsgUri string `json:"listeningTimeoutAlertMsgUri"`
		ListeningTimeoutMinutes     string `json:"listeningTimeoutMinutes"`
		MaxStationsAllowed          int    `json:"maxStationsAllowed"`
		MinimumAdRefreshInterval    int    `json:"minimumAdRefreshInterval"`
		NowPlayingURL               string `json:"nowPlayingUrl"`
		SplashScreenAdURL           string `json:"splashScreenAdUrl"`
		StationCreationAdURL        string `json:"stationCreationAdUrl"`
		UserAuthToken               string `json:"userAuthToken"`
		UserID                      string `json:"userId"`
		UserProfileURL              string `json:"userProfileUrl"`
		Username                    string `json:"username"`
		VideoAdURL                  string `json:"videoAdUrl"`
	} `json:"result"`
}

type UserCanSubscribe struct {
	Result struct {
		CanSubscribe bool `json:"canSubscribe"`
		IsSubscriber bool `json:"isSubscriber"`
	} `json:"result"`
}

type UserCreateUser AuthUserLogin

type artistBookmark struct {
	ArtURL        string       `json:"artUrl"`
	ArtistName    string       `json:"artistName"`
	BookmarkToken string       `json:"bookmarkToken"`
	DateCreated   dateResponse `json:"dateCreated"`
	MusicToken    string       `json:"musicToken"`
}
type BookmarkAddArtistBookmark artistBookmark

type songBookmark struct {
	AlbumName     string       `json:"artistName"`
	ArtURL        string       `json:"artUrl"`
	ArtistName    string       `json:"artistName"`
	BookmarkToken string       `json:"bookmarkToken"`
	DateCreated   dateResponse `json:"dateCreated"`
	MusicToken    string       `json:"musicToken"`
	SampleGain    string       `json:"sampleGain"`
	SampleURL     string       `json:"sampleUrl"`
	SongName      string       `json:"songName"`
}
type BookmarkAddSongBookmark songBookmark

type station struct {
	SuppressVideoAds bool         `json:"suppressVideoAds"`
	StationID        string       `json:"stationId"`
	AllowAddMusic    bool         `json:"allowAddMusic"`
	DateCreated      dateResponse `json:"dateCreated"`
	StationDetailURL string       `json:"stationDetailUrl"`
	ArtURL           string       `json:"artUrl"`
	RequiresCleanAds bool         `json:"requiresCleanAds"`
	StationToken     string       `json:"stationToken"`
	StationName      string       `json:"stationName"`
	Music            struct {
		Songs []struct {
			SeedID      string       `json:"seedId"`
			ArtistName  string       `json:"artistName"`
			SongName    string       `json:"songName"`
			DateCreated dateResponse `json:"dateCreated"`
		} `json:"songs"`
		Artists []struct {
			SeedID      string       `json:"seedId"`
			ArtistName  string       `json:"artistName"`
			DateCreated dateResponse `json:"dateCreated"`
		} `json:"songs"`
	} `json:"music"`
	IsShared          bool     `json:"isShared"`
	AllowDelete       bool     `json:"allowDelete"`
	Genre             []string `json:"genre"`
	IsQuickMix        bool     `json:"isQuickMix"`
	AllowRename       bool     `json:"allowRename"`
	StationSharingURL string   `json:"stationSharingUrl"`
	Feedback          struct {
		ThumbsDown []feedbackResponse `json:"thumsDown"`
		ThumbsUp   []feedbackResponse `json:"thumbsUp"`
	} `json:"feedback"`
}

type UserGetBookmarks struct {
	Result struct {
		Artists []artistBookmark `json:"artists"`
		Songs   []songBookmark   `json:"songs"`
	} `json:"result"`
}

type UserGetStationList struct {
	Result struct {
		Stations []station `json:"stations"`
		Checksum string    `json:"checksum"`
	} `json:"result"`
}

type checksumResponse struct {
	Result struct {
		Checksum string `json:"checksum"`
	} `json:"result"`
}
type UserGetStationListChecksum checksumResponse

type MusicSearch struct {
	Result struct {
		NearMatchesAvailable bool   `json:"nearMatchesAvailable"`
		Explanation          string `json:"explanation"`
		Songs                []struct {
			ArtistName string `json:"artistName"`
			MusicToken string `json:"musicToken"`
			SongName   string `json:"songName"`
			Score      int    `json:"score"`
		} `json:"songs"`
		Artists []struct {
			ArtistName  string `json:"artistName"`
			MusicToken  string `json:"musicToken"`
			LikelyMatch bool   `json:"likelyMatch"`
			Score       int    `json:"score"`
		} `json:"artists"`
	} `json:"result"`
}

type feedbackResponse struct {
	ArtistName  string       `json:"artistName"`
	SongName    string       `json:"songName"`
	DateCreated dateResponse `json:"dateCreated"`
	FeedbackID  string       `json:"feedbackId"`
	IsPositive  int          `json:"isPositive"`
}

type StationAddFeedback struct {
	Result feedbackResponse `json:"result"`
}

type StationAddMusic struct {
	Result struct {
		ArtistName  string       `json:"artistName"`
		DateCreated dateResponse `json:"dateCreated"`
		SeedID      string       `json:"seedId"`
	} `json:"result"`
}

type stationResponse struct {
	Result station `json:"result"`
}
type StationCreateStation stationResponse
type StationGetStation stationResponse
type StationRenameStation stationResponse
type StationTransformSharedStation stationResponse

type StationGetGenreStations struct {
	Result struct {
		Categories []struct {
			CategoryName string `json:"categoryName"`
			Stations     []struct {
				StationToken string `json:"stationToken"`
				StationName  string `json:"stationName"`
				StationID    string `json:"stationId"`
			}
		} `json:"categories"`
	} `json:"result"`
}

type StationGetGenreStationsChecksum checksumResponse

type StationGetPlaylist struct {
	Result struct {
		Items []struct {
			TrackToken      string `json:"trackToken"`
			ArtistName      string `json:"artistName"`
			AlbumName       string `json:"albumName"`
			AmazonAlbumURL  string `json:"amazonAlbumUrl"`
			SongExplorerURL string `json:"songExplorerUrl"`
			AlbumArtURL     string `json:"albumArtUrl"`
			ArtistDetailURL string `json:"artistDetailUrl"`
			AudioURLMap     map[string]struct {
				Bitrate  string `json:"bitrate"`
				Encoding string `json:"encoding"`
				AudioURL string `json:"audioUrl"`
				Protocol string `json:"protocol"`
			} `json:"audioUrlMap"`
			ITunesSongURL          string   `json:"itunesSongUrl"`
			AdditionalAudioURL     []string `json:"additionalAudioUrl,omitempty"` // FIXME This can also be just a single string.
			AmazonAlbumAsin        string   `json:"amazonAlbumAsin"`
			AmazonAlbumDigitalAsin string   `json:"amazonAlbumDigitalAsin"`
			ArtistExplorerURL      string   `json:"artistExplorerUrl"`
			SongName               string   `json:"songName"`
			AlbumDetailURL         string   `json:"albumDetailUrl"`
			SongDetailURL          string   `json:"songDetailUrl"`
			StationID              string   `json:"stationId"`
			SongRating             int      `json:"songRating"`
			TrackGain              string   `json:"trackGain"`
			AlbumExplorerURL       string   `json:"albumExplorerUrl"`
			AllowFeedback          bool     `json:"allowFeedback"`
			AmazonSongDigitalAsin  string   `json:"amazonSongDigitalAsin"`
			NowPlayingStationAdURL string   `json:"nowPlayingStationAdUrl"`
			AdToken                string   `json:"adToken"`
		} `json:"items"`
	} `json:"result"`
}

type TrackExplainTrack struct {
	Result struct {
		Explanations []struct {
			FocustTraitName string `json:"focusTraitName"`
			FocusTraitID    string `json:"focustTraitId"`
		} `json:"explanations"`
	} `json:"result"`
}
