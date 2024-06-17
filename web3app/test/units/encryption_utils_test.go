package testing

import (
	"SimpleId/internal/utils"
	"SimpleId/internal/utils/encryption"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEcnrypt(t *testing.T) {
	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("exampleplaintext")
	// nonce, _ := hex.DecodeString("bb8ef84243d2ee95a41c6c57")

	nonce, _ := encryption.GenerateRandom(12)
	fmt.Println(nonce)

	ciphertext, err := encryption.AesGCMEncrypt(plaintext, key, nonce)
	encodedBase64 := utils.ByteArrayToBase64(ciphertext)
	decodedBase64, _ := utils.Base64ToByteArray(encodedBase64)
	fmt.Println(encodedBase64)
	assert.NoError(t, err)
	assert.Equal(t, ciphertext, decodedBase64)
}

func TestDecrpy(t *testing.T) {
	expectedPlaintext := []byte("exampleplaintext")

	key := []byte("AES256Key-32Characters1234567890")
	nonce, _ := hex.DecodeString("bb8ef84243d2ee95a41c6c57")
	cipherTextByte, _ := utils.Base64ToByteArray("jkClBZ1UZANRTvOU9zKtXAuNACBT7lniqK77NDg+0Y0=")
	results, err := encryption.AesGCMDecrypt(cipherTextByte, key, nonce)
	assert.NoError(t, err)
	assert.Equal(t, expectedPlaintext, results)

}

func TestRandom(t *testing.T) {
	nonce, _ := encryption.GenerateRandom(12)
	fmt.Println("Nonce : " + utils.ByteArrayToBase64(nonce))

	key, _ := encryption.GenerateRandom(32)
	fmt.Println("Key : " + utils.ByteArrayToBase64(key))
}

func TestStringEncryption(t *testing.T) {
	expected := "fAph+VQ9QOT3bJTRPnl+FluaBqtwx7r1r/+aWf800yHy"
	plaintext := utils.StringToByteArray("testing plaintext")
	nonce, _ := utils.Base64ToByteArray("9iUqWZT0gayoe22w")
	key, _ := utils.Base64ToByteArray("KhZRw2bwF1I806uLQmNpmjstyq2NluoIyUNS8JltaW4=")

	results, err := encryption.AesGCMEncrypt(plaintext, key, nonce)
	if err != nil {
		fmt.Println("Terjadi kesalahan")
	}
	// fmt.Printf(utils.ByteArrayToBase64(results))
	assert.NoError(t, err)
	assert.Equal(t, expected, utils.ByteArrayToBase64(results))

}

func TestStringDecryption(t *testing.T) {
	expected := "testing plaintext"

	ciphertext, _ := utils.Base64ToByteArray("fAph+VQ9QOT3bJTRPnl+FluaBqtwx7r1r/+aWf800yHy")
	nonce, _ := utils.Base64ToByteArray("9iUqWZT0gayoe22w")
	key, _ := utils.Base64ToByteArray("KhZRw2bwF1I806uLQmNpmjstyq2NluoIyUNS8JltaW4=")

	results, err := encryption.AesGCMDecrypt(ciphertext, key, nonce)
	if err != nil {
		fmt.Println("Terjadi kesalahan")
	}
	fmt.Print(string(results))

	assert.NoError(t, err)
	assert.Equal(t, expected, string(results))
}

func TestSplitKey(t *testing.T) {
	// secretByte, _ := encryption.GenerateRandom(44)
	// secret := utils.ByteArrayToBase64(secretByte)
	secretByte, _ := utils.Base64ToByteArray("Pj1d1A4kc5fPRpTAImZKXBr3fdzcfCHo5e8w0zd/r98eJHEW/WvKGS2DamUXl/N1W9evd2f+oQypHnwy8A6QBTEPm3ead2R4rwFBQg==")
	secret := "Pj1d1A4kc5fPRpTAImZKXBr3fdzcfCHo5e8w0zd/r98eJHEW/WvKGS2DamUXl/N1W9evd2f+oQypHnwy8A6QBTEPm3ead2R4rwFBQg=="

	key, nonce, err := encryption.SplitKey(secret)
	assert.NoError(t, err)
	keyByte, err := utils.Base64ToByteArray(*key)
	assert.NoError(t, err)
	nonceByte, err := utils.Base64ToByteArray(*nonce)
	assert.NoError(t, err)

	combinedByte := append(keyByte, nonceByte...)
	assert.Equal(t, secretByte, combinedByte)

}
