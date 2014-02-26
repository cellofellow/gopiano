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
		PartnerId        string `json:"partnerId"`
		StationSkipUnit  string `json:"stationSkipUnit"`
		DeviceProperties struct {
			VideoAdRefreshInterval int `json:"videoAdRefreshInterval"`
			VideoAdUniqueInterval  int `json:"videoAdUniqueInterval"`
			AdRefreshInterval      int `json:"adRefreshInterval"`
			VideoAdStartInterval   int `json:"videoAdStartInterval"`
		} `json:"deviceProperties"`
		Urls             struct {
			AutoComplete string `json:"autoComplete"`
		} `json:"urls"`
	}
}
