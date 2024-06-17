package usecases

import (
	"SimpleId/internal/model"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/services"
	"SimpleId/internal/utils"
	"SimpleId/internal/utils/encryption"
	"fmt"
	"log"
)

type UserUseCase interface {
	Login(input model.LoginUserRequest) (*model.LoginResponse, error)
	Register(input model.RegisterUserRequest) (*model.UserResponse, error)
	GetInformation(input model.UserRequest) (*model.UserInformation, error)
	AddInformation(input model.UserInformation) (*model.UserInformation, *model.TransactionResults, error)
	UpdateInformation(input model.UserInformation) (*model.UserInformation, *model.TransactionResults, error)
	Approve(input model.UserRequest) (*model.UserQuery, error)
}

type userUseCase struct {
	userService          services.UserRepository
	cryptoService        services.CryptoService
	smartContractService services.SmartContractService
}

func NewUserUseCase(userservice services.UserRepository, cryptoservices services.CryptoService, smartcontract services.SmartContractService) *userUseCase {
	return &userUseCase{
		userService:          userservice,
		cryptoService:        cryptoservices,
		smartContractService: smartcontract,
	}
}

func (usc *userUseCase) Login(input model.LoginUserRequest) (*model.LoginResponse, error) {
	results, err := usc.userService.LoginUser(input)
	if err != nil {
		return nil, fmt.Errorf("username atau password yang dimasukkan salah: %v", err)
	}
	return results, nil
}

func (usc *userUseCase) Register(input model.RegisterUserRequest) (*model.UserResponse, error) {
	results, err := usc.userService.CreateUser(input)
	if err != nil {
		return results, nil
	}
	return results, nil
}

// BUTUH CRYPTO SERVCIE
// BUTUH SMART CONTRACT SERVICE
func (usc *userUseCase) GetInformation(input model.UserRequest) (*model.UserInformation, error) {
	var user *model.UserQuery
	var err error

	if input.UID != "" {
		user, err = usc.userService.GetUserById(input.UID)
		if err != nil {
			return nil, fmt.Errorf("user not found by UID: %v", err)
		}
	} else if input.Username != "" {
		user, err = usc.userService.GetUserByUsername(input.Username)
		if err != nil {
			return nil, fmt.Errorf("user not found by Username: %v", err)
		}
	} else {
		return nil, fmt.Errorf("bad request: UID or Username must be provided")
	}

	encryptedUserData, err := usc.smartContractService.GetUserInformation(user.ContractAddress)
	if err != nil {
		return nil, err
	}

	secret, err := usc.cryptoService.DecryptString(user.Secret)
	if err != nil {
		return nil, err
	}
	userKey, userNonce, _ := encryption.SplitKey(secret)
	cryptoServiceForUser := services.SetKeyAndNonce(*userKey, *userNonce)
	results, err := cryptoServiceForUser.DecryptUserData(encryptedUserData)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

// BUTUH CRYPTO SERVCIE
// BUTUH SMART CONTRACT SERVICE
func (usc *userUseCase) AddInformation(input model.UserInformation) (*model.UserInformation, *model.TransactionResults, error) {
	var user *model.UserQuery
	var err error

	//Add info gak boleh pakai UID sebenarnya , bisa pakai username
	if input.UID != "" {
		user, err = usc.userService.GetUserById(input.UID)
		if err != nil {
			return nil, nil, fmt.Errorf("user not found by UID: %v", err)
		}
	} else {
		return nil, nil, fmt.Errorf("bad request: UID or Username must be provided")
	}

	secret, err := usc.cryptoService.DecryptString(user.Secret) //DECRYPT USER SECRET
	if err != nil {
		return nil, nil, err
	}

	userKey, userNonce, _ := encryption.SplitKey(secret)                  //SPLIT DECRYPTED SECRET
	cryptoServiceForUser := services.SetKeyAndNonce(*userKey, *userNonce) //SETUP USER KEY
	encryptedUserInfo, err := cryptoServiceForUser.EncryptUserData(input) //ENCRYPT USER INFO
	if err != nil {
		return nil, nil, err
	}

	results, err := usc.smartContractService.AddUserInformation(user.ContractAddress, encryptedUserInfo)
	if err != nil {
		return nil, nil, err
	}

	return &input, &results, nil
}

// BUTUH CRYPTO SERVCIE
// BUTUH SMART CONTRACT SERVICE
func (usc *userUseCase) UpdateInformation(input model.UserInformation) (*model.UserInformation, *model.TransactionResults, error) {
	var user *model.UserQuery
	var err error

	if input.UID != "" {
		user, err = usc.userService.GetUserById(input.UID)
		if err != nil {
			return nil, nil, fmt.Errorf("user not found by UID: %v", err)
		}
	} else {
		return nil, nil, fmt.Errorf("bad request: UID or Username must be provided")
	}

	secret, err := usc.cryptoService.DecryptString(user.Secret) //DECRYPT USER SECRET
	if err != nil {
		return nil, nil, err
	}

	userKey, userNonce, _ := encryption.SplitKey(secret)                  //SPLIT DECRYPTED SECRET
	cryptoServiceForUser := services.SetKeyAndNonce(*userKey, *userNonce) //SETUP USER KEY
	encryptedUserInfo, err := cryptoServiceForUser.EncryptUserData(input) //ENCRYPT USER INFO
	if err != nil {
		return nil, nil, err
	}

	results, err := usc.smartContractService.UpdateUserInformation(user.ContractAddress, encryptedUserInfo)
	if err != nil {
		return nil, nil, err
	}

	return &input, &results, nil
}

// BUTUH CRYPTO SERVCIE
// BUTUH SMART CONTRACT SERVICE
func (usc *userUseCase) Approve(input model.UserRequest) (*model.UserQuery, error) {

	user, err := usc.userService.GetUserById(input.UID)

	if err != nil {
		return nil, err
	}

	if user.Status == enum.Approved.String() {
		return nil, fmt.Errorf("bad request: User Already Approved")
	}

	generated, _ := encryption.GenerateRandom(44)                                         //GENERATE USER SECRET
	encrypted, err := usc.cryptoService.EncryptString(utils.ByteArrayToBase64(generated)) //ENCRPYT BASE64 USER SECRET
	if err != nil {
		return nil, err
	}
	log.Println("Trying to deploy smart contract....")

	deployed, err := usc.smartContractService.DeployIdentityContract()
	if err != nil {
		return nil, err
	}

	log.Println("Already deployed")
	fmt.Print(deployed)

	updateRequest := model.UpdateUserRequest{
		Status:          enum.Approved.String(),
		Secret:          encrypted,
		ContractAddress: deployed.ContractAddress,
	}
	results, err := usc.userService.UpdateUser(input.UID, updateRequest)
	if err != nil {
		return nil, err
	}
	return results, nil
}
