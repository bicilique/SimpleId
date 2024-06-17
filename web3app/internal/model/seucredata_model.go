package model

import "github.com/ethereum/go-ethereum/common"

type DataToSend struct {
	Receiver      common.Address `json:"receiver,omitempty"`
	EncryptedData string         `json:"encryptedData,omitempty"`
	DataHash      string         `json:"dataHash,omitempty"`
	SignedHash    string         `json:"signedHash,omitempty"`
}
