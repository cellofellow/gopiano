package requests

type PartnerLogin struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	DeviceModel string `json:"deviceModel"`
	Version     string `json:"version"`
	IncludeUrls bool   `json:"includeUrls,omitempty"`
}
