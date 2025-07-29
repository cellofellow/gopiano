package gopiano

import "testing"

func Test_Encrypt_1(t *testing.T) {
	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatal(err)
	}
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
	client, err := NewClient(AndroidClient)
	if err != nil {
		t.Fatal(err)
	}
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
