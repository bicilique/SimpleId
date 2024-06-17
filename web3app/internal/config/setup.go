package config

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/handlers"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/repository"
	"SimpleId/internal/services"
	"SimpleId/internal/smartcontracts/identityManagement"
	"SimpleId/internal/smartcontracts/secureDataTransfer"
	"SimpleId/internal/usecases"
	"SimpleId/internal/usecases/route"
	"SimpleId/internal/utils"
	"SimpleId/internal/utils/blockchains"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SetupApp() {
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development") // Default to development
	}
	cfg := LoadConfig()
	db := InitDb(cfg)

	client, err := blockchains.InitConnection(cfg.ClientPathNode1)
	if err != nil {
		log.Fatal(err.Error())
	}

	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	idm := identityManagement.NewIdentitySmartContract(cfg.WalletNode1, client)
	sdc := secureDataTransfer.NewSecureDataSmartContract(cfg.WalletNode1, client)
	scService := services.NewSmartContractService(idm, sdc)
	cryptService := services.SetKeyAndNonce(cfg.MKey, cfg.MNonce)

	sharingUseCases := usecases.NewRequestUseCase(requestService, userService, cryptService, scService)
	userUseCases := usecases.NewUserUseCase(userService, cryptService, scService)

	userHandler := handlers.NewUserHandler(userUseCases)
	adminHandler := handlers.NewAdminHandler(userUseCases, sharingUseCases)

	router := gin.Default()

	routeConfig := route.RouteConfig{
		Api:          router,
		UserHandler:  userHandler,
		AdminHandler: adminHandler,
	}
	routeConfig.Setup()
	seedData(db)

	router.Run()
}

func seedData(db *gorm.DB) {
	var roles = []entity.Role{
		{Name: "admin", Description: "Administrator role"},
		{Name: "customer", Description: "Authenticated customer role"},
		{Name: "anonymous", Description: "Unauthenticated customer role"},
	}
	db.Save(&roles)

	defaultPass, _ := utils.HashAndEncodePassword("qwerty")
	roles[0].ID = 1
	var admin = entity.User{
		UID:      uuid.New().String(),
		NIK:      "0000111100001111",
		Username: "admin",
		Name:     "admin",
		Password: defaultPass,
		Status:   enum.Approved.String(),
		RoleID:   1,
		Role:     roles[0],
	}

	db.Save(&admin)

}
