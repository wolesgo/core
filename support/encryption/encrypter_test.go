package encrypter

import (
	"testing"
)

func TestDecrypt(t *testing.T) {
	encrypter := New()

	encrypter.SetKey(DummyKey())

	decryptedText := "FooBar"
	encrypted := "eyJpdiI6IkVRM3htMnlPanlRbUlKZ2k0N2pMdkE9PSIsIm1hYyI6ImMxYjhmOTUyZDkwNWY0NzliNDgwMzcyMWI2NmY5NGIwM2IwMmY4OThmODMzNzNiMGUwN2MyZjhlZWVlMWJkYTMiLCJ2YWx1ZSI6IjhRQUhkRjVZSkIraW15OWppVEhFWkE9PSJ9"

	decrypted := encrypter.DecryptString(encrypted)

	if decrypted != decryptedText {
		t.Errorf("Decrypted value must be %s", decryptedText)
	}
}

func TestEncryption(t *testing.T) {
	encrypter := New()

	encrypter.SetKey(DummyKey())

	testText := "FooBar"

	encrypted := encrypter.EncryptString(testText)

	decrypted := encrypter.DecryptString(encrypted)

	if decrypted != testText {
		t.Errorf("Decrypted value must be %s", testText)
	}
}

func DummyKey() string {
	return "GReveBscqKSlgaUqKhNhGdDeZXVrOjZh"
}
