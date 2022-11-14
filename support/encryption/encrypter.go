package encrypter

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/forgoer/openssl"
)

type Payload struct {
	Iv    string
	Value string
	Mac   string
}

var encrypter Encrypter

type Interface interface {
	Encrypt(text string, serialize bool) string
	Decrypt(payload string, serialize bool) string
}

type Encrypter struct {
	initiated bool
	key       string
}

/* Static */

func initialize() {
	if (Encrypter{}) == encrypter {
		encrypter = New()
	}
}

func Encrypt(text string, serialize bool) string {
	initialize()
	return encrypter.Encrypt(text, serialize)
}

func EncryptString(text string) string {
	initialize()
	return encrypter.EncryptString(text)
}

func Decrypt(payload string, serialize bool) string {
	initialize()
	return encrypter.Decrypt(payload, serialize)
}

func DecryptString(payload string) string {
	initialize()
	return encrypter.DecryptString(payload)
}

func SetKey(key string) string {
	initialize()
	encrypter.key = key
	return encrypter.key
}

func GetKey() string {
	initialize()
	return encrypter.key
}

func New() Encrypter {
	return Encrypter{
		initiated: true,
	}
}

/* Object */

func (encrypter Encrypter) hash(iv string, value string) string {
	key := []byte(encrypter.key)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(iv + value))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

func (encrypter Encrypter) calculateMac(payload Payload) string {
	return encrypter.hash(payload.Iv, payload.Value)
}

func (encrypter Encrypter) validMac(payload Payload) bool {
	expectedMAC := encrypter.calculateMac(payload)
	return hmac.Equal([]byte(expectedMAC), []byte(payload.Mac))
}

func (encrypter Encrypter) Encrypt(text string, serialize bool) string {
	iv := make([]byte, 16)
	_, err := rand.Read(iv)
	if err != nil {
		panic(err)
	}

	encrypted, err := openssl.AesCBCEncrypt([]byte(text), []byte(encrypter.key), iv, openssl.PKCS7_PADDING)
	if err != nil {
		panic(err)
	}

	encryptedBase64 := base64.StdEncoding.EncodeToString(encrypted)
	ivBase64 := base64.StdEncoding.EncodeToString(iv)

	mac := encrypter.hash(ivBase64, encryptedBase64)

	ticket := make(map[string]interface{})
	ticket["iv"] = ivBase64
	ticket["mac"] = mac
	ticket["value"] = encryptedBase64

	jsonTicket, err := json.Marshal(ticket)
	if err != nil {
		panic(err)
	}

	encodedTicket := base64.StdEncoding.EncodeToString(jsonTicket)

	return fmt.Sprintf("%s", string(encodedTicket))
}

func (encrypter Encrypter) EncryptString(text string) string {
	return encrypter.Encrypt(text, false)
}

func (encrypter Encrypter) Decrypt(payload string, serialize bool) string {
	ciphertext, err := base64.StdEncoding.DecodeString(payload)

	if err != nil {
		panic(err)
	}

	var decodedJson Payload

	json.Unmarshal([]byte(string(ciphertext)), &decodedJson)

	iv, _ := base64.StdEncoding.DecodeString(decodedJson.Iv)
	ciphertext, _ = base64.StdEncoding.DecodeString(decodedJson.Value)

	checkIsMacValid := encrypter.validMac(decodedJson)
	if !checkIsMacValid {
		panic("mac valid failed")
	}

	decrypted, err := openssl.AesCBCDecrypt(ciphertext, []byte(encrypter.key), iv, openssl.PKCS7_PADDING)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", string(decrypted))
}

func (encrypter Encrypter) DecryptString(payload string) string {
	return encrypter.Decrypt(payload, false)
}

func (encrypter *Encrypter) SetKey(key string) string {
	encrypter.key = key
	return encrypter.key
}

func (encrypter Encrypter) GetKey() string {
	return encrypter.key
}
