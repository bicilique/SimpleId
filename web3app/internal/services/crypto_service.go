package services

import (
	"SimpleId/internal/model"
	"SimpleId/internal/utils"
	"SimpleId/internal/utils/encryption"
	"errors"
)

type CryptoService interface {
	EncryptString(input string) (string, error)
	DecryptString(input string) (string, error)
	EncryptUserData(input model.UserInformation) (model.UserInformation, error)
	DecryptUserData(input model.UserInformation) (model.UserInformation, error)
	// KeyGeneration(keyLength int) (string, error)
}

type cryptoService struct {
	key   []byte
	nonce []byte
}

func SetKeyAndNonce(key string, nonce string) *cryptoService {
	keyByte, err := utils.Base64ToByteArray(key)
	if err != nil {
		return nil
	}
	nonceByte, err := utils.Base64ToByteArray(nonce)
	if err != nil {
		return nil
	}

	return &cryptoService{
		key:   keyByte,
		nonce: nonceByte,
	}
}

func KeyGeneration(keyLength int) (string, error) {
	results, err := encryption.GenerateRandom(keyLength)
	if err != nil {
		return "", err
	}
	return utils.ByteArrayToBase64(results), nil
}

func (c *cryptoService) EncryptString(input string) (string, error) {
	results, err := encryption.AesGCMEncrypt(utils.StringToByteArray(input), c.key, c.nonce)
	if err != nil {
		return "", err
	}
	return utils.ByteArrayToBase64(results), nil
}

func (c *cryptoService) DecryptString(input string) (string, error) {
	ciphertext, err := utils.Base64ToByteArray(input)
	if err != nil {
		return "", err
	}

	results, err := encryption.AesGCMDecrypt(ciphertext, c.key, c.nonce)
	if err != nil {
		return "", err
	}
	return string(results), nil
}

func (c *cryptoService) EncryptUserData(input model.UserInformation) (model.UserInformation, error) {
	if (input == model.UserInformation{}) {
		return input, errors.ErrUnsupported
	}
	encryptedUser, err := encryption.AesGCMEncryptUser(input, c.key, c.nonce)
	if err != nil {
		return input, err
	}

	return encryptedUser, nil
}

func (c *cryptoService) DecryptUserData(input model.UserInformation) (model.UserInformation, error) {
	if (input == model.UserInformation{}) {
		return input, errors.ErrUnsupported
	}
	decryptedUser, err := encryption.AesGCMDecryptUser(input, c.key, c.nonce)
	if err != nil {
		return input, err
	}
	return decryptedUser, nil
}
