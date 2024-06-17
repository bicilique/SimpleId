package utils

import "encoding/base64"

func ByteArrayToBase64(input []byte) string {
	reults := base64.StdEncoding.EncodeToString([]byte(input))
	return reults
}

func Base64ToByteArray(input string) ([]byte, error) {
	results, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return nil, err
	}
	return results, nil
}

func StringToByteArray(input string) []byte {
	return []byte(input)
}

func ByteArrayToString(byteArray []byte) string {
	return string(byteArray)
}
