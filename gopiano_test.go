package gopiano

import "testing"

var client *Client = NewClient(AndroidClient)

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
	decrypted := client.decrypt(testString)
	if decrypted != expected {
		t.Error("decrypt failed.")
	} else {
		t.Log("decrypt passed")
	}
}

func Test_PartnerLogin_1(t *testing.T) {
	response, err := client.PartnerLogin()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}

func Test_UserLogin_1(t *testing.T) {
	response, err := client.UserLogin("mellowcellofellow@gmail.com", "Great8")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", response)
}
