package secureDataTransfer

import (
	"SimpleId/internal/model"
	"SimpleId/internal/utils/blockchains"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SecureDataSmartContract interface {
	DeploySecureDataContract() (model.DeployResponse, error)
	LoadSecureDataContract(contractAddress string) bool
	SendData(input model.DataToSend) (model.TransactionResults, error)
	ReceiveData() (model.DataResults, error)
}

type secureDataSmartContract struct {
	Client   *ethclient.Client
	Config   *bind.TransactOpts
	Instance *SecureDataTransfer
}

func NewSecureDataSmartContract(walletPrivateKey string, client *ethclient.Client) *secureDataSmartContract {
	publicKey, privKey := blockchains.ImportWallet(walletPrivateKey)
	config := blockchains.PrepareTransaction(publicKey, privKey, client)
	return &secureDataSmartContract{
		Config: config,
		Client: client,
	}
}

func (sdsc *secureDataSmartContract) DeploySecureDataContract() (model.DeployResponse, error) {
	contractAddress, transcactionAddress, instance, err := DeploySecureDataTransfer(sdsc.Config, sdsc.Client)
	if err != nil {
		return model.DeployResponse{}, err
	}
	sdsc.Instance = instance
	return model.DeployResponse{
		ContractAddress:    contractAddress.Hex(),
		TransactionAddress: transcactionAddress.Hash().Hex(),
	}, nil
}

func (sdsc *secureDataSmartContract) SendData(input model.DataToSend) (model.TransactionResults, error) {
	tx, err := sdsc.Instance.SendData(sdsc.Config,
		input.Receiver,
		input.EncryptedData,
		input.DataHash,
		input.SignedHash,
	)

	if err != nil {
		return model.TransactionResults{}, err
	}

	return model.TransactionResults{
		TransactionAddress: tx.Hash().Hex(),
		Data:               tx.Data(),
		Status:             true,
	}, nil
}

func (sdsc *secureDataSmartContract) ReceiveData() (model.DataResults, error) {
	callOpts := &bind.CallOpts{
		Pending: true,
		From:    sdsc.Config.From,
	}
	results, err := sdsc.Instance.GetData(callOpts)
	if err != nil {
		log.Fatalf("An error occurred while reading data from smart contract : %v", err)
		return model.DataResults{}, err
	}

	return model.DataResults{
		EncryptedData: results.EncryptedData,
	}, nil
}

func (sdsc *secureDataSmartContract) LoadSecureDataContract(contractAddress string) bool {
	secureDataContract, err := NewSecureDataTransfer(common.HexToAddress(contractAddress), sdsc.Client)
	if err != nil {
		return false
	}
	sdsc.Instance = secureDataContract
	return true
}
