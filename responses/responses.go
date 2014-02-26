package responses

import "fmt"

type ErrorResponse struct {
	Stat    string `json:"stat"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ErrorResponse) Error() string {
	return fmt.Sprintf("Pandora Error: %d %s", e.Code, e.Message)
}

type PartnerLogin struct {
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

type UserLogin struct {
	Result struct {
		CanListen                   bool   `json:"canListen"`
		HasAudioAds                 bool   `json:"hasAudioAds"`
		IsCapped                    bool   `json:"isCapped,omitempty"`
		ListeningTimeoutAlertMsgUri string `json:"listeningTimeoutAlertMsgUri"`
		ListeningTimeoutMinutes     string `json:"listeningTimeoutMinutes"`
		MaxStationsAllowed          int    `json:"maxStationsAllowed"`
		MinimumAdRefreshInterval    int    `json:"minimumAdRefreshInterval"`
		NowPlayingUrl               string `json:"nowPlayingUrl"`
		SplashScreenAdUrl           string `json:"splashScreenAdUrl"`
		StationCreationAdUrl        string `json:"stationCreationAdUrl"`
		UserAuthToken               string `json:"userAuthToken"`
		UserID                      string `json:"userId"`
		UserProfileUrl              string `json:"userProfileUrl"`
		Username                    string `json:"username"`
		VideoAdUrl                  string `json:"videoAdUrl"`
	} `json:"result"`
}
