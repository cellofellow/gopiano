//go:build integration

package gopiano

import "testing"

var client *Client

func Test_Setup(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}
	client, _ = NewClient(AndroidClient)
}

func Test_AuthPartnerLogin_1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}
	response, err := client.AuthPartnerLogin()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_AuthUserLogin_1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}
	response, err := client.AuthUserLogin("mellowcellofellow@gmail.com", "Great8")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserCanSubscribe_1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}
	response, err := client.UserCanSubscribe()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserBetBookmarks_1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}
	response, err := client.UserGetBookmarks()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserGetStationList_1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}
	response, err := client.UserGetStationList(true)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserGetStationListChecksum_1(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping network test in short mode")
	}
	response, err := client.UserGetStationListChecksum()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}
