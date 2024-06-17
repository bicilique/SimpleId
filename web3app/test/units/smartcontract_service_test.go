package testing

// import (
// 	"SimpleId/internal/model"
// 	"SimpleId/internal/services"
// 	"SimpleId/test/mock_identityManagement"
// 	"SimpleId/test/mock_secureDataTransfer"
// 	"fmt"
// 	"testing"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// // var client = &ethclient.Client{}
// // var walletKey = "fef922214731e0086fcbaa19b98519f2ef2a857a3c2663c9daae2456d9290d5c"

// func setupTest(ctrl *gomock.Controller) (*mock_identityManagement.MockIdentitySmartContract, *mock_secureDataTransfer.MockSecureDataSmartContract) {
// 	idm := mock_identityManagement.NewMockIdentitySmartContract(ctrl)
// 	sdc := mock_secureDataTransfer.NewMockSecureDataSmartContract(ctrl)
// 	return idm, sdc
// }

// func TestDeployIdentityManagement(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	idm, scd := setupTest(ctrl)
// 	sc_service := services.NewSmartContractService(idm, scd)

// 	dummy := model.DeployResponse{
// 		ContractAddress:    "0xAABBCCDD",
// 		TransactionAddress: "OxFFGGHH",
// 	}

// 	expected := model.DeployResponse{
// 		ContractAddress:    "0xAABBCCDD",
// 		TransactionAddress: "OxFFGGHH",
// 	}

// 	idm.EXPECT().DeployIdentityContract().Return(dummy, nil)
// 	results, err := sc_service.DeployIdentityContract()
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected.ContractAddress, results.ContractAddress)
// 	assert.Equal(t, expected.TransactionAddress, results.TransactionAddress)
// }

// func TestAddUserInformation(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	idm, scd := setupTest(ctrl)
// 	sc_service := services.NewSmartContractService(idm, scd)

// 	contractAddrress := "0xCCBBDD"

// 	userInput := model.UserInformation{
// 		UID:       "123",
// 		Name:      "Momo",
// 		NIK:       "112233445566",
// 		Address:   "Depok",
// 		Country:   "Indonesia",
// 		Email:     "momo@gmail.com",
// 		BirthDate: "1997/07/26",
// 		Issuer:    "Dukcapil Depok",
// 		Active:    "Active",
// 	}

// 	dummy := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}
// 	expected := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}

// 	idm.EXPECT().LoadIdentityContract(contractAddrress).Return(true)
// 	idm.EXPECT().AddUserInformation(gomock.Any()).Return(dummy, nil)
// 	results, err := sc_service.AddUserInformation(contractAddrress, userInput)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected.TransactionAddress, results.TransactionAddress)
// 	assert.Equal(t, expected.Data, results.Data)
// 	assert.Equal(t, expected.Status, results.Status)
// }

// func TestUpdateUserInformation(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	idm, scd := setupTest(ctrl)
// 	sc_service := services.NewSmartContractService(idm, scd)

// 	contractAddrress := "0xCCBBDD"

// 	userInput := model.UserInformation{
// 		UID:       "123",
// 		Name:      "Momo",
// 		NIK:       "112233445566",
// 		Address:   "Depok",
// 		Country:   "Indonesia",
// 		Email:     "momo@gmail.com",
// 		BirthDate: "1997/07/26",
// 		Issuer:    "Dukcapil Depok",
// 		Active:    "Active",
// 	}

// 	dummy := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}
// 	expected := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}

// 	idm.EXPECT().LoadIdentityContract(contractAddrress).Return(true)
// 	idm.EXPECT().UpdateUserInformation(gomock.Any()).Return(dummy, nil)
// 	results, err := sc_service.UpdateUserInformation(contractAddrress, userInput)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected.TransactionAddress, results.TransactionAddress)
// 	assert.Equal(t, expected.Data, results.Data)
// 	assert.Equal(t, expected.Status, results.Status)
// }

// func TestGetUserInformation(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	idm, scd := setupTest(ctrl)
// 	sc_service := services.NewSmartContractService(idm, scd)

// 	contractAddrress := "0xCCBBDD"

// 	dummy := model.UserInformation{
// 		UID:       "123",
// 		Name:      "Momo",
// 		NIK:       "112233445566",
// 		Address:   "Depok",
// 		Country:   "Indonesia",
// 		Email:     "momo@gmail.com",
// 		BirthDate: "1997/07/26",
// 		Issuer:    "Dukcapil Depok",
// 		Active:    "Active",
// 	}

// 	expected := model.UserInformation{
// 		UID:       "123",
// 		Name:      "Momo",
// 		NIK:       "112233445566",
// 		Address:   "Depok",
// 		Country:   "Indonesia",
// 		Email:     "momo@gmail.com",
// 		BirthDate: "1997/07/26",
// 		Issuer:    "Dukcapil Depok",
// 		Active:    "Active",
// 	}

// 	idm.EXPECT().LoadIdentityContract(contractAddrress).Return(true)
// 	idm.EXPECT().GetUserInformation().Return(dummy, nil)
// 	results, err := sc_service.GetUserInformation(contractAddrress)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected.Name, results.Name)
// 	assert.Equal(t, expected.NIK, results.NIK)
// 	assert.Equal(t, expected.NIK, results.NIK)
// 	assert.Equal(t, expected.Address, results.Address)
// 	assert.Equal(t, expected.Country, results.Country)
// 	assert.Equal(t, expected.Email, results.Email)
// 	assert.Equal(t, expected.BirthDate, results.BirthDate)
// 	assert.Equal(t, expected.Issuer, results.Issuer)
// 	assert.Equal(t, expected.Active, results.Active)
// }

// func TestAddStakeholder(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	idm, scd := setupTest(ctrl)
// 	sc_service := services.NewSmartContractService(idm, scd)

// 	contractAddrress := "0xCCBBDD"

// 	userInput := "0xAAACCC"

// 	dummy := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}
// 	expected := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}

// 	idm.EXPECT().LoadIdentityContract(contractAddrress).Return(true)
// 	idm.EXPECT().AddStakeholder(gomock.Any()).Return(dummy, nil)
// 	results, err := sc_service.AddStakeholder(contractAddrress, userInput)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected.TransactionAddress, results.TransactionAddress)
// 	assert.Equal(t, expected.Data, results.Data)
// 	assert.Equal(t, expected.Status, results.Status)
// }

// func TestDeploySecureDataContract(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	idm, scd := setupTest(ctrl)
// 	sc_service := services.NewSmartContractService(idm, scd)

// 	dummy := model.DeployResponse{
// 		ContractAddress:    "0xAABBCCDD",
// 		TransactionAddress: "OxFFGGHH",
// 	}

// 	expected := model.DeployResponse{
// 		ContractAddress:    "0xAABBCCDD",
// 		TransactionAddress: "OxFFGGHH",
// 	}

// 	scd.EXPECT().DeploySecureDataContract().Return(dummy, nil)
// 	results, err := sc_service.DeploySecureDataContract()
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected.ContractAddress, results.ContractAddress)
// 	assert.Equal(t, expected.TransactionAddress, results.TransactionAddress)
// }

// func TestSendData(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	idm, scd := setupTest(ctrl)
// 	sc_service := services.NewSmartContractService(idm, scd)

// 	contractAddrress := "0xCCBBDD"

// 	receiver := common.HexToAddress("0xAABBCCDDEEFF")

// 	input := model.DataToSend{
// 		Receiver:      receiver,
// 		EncryptedData: "Encrypted Data",
// 		DataHash:      "Hash Data",
// 		SignedHash:    "Signed hash data",
// 	}

// 	dummy := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}
// 	expected := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}

// 	scd.EXPECT().LoadSecureDataContract(contractAddrress).Return(true)
// 	scd.EXPECT().SendData(gomock.Any()).Return(dummy, nil)

// 	results, err := sc_service.SendData(contractAddrress, input)
// 	assert.NoError(t, err)
// 	assert.Equal(t, expected.TransactionAddress, results.TransactionAddress)
// 	assert.Equal(t, expected.Data, results.Data)
// 	assert.Equal(t, expected.Status, results.Status)
// }

// func TestReceiveData(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	idm, scd := setupTest(ctrl)
// 	sc_service := services.NewSmartContractService(idm, scd)

// 	contractAddrress := "0xCCBBDD"

// 	dummy := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}
// 	expected := model.TransactionResults{
// 		TransactionAddress: "0XAAFFCC",
// 		Data:               "AAAA",
// 		Status:             true,
// 	}

// 	scd.EXPECT().LoadSecureDataContract(contractAddrress).Return(true)
// 	scd.EXPECT().ReceiveData().Return(dummy, nil)

// 	results, err := sc_service.ReceiveData(contractAddrress)
// 	assert.NoError(t, err)
// 	fmt.Print(results)
// 	// assert.Equal(t, expected.TransactionAddress, results.TransactionAddress)
// 	// assert.Equal(t, expected.Data, results.Data)
// 	// assert.Equal(t, expected.Status, results.Status)
// }
