package gopiano

import (
	"bytes"
	"code.google.com/p/go.crypto/blowfish"
	"encoding/hex"
	"encoding/json"
	"github.com/cellofellow/gopiano/requests"
	"github.com/cellofellow/gopiano/responses"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type ClientDescription map[string]string

var AndroidClient = ClientDescription{
	"deviceModel": "android-generic",
	"username":    "android",
	"password":    "AC7IBG09A3DTSYM4R41UJWL07VLN8JI7",
	"baseUrl":     "tuner.pandora.com/services/json/",
	"encryptKey":  "6#26FRL$ZWD",
	"decryptKey":  "R=U!LH$O2B#",
	"version":     "5",
}

type Client struct {
	description      ClientDescription
	http             *http.Client
	encrypter        *blowfish.Cipher
	decrypter        *blowfish.Cipher
	timeOffset       time.Duration
	partnerAuthToken string
	partnerID        string
	userAuthToken    string
	userID           string
}

func NewClient(d ClientDescription) *Client {
	client := &http.Client{}
	encrypter, err := blowfish.NewCipher([]byte(d["encryptKey"]))
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}
	decrypter, err := blowfish.NewCipher([]byte(d["decryptKey"]))
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}
	return &Client{
		description: d,
		http:        client,
		encrypter:   encrypter,
		decrypter:   decrypter,
	}
}

func (c *Client) encrypt(data string) string {
	chunks := make([]string, 0)
	for i := 0; i < len(data); i += 8 {
		var buf [8]byte
		var crypt [8]byte
		copy(buf[:], data[i:])
		c.encrypter.Encrypt(crypt[:], buf[:])
		encoded := hex.EncodeToString(crypt[:])
		chunks = append(chunks, encoded)
	}
	return strings.Join(chunks, "")
}

func (c *Client) decrypt(data string) string {
	chunks := make([]string, 0)
	for i := 0; i < len(data); i += 16 {
		var buf [16]byte
		var decoded, decrypted [8]byte
		copy(buf[:], data[i:])
		_, err := hex.Decode(decoded[:], buf[:])
		if err != nil {
			// TODO Handle error
			log.Fatal(err)
		}
		c.decrypter.Decrypt(decrypted[:], decoded[:])
		chunks = append(chunks, strings.Trim(string(decrypted[:]), "\x00"))
	}
	return strings.Join(chunks, "")
}

func (c *Client) PandoraCall(protocol string, method string, body io.Reader, data interface{}) error {
	urlArgs := url.Values{
		"method": {method},
	}

	if c.partnerID != "" {
		urlArgs.Add("partner_id", c.partnerID)
	}
	if c.userID != "" {
		urlArgs.Add("user_id", c.userID)
	}
	if c.userAuthToken != "" && c.partnerAuthToken == "" {
		urlArgs.Add("auth_token", c.userAuthToken)
	} else if c.partnerAuthToken != "" && c.userAuthToken == "" {
		urlArgs.Add("auth_token", c.partnerAuthToken)
	}
	callUrl := protocol + c.description["baseUrl"] + "?" + urlArgs.Encode()

	//// Clone of actual request for debugging.
	//bodyBytes, err := ioutil.ReadAll(body)
	//debugBody := strings.NewReader(string(bodyBytes))
	//debugRequest, err := http.NewRequest("POST", callUrl, debugBody)
	//if err != nil {
	//	// TODO Handle error.
	//	log.Fatal(err)
	//}
	//debugRequest.Header.Add("User-Agent", "pithos")
	//debugRequest.Header.Add("Content-type", "text/plain")
	//debugRequest.Write(os.Stderr)
	//body = strings.NewReader(string(bodyBytes))

	req, err := http.NewRequest("POST", callUrl, body)
	if err != nil {
		// TODO Handle error.
		log.Fatal(err)
	}
	req.Header.Add("User-Agent", "gopiano")
	req.Header.Add("Content-type", "text/plain")

	resp, err := c.http.Do(req)
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}

	var errResp responses.ErrorResponse
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(responseBody, &errResp)
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}

	if errResp.Stat == "fail" {
		return errResp
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}
	return nil
}

func (c *Client) BlowfishCall(protocol string, method string, body io.Reader, data interface{}) error {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}
	encrypted := strings.NewReader(c.encrypt(string(bodyBytes)))
	return c.PandoraCall(protocol, method, encrypted, data)
}

func (c *Client) PartnerLogin() (*responses.PartnerLogin, error) {
	requestData := requests.PartnerLogin{
		Username:    c.description["username"],
		Password:    c.description["password"],
		Version:     c.description["version"],
		DeviceModel: c.description["deviceModel"],
		IncludeUrls: true,
	}
	requestDataEncoded, err := json.Marshal(requestData)
	if err != nil {
		// TODO Handle error
		log.Fatal(err)
	}
	requestDataReader := bytes.NewReader(requestDataEncoded)
	var resp responses.PartnerLogin
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

func (c *Client) UserLogin(username, password string) (*responses.UserLogin, error) {
	requestData := requests.UserLogin{
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
	var resp responses.UserLogin
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
