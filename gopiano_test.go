package gopiano

import "testing"

var client *Client

func Test_Setup(t *testing.T) {
	client, _ = NewClient(AndroidClient)
}

func Test_Encrypt_1(t *testing.T) {
	testString := "foobar"
	expected := "3c739d4e29b5d6c6"
	encrypted := client.encrypt(testString)
	if encrypted != expected {
		t.Error("encrypt failed.")
	} else {
		t.Log("encrypt passed")
	}
}

func Test_Decrypt_1(t *testing.T) {
	expected := "foobar"
	testString := "95b6027f2d427dc0"
	decrypted, err := client.decrypt(testString)
	if err != nil {
		t.Error(err)
	}
	if decrypted != expected {
		t.Error("decrypt failed.")
	} else {
		t.Log("decrypt passed")
	}
}

func Test_AuthPartnerLogin_1(t *testing.T) {
	response, err := client.AuthPartnerLogin()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_AuthUserLogin_1(t *testing.T) {
	response, err := client.AuthUserLogin("mellowcellofellow@gmail.com", "Great8")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserCanSubscribe_1(t *testing.T) {
	response, err := client.UserCanSubscribe()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserBetBookmarks_1(t *testing.T) {
	response, err := client.UserGetBookmarks()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserGetStationList_1(t *testing.T) {
	response, err := client.UserGetStationList(true)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserGetStationListChecksum_1(t *testing.T) {
	response, err := client.UserGetStationListChecksum()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}
