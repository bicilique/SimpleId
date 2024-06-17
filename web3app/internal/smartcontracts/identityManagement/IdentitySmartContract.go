package identityManagement

import (
	"SimpleId/internal/model"
	"SimpleId/internal/utils/blockchains"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type IdentitySmartContract interface {
	DeployIdentityContract() (model.DeployResponse, error)
	LoadIdentityContract(contractAddress string) bool
	AddUserInformation(input model.UserInformation) (model.TransactionResults, error)
	UpdateUserInformation(input model.UserInformation) (model.TransactionResults, error)
	GetUserInformation() (model.UserInformation, error)
	AddStakeholder(stakeholderAddress string) (model.TransactionResults, error)
}

type identitySmartContract struct {
	Client   *ethclient.Client
	Config   *bind.TransactOpts
	Instance *IdentityManagement
}

func NewIdentitySmartContract(walletPrivateKey string, client *ethclient.Client) *identitySmartContract {
	publicKey, privKey := blockchains.ImportWallet(walletPrivateKey)
	config := blockchains.PrepareTransaction(publicKey, privKey, client)
	return &identitySmartContract{
		Config: config,
		Client: client,
	}
}

func (idsc *identitySmartContract) DeployIdentityContract() (model.DeployResponse, error) {
	contractAddress, transcactionAddress, instance, err := DeployIdentityManagement(idsc.Config, idsc.Client)
	if err != nil {
		return model.DeployResponse{}, err
	}
	idsc.Instance = instance
	return model.DeployResponse{
		ContractAddress:    contractAddress.Hex(),
		TransactionAddress: transcactionAddress.Hash().Hex(),
	}, nil
}

func (idsc *identitySmartContract) LoadIdentityContract(contractAddress string) bool {
	identityContract, err := NewIdentityManagement(common.HexToAddress(contractAddress), idsc.Client)
	if err != nil {
		return false
	}
	idsc.Instance = identityContract
	return true
}

func (idsc *identitySmartContract) AddUserInformation(input model.UserInformation) (model.TransactionResults, error) {
	tx, err := idsc.Instance.CreateIdentity(idsc.Config,
		input.Name,
		input.NIK,
		input.Email,
		input.Address,
		input.Country,
		input.BirthDate,
		input.Status,
		true,
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

func (idsc *identitySmartContract) UpdateUserInformation(input model.UserInformation) (model.TransactionResults, error) {
	tx, err := idsc.Instance.UpdateIdentity(idsc.Config,
		input.Name,
		input.NIK,
		input.Email,
		input.Address,
		input.Country,
		input.BirthDate,
		input.Status,
		true,
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

func (idsc *identitySmartContract) GetUserInformation() (model.UserInformation, error) {
	callOpts := &bind.CallOpts{
		Pending: true,
		From:    idsc.Config.From,
	}
	results, err := idsc.Instance.GetIdentity(callOpts)
	if err != nil {
		return model.UserInformation{}, err
	}

	isActive := ""
	if results.Active {
		isActive = "true"
	} else {
		isActive = "false"
	}

	return model.UserInformation{
		Name:      results.Name,
		NIK:       results.Nik,
		Address:   results.Street,
		Country:   results.Country,
		Email:     results.Email,
		BirthDate: results.Birthdate,
		Issuer:    results.Issuer.Hex(),
		Status:    results.Status,
		Active:    isActive,
	}, nil
}

func (idsc *identitySmartContract) AddStakeholder(stakeholderAddress string) (model.TransactionResults, error) {
	tx, err := idsc.Instance.AddStakeholder(idsc.Config, common.HexToAddress(stakeholderAddress))
	if err != nil {
		return model.TransactionResults{}, err
	}

	return model.TransactionResults{
		TransactionAddress: tx.Hash().Hex(),
		Data:               tx.Data(),
		Status:             true,
	}, nil
}
