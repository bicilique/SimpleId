package usecases

import (
	"SimpleId/internal/model"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/services"
	"SimpleId/internal/utils"
	"SimpleId/internal/utils/encryption"

	"github.com/ethereum/go-ethereum/common"
)

type SharingUseCase interface {
	RequestSharing(input model.SharingRequest) (*model.RequestResponse, error)
	ApproveSharing(receiverAddress string, input model.GetSharingRequest) (*model.RequestResponse, error)
	GetRequest(input model.GetSharingRequest) (*model.RequestResponse, error)
	ShowAll() ([]*model.RequestResponse, error)
	ShowAllByStatus(status enum.Status) ([]*model.RequestResponse, error)
	GetSharingInformation(inputSecret string, contractAddress string) (string, error)
}

type sharingUseCase struct {
	userService          services.UserRepository
	requestService       services.RequestRepository
	cryptoService        services.CryptoService
	smartContractService services.SmartContractService
}

func NewRequestUseCase(requestService services.RequestRepository, userService services.UserRepository, cryptoService services.CryptoService, smartContractService services.SmartContractService) *sharingUseCase {
	return &sharingUseCase{
		requestService:       requestService,
		userService:          userService,
		cryptoService:        cryptoService,
		smartContractService: smartContractService,
	}
}

func (ssc *sharingUseCase) RequestSharing(input model.SharingRequest) (*model.RequestResponse, error) {
	results, err := ssc.requestService.CreateNewRequest(input)
	if err != nil {
		return results, nil
	}
	return results, nil
}

func (ssc *sharingUseCase) ApproveSharing(receiverAddress string, input model.GetSharingRequest) (*model.RequestResponse, error) {
	user, err := ssc.userService.GetUserByNIK(input.NIK)
	if err != nil {
		return nil, err
	}

	requested, err := ssc.requestService.GetRequestByNIK(input)
	if err != nil {
		return nil, err
	}

	//DECRYPT USER KEY WITH MASTER KEY
	userKey, err := ssc.cryptoService.DecryptString(user.Secret)
	if err != nil {
		return nil, err
	}

	//GENERATE KEK
	generated, err := services.KeyGeneration(44)
	if err != nil {
		return nil, err
	}

	kek, nonce, _ := encryption.SplitKey(generated)
	encryptService := services.SetKeyAndNonce(*kek, *nonce)

	//ENCRYPT USER KET WITH KEK
	encryptedUserKey, err := encryptService.EncryptString(userKey)
	if err != nil {
		return nil, err
	}

	filePath, _ := utils.GetProjectRoot()
	utils.WriteStringToFile(generated, filePath+"/web3app/example/generatedKey/"+user.UID)

	deployResults, err := ssc.smartContractService.DeploySecureDataContract()
	if err != nil {
		return nil, err
	}
	receiver := common.HexToAddress(receiverAddress)

	dataToSend := model.DataToSend{
		Receiver:      receiver,
		EncryptedData: encryptedUserKey,
		DataHash:      "DATAHASH",
		SignedHash:    "SIGNEDHASH",
	}

	sendingResults, err := ssc.smartContractService.SendData(deployResults.ContractAddress, dataToSend)
	if err != nil {
		return nil, err
	}

	_ = sendingResults // INI BUG

	utils.WriteStringToFile(deployResults.ContractAddress, filePath+"/web3app/example/tx/"+deployResults.TransactionAddress)

	updatedData := model.UpdateSharingRequest{
		RequestID:       requested.RequestID,
		Status:          enum.Approved.String(),
		ContractAddress: deployResults.ContractAddress,
	}

	results, err := ssc.requestService.UpdateSharingRequest(updatedData)
	if err != nil {
		return nil, err
	}

	return results, nil

	//SIGN KEK WITH PUBLIC KEY SEMENTARA GAK DLU

	//DEPLOY SC
	//SEND INFO TO SC
	//UPDATE DATA IN DB
	// ssc.requestService.UpdateSharingRequest()
}

func (ssc *sharingUseCase) GetRequest(input model.GetSharingRequest) (*model.RequestResponse, error) {
	results, err := ssc.requestService.GetRequestById(input)
	if err != nil {
		return results, nil
	}
	return results, nil
}

func (ssc *sharingUseCase) ShowAll() ([]*model.RequestResponse, error) {
	results, err := ssc.requestService.ShowAllRequest()
	if err != nil {
		return results, nil
	}
	return results, nil
}

func (ssc *sharingUseCase) ShowAllByStatus(status enum.Status) ([]*model.RequestResponse, error) {
	results, err := ssc.requestService.ShowFilteredRequest(status)
	if err != nil {
		return results, nil
	}
	return results, nil
}

func (ssc *sharingUseCase) GetSharingInformation(inputSecret string, contractAddress string) (string, error) {
	sharingData, err := ssc.smartContractService.ReceiveData(contractAddress)
	if err != nil {
		return "", err
	}

	key, nonce, err := encryption.SplitKey(inputSecret)
	if err != nil {
		return "", err
	}

	cryptService := services.SetKeyAndNonce(*key, *nonce)
	results, err := cryptService.DecryptString(sharingData.EncryptedData)
	if err != nil {
		return "", err
	}

	return results, nil

}
