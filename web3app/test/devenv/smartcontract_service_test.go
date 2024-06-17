package devenv

import (
	"SimpleId/internal/config"
	"SimpleId/internal/model"
	"SimpleId/internal/services"
	"SimpleId/internal/smartcontracts/identityManagement"
	"SimpleId/internal/smartcontracts/secureDataTransfer"
	"SimpleId/internal/utils/blockchains"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func setupSmartContractEnv() (*ethclient.Client, string) {
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development") // Default to development
	}
	cfg := config.LoadConfig()

	client, err := blockchains.InitConnection(cfg.ClientPathNode1)
	if err != nil {
		log.Print("Failed load connection")
	}
	return client, cfg.WalletNode1
	// return client, "2bf5210b7621ba31b42aac698d2f81b4c5906c5790aaea443a8963be9886161a"
	// return client, "423f89ac4b6c9c049aae64c4cc673e477817b57e8de952d84fb349c122b49ac6"
	// return client, "30b3e5639f807b1803d928b5541cc0654ec1c108b31d08585277f83e59f17fa3"
}

func TestShowKey(t *testing.T) {
	// blockchains.GetPrivateKey("../../../node1/keystore/UTC--2024-05-18T02-08-14.416490000Z--70c726b420ac92095fa73b4ab298bc2fff7d9b0d", "../../example/Priv.txt", "1234")

	// blockchains.GetPrivateKey("../../../sealer1/keystore/UTC--2024-05-18T02-05-40.183973000Z--9374f2de64bda340f85a5c38a708a2a42889375f", "../../example/PrivSealer1.txt", "1qaz2wsx3edc")
	// blockchains.GetPrivateKey("../../../sealer2/keystore/UTC--2024-05-18T02-06-45.334895000Z--a592b54d0a7e28d1508a05ca58f167c950e54607", "../../example/PrivSealer2.txt", "0okm9ijn8uhb")
	// blockchains.GetPrivateKey("../../../sealer3/keystore/UTC--2024-05-18T02-07-42.072794000Z--c93eb278dd90c9d258f78e962cb4f11f8c3851fb", "../../example/PrivSealer3.txt", "1234qwerasdf")
	blockchains.GetPrivateKey("../../../node2/keystore/UTC--2024-05-18T02-08-47.530166000Z--0e98f547266bfdef8afa67ba15d5fc0700f6988c", "../../example/Node2.txt", "qwer")

}

func TestDeployIdentityMan(t *testing.T) {
	client, wallet := setupSmartContractEnv()
	idm := identityManagement.NewIdentitySmartContract(wallet, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(wallet, client)
	scService := services.NewSmartContractService(idm, sdc)
	results, err := scService.DeployIdentityContract()
	assert.NoError(t, err)
	fmt.Print(results)
	assert.NotNil(t, results.ContractAddress)
	assert.NotNil(t, results.TransactionAddress)

}

func TestAddUserInformation(t *testing.T) {
	contractAddress := "0x1EB121e2035f6b3e3EA5b3ba13fF48E83ADE93f1"
	client, wallet := setupSmartContractEnv()
	idm := identityManagement.NewIdentitySmartContract(wallet, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(wallet, client)
	scService := services.NewSmartContractService(idm, sdc)

	dummy := model.UserInformation{
		UID:       "123",
		Name:      "Momo",
		NIK:       "112233445566",
		Address:   "Depok",
		Country:   "Indonesia",
		Email:     "momo@gmail.com",
		BirthDate: "1997/07/26",
		Issuer:    "Dukcapil Depok",
		Active:    "Active",
	}

	results, err := scService.AddUserInformation(contractAddress, dummy)
	assert.NoError(t, err)
	fmt.Print(results)
	// assert.NotNil(t, results.Data)
	assert.NotNil(t, results.TransactionAddress)
}

func TestGetUserInformation(t *testing.T) {
	contractAddress := "0x71CEBDB465456b824Af2598e754F2FA283E13570"
	client, wallet := setupSmartContractEnv()
	idm := identityManagement.NewIdentitySmartContract(wallet, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(wallet, client)
	scService := services.NewSmartContractService(idm, sdc)

	expected := model.UserInformation{
		UID:       "123",
		Name:      "Momo",
		NIK:       "112233445566",
		Address:   "Depok",
		Country:   "Indonesia",
		Email:     "momo@gmail.com",
		BirthDate: "1997/07/26",
		Issuer:    "Dukcapil Depok",
		Active:    "Active",
	}

	_ = expected
	results, err := scService.GetUserInformation(contractAddress)
	assert.NoError(t, err)
	fmt.Print(results)
}

func TestAddStakeholder(t *testing.T) {

	contractAddress := "0x1EB121e2035f6b3e3EA5b3ba13fF48E83ADE93f1"
	client, wallet := setupSmartContractEnv()
	idm := identityManagement.NewIdentitySmartContract(wallet, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(wallet, client)
	scService := services.NewSmartContractService(idm, sdc)

	stakeholderPublickey := "0x70c726B420AC92095fa73B4ab298bC2fFF7d9B0d"

	results, err := scService.AddStakeholder(contractAddress, stakeholderPublickey)
	assert.NoError(t, err)
	fmt.Print(results)
}

func TestUpdateUserInfo(t *testing.T) {
	contractAddress := "0x1EB121e2035f6b3e3EA5b3ba13fF48E83ADE93f1"
	client, wallet := setupSmartContractEnv()
	idm := identityManagement.NewIdentitySmartContract(wallet, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(wallet, client)
	scService := services.NewSmartContractService(idm, sdc)

	dummy := model.UserInformation{
		UID:       "123",
		Name:      "Supanova",
		NIK:       "1234",
		Address:   "Depok",
		Country:   "Indonesia",
		Email:     "momo@gmail.com",
		BirthDate: "1997/07/26",
		Issuer:    "Dukcapil Depok",
		Active:    "Active",
	}

	results, err := scService.UpdateUserInformation(contractAddress, dummy)
	assert.NoError(t, err)
	fmt.Print(results)
}

func TestDeploySecureDataContract(t *testing.T) {
	client, wallet := setupSmartContractEnv()
	idm := identityManagement.NewIdentitySmartContract(wallet, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(wallet, client)
	scService := services.NewSmartContractService(idm, sdc)
	results, err := scService.DeploySecureDataContract()
	assert.NoError(t, err)
	fmt.Print(results)
	assert.NotNil(t, results.ContractAddress)
	assert.NotNil(t, results.TransactionAddress)
}

func TestSendData(t *testing.T) {
	client, wallet := setupSmartContractEnv()
	idm := identityManagement.NewIdentitySmartContract(wallet, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(wallet, client)
	scService := services.NewSmartContractService(idm, sdc)

	contractAddress := "0x8535F9cEBFd35E6E4e48502027Afed8d9B5744F8"

	receiver := common.HexToAddress("0x0E98f547266BfdeF8aFA67BA15d5fc0700F6988c") //NODE2

	input := model.DataToSend{
		Receiver:      receiver,
		EncryptedData: "Encrypted Data",
		DataHash:      "Hash Data",
		SignedHash:    "Signed hash data",
	}

	results, err := scService.SendData(contractAddress, input)
	assert.NoError(t, err)
	fmt.Print(results)
}

func TestReceiveData(t *testing.T) {
	client, wallet := setupSmartContractEnv()
	idm := identityManagement.NewIdentitySmartContract(wallet, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(wallet, client)
	scService := services.NewSmartContractService(idm, sdc)

	contractAddress := "0x8535F9cEBFd35E6E4e48502027Afed8d9B5744F8"

	results, err := scService.ReceiveData(contractAddress)
	assert.NoError(t, err)
	fmt.Print(results)
}
