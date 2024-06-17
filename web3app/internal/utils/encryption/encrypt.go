package encryption

import (
	"SimpleId/internal/model"
	"SimpleId/internal/utils"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func AesGCMEncrypt(plaintext []byte, key []byte, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func AesGCMDecrypt(ciphertext []byte, key []byte, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}

func GenerateRandom(size int) ([]byte, error) {
	nonce := make([]byte, size)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}

func AesGCMEncryptUser(input model.UserInformation, key []byte, nonce []byte) (model.UserInformation, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return model.UserInformation{}, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return model.UserInformation{}, err
	}

	uid := ""
	if input.Name != "" {
		uid = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.UID), nil))
	}

	name := ""
	if input.Name != "" {
		name = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.Name), nil))
	}

	nik := ""
	if input.NIK != "" {
		nik = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.NIK), nil))
	}

	address := ""
	if input.Address != "" {
		address = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.Address), nil))
	}

	county := ""
	if input.Country != "" {
		county = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.Country), nil))
	}

	email := ""
	if input.Email != "" {
		email = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.Email), nil))
	}

	birthdate := ""
	if input.BirthDate != "" {
		birthdate = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.BirthDate), nil))
	}

	issuer := ""
	if input.Issuer != "" {
		issuer = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.Issuer), nil))
	}

	status := ""
	if input.Status != "" {
		status = utils.ByteArrayToBase64(aesgcm.Seal(nil, nonce, utils.StringToByteArray(input.Status), nil))
	}

	return model.UserInformation{
		UID:       uid,
		Name:      name,
		NIK:       nik,
		Address:   address,
		Country:   county,
		Email:     email,
		BirthDate: birthdate,
		Issuer:    issuer,
		Status:    status,
		Active:    input.Active,
	}, nil
}

func aesGCMDecrypString(input string, aesgcm cipher.AEAD, nonce []byte) (string, error) {
	inputByte, err := utils.Base64ToByteArray(input)
	if err != nil {
		return "", err
	}

	decrypted, err := aesgcm.Open(nil, nonce, inputByte, nil)
	if err != nil {
		return "", err
	}
	return string(decrypted), err

}

func AesGCMDecryptUser(input model.UserInformation, key []byte, nonce []byte) (model.UserInformation, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return model.UserInformation{}, err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return model.UserInformation{}, err
	}

	uid := ""
	if input.UID != "" {
		uid, _ = aesGCMDecrypString(input.UID, aesgcm, nonce)
	}

	name := ""
	if input.Name != "" {
		name, _ = aesGCMDecrypString(input.Name, aesgcm, nonce)
	}

	nik := ""
	if input.Name != "" {
		nik, _ = aesGCMDecrypString(input.NIK, aesgcm, nonce)
	}

	address := ""
	if input.Name != "" {
		address, _ = aesGCMDecrypString(input.Address, aesgcm, nonce)
	}

	county := ""
	if input.Name != "" {
		county, _ = aesGCMDecrypString(input.Country, aesgcm, nonce)
	}

	email := ""
	if input.Name != "" {
		email, _ = aesGCMDecrypString(input.Email, aesgcm, nonce)
	}

	birthdate := ""
	if input.Name != "" {
		birthdate, _ = aesGCMDecrypString(input.BirthDate, aesgcm, nonce)
	}

	issuer := ""
	if input.Name != "" {
		issuer, _ = aesGCMDecrypString(input.Issuer, aesgcm, nonce)
	}

	status := ""
	if input.Status != "" {
		status, _ = aesGCMDecrypString(input.Status, aesgcm, nonce)
	}

	return model.UserInformation{
		UID:       uid,
		Name:      name,
		NIK:       nik,
		Address:   address,
		Country:   county,
		Email:     email,
		BirthDate: birthdate,
		Issuer:    issuer,
		Status:    status,
		Active:    input.Active,
	}, nil

}

func SplitKey(input string) (*string, *string, error) {
	secretByte, err := utils.Base64ToByteArray(input)
	if err != nil {
		return nil, nil, err
	}

	key := utils.ByteArrayToBase64(secretByte[:32])
	nonce := utils.ByteArrayToBase64(secretByte[32:])
	return &key, &nonce, nil
}
