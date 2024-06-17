package services

import (
	"SimpleId/internal/model"
	"SimpleId/internal/smartcontracts/identityManagement"
	"SimpleId/internal/smartcontracts/secureDataTransfer"
)

type SmartContractService interface {
	DeployIdentityContract() (model.DeployResponse, error)
	AddUserInformation(contractAddress string, input model.UserInformation) (model.TransactionResults, error)
	UpdateUserInformation(contractAddress string, input model.UserInformation) (model.TransactionResults, error)
	GetUserInformation(contractAddress string) (model.UserInformation, error)
	AddStakeholder(contractAddress string, stakeholderAddress string) (model.TransactionResults, error)

	DeploySecureDataContract() (model.DeployResponse, error)
	SendData(contractAddress string, input model.DataToSend) (model.TransactionResults, error)
	ReceiveData(contractAddress string) (model.DataResults, error)
}

type smartContractService struct {
	identityContract   identityManagement.IdentitySmartContract
	secureDataContract secureDataTransfer.SecureDataSmartContract
}

func NewSmartContractService(idm identityManagement.IdentitySmartContract, scd secureDataTransfer.SecureDataSmartContract) *smartContractService {
	// idm := identityManagement.NewIdentitySmartContract(walletPrivateKey, client)
	// scd := secureDataTransfer.NewSecureDataSmartContract(walletPrivateKey, client)
	return &smartContractService{
		identityContract:   idm,
		secureDataContract: scd,
	}
}

func (smc *smartContractService) loadIdentityContract(contractAddress string) bool {
	reults := smc.identityContract.LoadIdentityContract(contractAddress)
	return reults
}

func (smc *smartContractService) loadSecureDataContract(contractAddress string) bool {
	results := smc.secureDataContract.LoadSecureDataContract(contractAddress)
	return results
}

func (smc *smartContractService) DeployIdentityContract() (model.DeployResponse, error) {
	results, err := smc.identityContract.DeployIdentityContract()
	if err != nil {
		return model.DeployResponse{}, err
	}
	return results, nil
}

func (smc *smartContractService) AddUserInformation(contractAddress string, input model.UserInformation) (model.TransactionResults, error) {
	smc.loadIdentityContract(contractAddress) //TODO Butuh custom error
	results, err := smc.identityContract.AddUserInformation(input)
	if err != nil {
		return model.TransactionResults{}, err
	}
	return results, nil
}

func (smc *smartContractService) UpdateUserInformation(contractAddress string, input model.UserInformation) (model.TransactionResults, error) {
	smc.loadIdentityContract(contractAddress) //TODO Butuh custom error
	results, err := smc.identityContract.UpdateUserInformation(input)
	if err != nil {
		return model.TransactionResults{}, err
	}
	return results, nil
}

func (smc *smartContractService) GetUserInformation(contractAddress string) (model.UserInformation, error) {
	smc.loadIdentityContract(contractAddress) //TODO Butuh custom error
	results, err := smc.identityContract.GetUserInformation()
	if err != nil {
		return model.UserInformation{}, err
	}
	return results, nil
}

func (smc *smartContractService) AddStakeholder(contractAddress string, stakeholderAddress string) (model.TransactionResults, error) {
	smc.loadIdentityContract(contractAddress) //TODO Butuh custom error
	results, err := smc.identityContract.AddStakeholder(stakeholderAddress)
	if err != nil {
		return model.TransactionResults{}, err
	}
	return results, nil
}

func (smc *smartContractService) DeploySecureDataContract() (model.DeployResponse, error) {
	results, err := smc.secureDataContract.DeploySecureDataContract()
	if err != nil {
		return model.DeployResponse{}, err
	}
	return results, nil
}

func (smc *smartContractService) SendData(contractAddress string, input model.DataToSend) (model.TransactionResults, error) {
	smc.loadSecureDataContract(contractAddress)
	results, err := smc.secureDataContract.SendData(input)
	if err != nil {
		return model.TransactionResults{}, err
	}
	return results, nil
}

func (smc *smartContractService) ReceiveData(contractAddress string) (model.DataResults, error) {
	smc.loadSecureDataContract(contractAddress)
	results, err := smc.secureDataContract.ReceiveData()
	if err != nil {
		return model.DataResults{}, err
	}
	return results, nil
}
