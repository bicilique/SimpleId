package devenv

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/repository"
	"SimpleId/internal/services"
	"SimpleId/internal/smartcontracts/identityManagement"
	"SimpleId/internal/smartcontracts/secureDataTransfer"
	"SimpleId/internal/usecases"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestSharing(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	sharingUseCases := usecases.NewRequestUseCase(requestService, userService, cryptService, scService)
	dummy := model.SharingRequest{
		NIK:         "777888",
		Stakeholder: "PT. GHI",
	}

	excpected := model.RequestResponse{
		NIK:         "777888",
		Stakeholder: "PT. GHI",
	}

	results, err := sharingUseCases.RequestSharing(dummy)
	assert.NoError(t, err)
	fmt.Print(results)
	assert.Equal(t, excpected.NIK, results.NIK)
	assert.Equal(t, excpected.Stakeholder, results.Stakeholder)

}

func TestShowAll(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	sharingUseCases := usecases.NewRequestUseCase(requestService, userService, cryptService, scService)

	results, err := sharingUseCases.ShowAll()
	assert.NoError(t, err)
	fmt.Print(results)

	for _, result := range results {
		fmt.Println(result)
	}
}

func TestShowByStatus(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	sharingUseCases := usecases.NewRequestUseCase(requestService, userService, cryptService, scService)

	results, err := sharingUseCases.ShowAllByStatus(enum.Approved)
	assert.NoError(t, err)
	fmt.Print(results)

	for _, result := range results {
		fmt.Println(result)
	}
}

func TestApproveRequest(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	sharingUseCases := usecases.NewRequestUseCase(requestService, userService, cryptService, scService)

	receiverAddr := "0x0E98f547266BfdeF8aFA67BA15d5fc0700F6988c"
	input := model.GetSharingRequest{
		RequestID: "4883d29f-1cfb-4b14-adde-5587263e79ee",
		NIK:       "777888",
	}

	results, err := sharingUseCases.ApproveSharing(receiverAddr, input)
	assert.NoError(t, err)
	fmt.Print(results)
}

func TestReceiveInfo(t *testing.T) {
	db, client, cfg, err := setupUserUseCase()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	sharingUseCases := usecases.NewRequestUseCase(requestService, userService, cryptService, scService)

	conntractAddress := "0x7C3D2615969430aDeB263f751DBFF04a144DD1e4"
	kek := "i70+daYZR2i5lFo7r+/3TWKK6r7Y6g29/dO9H1o29gqpWTbj1PslNhEPhm4="
	results, err := sharingUseCases.GetSharingInformation(kek, conntractAddress)

	assert.NoError(t, err)
	fmt.Println(results)

	inkey := "Pj1d1A4kc5fPRpTAImZKXBr3fdzcfCHo5e8w0zd/r98eJHEW/WvKGS2DamUXl/N1W9evd2f+oQypHnwy8A6QBTEPm3ead2R4rwFBQg=="

	hasil, err := cryptService.DecryptString(inkey)
	assert.NoError(t, err)
	fmt.Println(hasil)

	assert.Equal(t, results, hasil)
}
