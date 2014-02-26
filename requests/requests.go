package requests

type PartnerLogin struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DeviceModel string `json:"deviceModel"`
	Version     string `json:"version"`
	IncludeUrls bool   `json:"includeUrls,omitempty"`
}

type UserLogin struct {
	IncludeAdAttributes           bool   `json:"includeAdAttributes,omitempty"`
	IncludeDemographics           bool   `json:"IncludeDemographics,omitempty"`
	IncludePandoraOneInfo         bool   `json:"includePandoraOneInfo,omitempty"` // Appears to do nothing.
	IncludeStationArtUrl          bool   `json:"includeStationArtUrl,omitempty"`
	IncludeSubscriptionExpiration bool   `json:"includeSubscriptionExpiration,omitempty"`
	LoginType                     string `json:"loginType"`
	PartnerAuthToken              string `json:"partnerAuthToken"`
	Password                      string `json:"password"`
	ReturnCapped                  bool   `json:"returnCapped,omitempty"`
	ReturnGenreStations           bool   `json:"returnGenreStations,omitempty"`
	ReturnStationList             bool   `json:"returnStationList,omitempty"`
	SyncTime                      int    `json:"syncTime"`
	Username                      string `json:"username"`
}
