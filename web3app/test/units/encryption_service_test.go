package testing

import (
	"SimpleId/internal/model"
	"SimpleId/internal/services"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var key = "KhZRw2bwF1I806uLQmNpmjstyq2NluoIyUNS8JltaW4="
var nonce = "9iUqWZT0gayoe22w"

func TestKeyGeneration(t *testing.T) {
	key, err := services.KeyGeneration(32)
	assert.NoError(t, err)

	fmt.Println(len(key))
	fmt.Println(key)
	nonce, err := services.KeyGeneration(12)
	assert.NoError(t, err)
	fmt.Println(len(nonce))
	fmt.Println(nonce)
	assert.Equal(t, 44, len(key))
	assert.Equal(t, 16, len(nonce))

}

func TestEncryptString(t *testing.T) {
	crypto_service := services.SetKeyAndNonce(key, nonce)
	input := "This is input string!"
	expected := "XAd7/h06VOTuboXNJC1oGl3af4qUT6covsAvPDlo1mVSGMqW7Q=="

	encrypted, err := crypto_service.EncryptString(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, encrypted)

}

func TestDecryptString(t *testing.T) {
	crypto_service := services.SetKeyAndNonce(key, nonce)
	input := "XAd7/h06VOTuboXNJC1oGl3af4qUT6covsAvPDlo1mVSGMqW7Q=="
	expected := "This is input string!"

	decrypted, err := crypto_service.DecryptString(input)
	assert.NoError(t, err)
	assert.Equal(t, expected, decrypted)
}

func TestEncryptUserData(t *testing.T) {
	crypto_service := services.SetKeyAndNonce(key, nonce)
	user := model.UserInformation{
		UID:       "123",
		Name:      "Momo",
		NIK:       "112233445566",
		Address:   "Depok",
		Country:   "Indonesia",
		Email:     "momo@gmail.com",
		BirthDate: "1997/07/26",
		Issuer:    "Dukcapil Depok",
		Active:    "Active",
	}

	expected := model.UserInformation{

		UID:       "OV0hf52NOpw/0FfLBg3w6cP5dQ==",
		Name:      "RQB/4oHOxwlbqUF16DNgMU0kr3c=",
		NIK:       "OV4gvw5gE/CyNcOOnHSd7AgOWEEZCNLBhmXHfA==",
		Address:   "TApi4lawAoi28awYUJGpFbf6qb8q",
		Country:   "QQF24lM2VK3m9c+53ZT+/e9n24vApRlJCg==",
		Email:     "ZQB/4n00SqXubNvbP2DLp6jxNl8WU3BiWMajnqiS",
		BirthDate: "OVYruhJjEOu1NtJ4s7LPhWUJi5DhFutEFV8=",
		Issuer:    "TBp57lwjTqinRJDIP2Yy9ambtqnv2tIVHYk8Mar/",
		Active:    "Active",
	}
	encrypted, err := crypto_service.EncryptUserData(user)
	assert.NoError(t, err)
	assert.Equal(t, expected, encrypted)
}

func TestDecryptUserData(t *testing.T) {
	crypto_service := services.SetKeyAndNonce(key, nonce)
	encrypted := model.UserInformation{
		UID:       "OV0hf52NOpw/0FfLBg3w6cP5dQ==",
		Name:      "RQB/4oHOxwlbqUF16DNgMU0kr3c=",
		NIK:       "OV4gvw5gE/CyNcOOnHSd7AgOWEEZCNLBhmXHfA==",
		Address:   "TApi4lawAoi28awYUJGpFbf6qb8q",
		Country:   "QQF24lM2VK3m9c+53ZT+/e9n24vApRlJCg==",
		Email:     "ZQB/4n00SqXubNvbP2DLp6jxNl8WU3BiWMajnqiS",
		BirthDate: "OVYruhJjEOu1NtJ4s7LPhWUJi5DhFutEFV8=",
		Issuer:    "TBp57lwjTqinRJDIP2Yy9ambtqnv2tIVHYk8Mar/",
		Active:    "Active",
	}

	expected := model.UserInformation{
		UID:       "123",
		Name:      "Momo",
		NIK:       "112233445566",
		Address:   "Depok",
		Country:   "Indonesia",
		Email:     "momo@gmail.com",
		BirthDate: "1997/07/26",
		Issuer:    "Dukcapil Depok",
		Active:    "Active",
	}
	user, err := crypto_service.DecryptUserData(encrypted)
	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}
