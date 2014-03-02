package gopiano

import (
	"bytes"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/cellofellow/gopiano/requests"
	"github.com/cellofellow/gopiano/responses"
)

func (c *Client) AuthPartnerLogin() (*responses.AuthPartnerLogin, error) {
	requestData := requests.AuthPartnerLogin{
		Username:    c.description["username"],
		Password:    c.description["password"],
		Version:     c.description["version"],
		DeviceModel: c.description["deviceModel"],
		IncludeURLs: true,
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.AuthPartnerLogin
	err = c.PandoraCall("https://", "auth.partnerLogin", requestDataReader, &resp)
	if err != nil {
		// TODO Handle error
		return nil, err
	}

	syncTime := c.decrypt(resp.Result.SyncTime)
	resp.Result.SyncTime = string(syncTime[4:14])
	i, err := strconv.ParseInt(resp.Result.SyncTime, 10, 32)
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}

	// Set partner data onto client for later use.
	c.timeOffset = time.Unix(i, 0).Sub(time.Now())
	c.partnerAuthToken = resp.Result.PartnerAuthToken
	c.partnerID = resp.Result.PartnerID

	return &resp, nil
}

func (c *Client) AuthUserLogin(username, password string) (*responses.AuthUserLogin, error) {
	requestData := requests.AuthUserLogin{
		PartnerAuthToken: c.partnerAuthToken,
		LoginType:        "user",
		Username:         username,
		Password:         password,
		SyncTime:         int(time.Now().Add(c.timeOffset).Unix()),
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.AuthUserLogin
	err = c.BlowfishCall("https://", "auth.userLogin", requestDataReader, &resp)
	if err != nil {
		// TODO Handle error
		return nil, err
	}

	// Set user data onto client for later use.
	c.userAuthToken = resp.Result.UserAuthToken
	c.userID = resp.Result.UserID

	return &resp, nil
}
