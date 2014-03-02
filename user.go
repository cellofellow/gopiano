package gopiano

import (
	"bytes"
	"encoding/json"
	"log"

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
		// TODO Handle error
		log.Fatal(err)
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.UserCanSubscribe
	err = c.BlowfishCall("http://", "user.canSubscribe", requestDataReader, &resp)
	if err != nil {
		// TODO Handle error
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
		// TODO Handle error
		log.Fatal(err)
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.UserCreateUser
	err = c.BlowfishCall("https://", "user.createUser", requestDataReader, &resp)
	if err != nil {
		// TODO Handle error
		return nil, err
	}

	// Set user data onto client for later use.
	c.userAuthToken = resp.Result.UserAuthToken
	c.userID = resp.Result.UserID

	return &resp, nil
}
