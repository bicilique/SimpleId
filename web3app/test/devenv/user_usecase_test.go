package devenv

import (
	"SimpleId/internal/config"
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/repository"
	"SimpleId/internal/services"
	"SimpleId/internal/smartcontracts/identityManagement"
	"SimpleId/internal/smartcontracts/secureDataTransfer"
	"SimpleId/internal/usecases"
	"SimpleId/internal/utils/blockchains"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupUserUseCase() (*gorm.DB, *ethclient.Client, *config.Config, error) {
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development") // Default to development
	}
	cfg := config.LoadConfig()

	username := cfg.DBUser
	password := cfg.DBPassword
	dbHost := "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" //"@tcp(localhost:3306)/"
	dbName := cfg.DBName
	dsn := username + ":" + password + dbHost + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, nil, err
	}

	db.AutoMigrate(
		&entity.User{},
		&entity.Request{},
		&entity.Shared{},
	)

	client, err := blockchains.InitConnection(cfg.ClientPathNode1)
	if err != nil {
		log.Print("tidak dapat melakukan koneksi")
		return nil, nil, nil, err
	}

	return db, client, cfg, nil
}

func TestRegister(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)
	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	userUseCases := usecases.NewUserUseCase(userService, cryptService, scService)

	dummy := model.RegisterUserRequest{
		NIK:      "777888",
		Name:     "Ruan Mei",
		Username: "ruanmei11",
		Password: "qwerty",
	}

	expected := model.UserResponse{
		NIK:  "777888",
		Name: "Ruan Mei",
	}
	results, err := userUseCases.Register(dummy)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Name, results.Name)
	assert.Equal(t, enum.Waiting.String(), results.Status)
}

func TestLoginUseCase(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)
	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	userUseCases := usecases.NewUserUseCase(userService, cryptService, scService)

	dummy := model.LoginUserRequest{
		Username: "ruanmei11",
		Password: "qwerty",
	}

	expected := model.UserResponse{
		NIK:  "777888",
		Name: "Ruan Mei",
	}

	results, err := userUseCases.Login(dummy)
	assert.NoError(t, err)
	_ = expected
	_ = results
}

func TestApprove(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)
	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	userUseCases := usecases.NewUserUseCase(userService, cryptService, scService)

	dummy := model.UserRequest{
		UID: "eef9a0cb-6b22-490c-a937-de8b8694a507",
	}

	expected := model.UserQuery{
		NIK:  "777888",
		Name: "Ruan Mei",
	}

	results, err := userUseCases.Approve(dummy)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Name, results.Name)
	log.Default().Print(results.ContractAddress)
	log.Default().Print(results.Secret)
}

func TestAddInformation(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)
	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	userUseCases := usecases.NewUserUseCase(userService, cryptService, scService)

	dummy := model.UserInformation{
		UID:     "eef9a0cb-6b22-490c-a937-de8b8694a507",
		Name:    "Ruan Mei",
		NIK:     "777888",
		Address: "Djakarta",
		Country: "Belgium",
		Email:   "ruan@mei.com",
		Issuer:  "PT. Indo Solusi",
		Status:  enum.Approved.String(),
		Active:  "true",
	}

	results, tx, err := userUseCases.AddInformation(dummy)
	assert.NoError(t, err)
	fmt.Print(results)
	fmt.Print(tx)
}

func TestGetInformation(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)
	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	userUseCases := usecases.NewUserUseCase(userService, cryptService, scService)

	dummy := model.UserRequest{
		UID: "c4eec187-4ffa-4c22-9002-ada2a06b5f9f",
	}

	results, err := userUseCases.GetInformation(dummy)
	assert.NoError(t, err)
	fmt.Print(results)
}

func TestUpdateInformation(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)
	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	userUseCases := usecases.NewUserUseCase(userService, cryptService, scService)

	dummy := model.UserInformation{
		UID:     "c4eec187-4ffa-4c22-9002-ada2a06b5f9f",
		Name:    "Fuan Mei",
		NIK:     "777888",
		Address: "Djakarta",
		Country: "Belgium",
		Email:   "ruan@mei.com",
		Issuer:  "PT. Indo Solusi",
		Status:  enum.Approved.String(),
		Active:  "true",
	}

	results, tx, err := userUseCases.UpdateInformation(dummy)
	assert.NoError(t, err)
	fmt.Print(tx)
	fmt.Print(results)

}
