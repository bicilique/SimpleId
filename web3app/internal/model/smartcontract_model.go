package model

import "math/big"

type DeployResponse struct {
	ContractAddress    string `json:"contract,omitempty"`
	TransactionAddress string `json:"transaction,omitempty"`
}

type TransactionResults struct {
	TransactionAddress string `json:"transaction,omitempty"`
	Data               any
	Status             bool
}

type DataResults struct {
	EncryptedData string   `json:"encryptedData,omitempty"`
	DataHash      string   `json:"dataHash,omitempty"`
	SignedHash    string   `json:"signedHash,omitempty"`
	TimeCreated   *big.Int `json:"timeCreated,omitempty"`
}

// type AddUserInfo struct {
// 	UID             string `json:"uid,omitempty" validate:"max=160"`
// 	Name            string `json:"name,omitempty" validate:"max=160"`
// 	NIK             string `json:"nik,omitempty" validate:"max=160"`
// 	Address         string `json:"address,omitempty" validate:"max=160"`
// 	Country         string `json:"country,omitempty" validate:"max=160"`
// 	Email           string `json:"email,omitempty" validate:"max=160"`
// 	BirthDate       string `json:"birthdate,omitempty" validate:"max=160"`
// 	Issuer          string `json:"issuer,omitempty" validate:"max=160"`
// 	Status          string `json:"status,omitempty" validate:"max=160"`
// 	Active          string `json:"active,omitempty" validate:"max=160"`
// 	ContractAddress string `json:"contractaddress,omitempty" validate:"max=160"`
// }
