package gopiano

import (
	"bytes"
	"encoding/json"

	"github.com/cellofellow/gopiano/requests"
	"github.com/cellofellow/gopiano/responses"
)

func (c *Client) UserCanSubscribe() (*responses.UserCanSubscribe, error) {
	requestData := requests.UserCanSubscribe{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.UserCanSubscribe
	err = c.BlowfishCall("http://", "user.canSubscribe", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (c *Client) UserCreateUser(username, password, gender, countryCode string, zipCode, birthYear int, emailOptin bool) (*responses.UserCreateUser, error) {
	requestData := requests.UserCreateUser{
		PartnerAuthToken: c.partnerAuthToken,
		AccountType:      "registered",
		RegisteredType:   "user",
		Username:         username,
		Password:         password,
		Gender:           gender,
		ZipCode:          zipCode,
		CountryCode:      countryCode,
		BirthYear:        birthYear,
		EmailOptin:       emailOptin,
		SyncTime:         c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil,  err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.UserCreateUser
	err = c.BlowfishCall("https://", "user.createUser", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}

	// Set user data onto client for later use.
	c.userAuthToken = resp.Result.UserAuthToken
	c.userID = resp.Result.UserID

	return &resp, nil
}

func (c *Client) UserEmailPassword(username string) error {
	requestData := requests.UserEmailPassword{
		Username:         username,
		PartnerAuthToken: c.partnerAuthToken,
		SyncTime:         c.GetSyncTime(),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp interface{}
	return c.BlowfishCall("https://", "user.emailPassword", requestDataReader, &resp)
}

func (c *Client) UserGetBookmarks() (*responses.UserGetBookmarks, error) {
	requestData := requests.UserGetBookmarks{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
	}

	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.UserGetBookmarks
	err = c.BlowfishCall("http://", "user.getBookmarks", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UserGetStationList(includeStationArtURL bool) (*responses.UserGetStationList, error) {
	requestData := requests.UserGetStationList{
		UserAuthToken: c.userAuthToken,
		SyncTime:      c.GetSyncTime(),
		IncludeStationArtURL: includeStationArtURL,
	}

	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		return nil, err
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)

	var resp responses.UserGetStationList
	err = c.BlowfishCall("http://", "user.getStationList", requestDataReader, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
