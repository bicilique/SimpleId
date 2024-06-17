package devenv

import (
	"SimpleId/internal/config"
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/repository"
	"SimpleId/internal/services"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupEnv() (*gorm.DB, error) {
	if os.Getenv("APP_ENV") == "" {
		os.Setenv("APP_ENV", "development") // Default to development
	}
	cfg := config.LoadConfig()

	// fmt.Println("DB Host:", cfg.DBHost)
	// fmt.Println("DB Port:", cfg.DBPort)
	// fmt.Println("DB User:", cfg.DBUser)
	// fmt.Println("DB Password:", cfg.DBPassword)
	// fmt.Println("DB Name:", cfg.DBName)

	username := cfg.DBUser
	password := cfg.DBPassword
	dbHost := "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" //"@tcp(localhost:3306)/"
	dbName := cfg.DBName
	dsn := username + ":" + password + dbHost + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&entity.User{})
	return db, nil
}

func TestSetupRepositort(t *testing.T) {
	db, err := setupEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)
	_ = userService
}

func TestCreateUser(t *testing.T) {
	db, err := setupEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	dummy := model.RegisterUserRequest{
		NIK:      "123456",
		Name:     "John Doe",
		Username: "johndoe",
		Password: "123",
	}

	expected := model.UserResponse{
		NIK:  "123456",
		Name: "John Doe",
	}
	results, err := userService.CreateUser(dummy)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Name, results.Name)
	assert.Equal(t, enum.Waiting.String(), results.Status)
}

func TestCreateUser2(t *testing.T) {
	db, err := setupEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	dummy := model.RegisterUserRequest{
		NIK:      "98765",
		Name:     "Afif Faizianur",
		Username: "faiz",
		Password: "456",
	}

	expected := model.UserResponse{
		NIK:  "98765",
		Name: "Afif Faizianur",
	}
	results, err := userService.CreateUser(dummy)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Name, results.Name)
	assert.Equal(t, enum.Waiting.String(), results.Status)
}

func TestLogin(t *testing.T) {
	db, err := setupEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	dummy := model.LoginUserRequest{
		Username: "johndoe",
		Password: "123",
	}

	expected := model.UserResponse{
		NIK:  "123456",
		Name: "John Doe",
	}

	results, err := userService.LoginUser(dummy)
	assert.NoError(t, err)
	_ = expected
	_ = results
}

func TestShowAllUser(t *testing.T) {
	db, err := setupEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	expectedUserResponses := []model.UserResponse{
		{NIK: "123456", Name: "John Doe"},
		{NIK: "98765", Name: "Afif Faizianur"},
	}

	results, err := userService.ShowAllUser()
	assert.NoError(t, err)

	for i, expected := range expectedUserResponses {
		assert.Equal(t, expected.NIK, results[i].NIK)
		assert.Equal(t, expected.Name, results[i].Name)
	}
}

func TestGetUserById(t *testing.T) {
	db, err := setupEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	dummyUID := "0890b8eb-84ae-4abf-b71e-1901ba7056c0"
	expected := model.UserQuery{
		UID:  "0890b8eb-84ae-4abf-b71e-1901ba7056c0",
		NIK:  "98765",
		Name: "Afif Faizianur",
	}

	results, err := userService.GetUserById(dummyUID)
	assert.NoError(t, err)
	assert.Equal(t, expected.UID, results.UID)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Name, results.Name)
}

func TestGetUserByUsername(t *testing.T) {
	db, err := setupEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	dummyUname := "faiz"
	expected := model.UserQuery{
		UID:  "0890b8eb-84ae-4abf-b71e-1901ba7056c0",
		NIK:  "98765",
		Name: "Afif Faizianur",
	}

	results, err := userService.GetUserByUsername(dummyUname)
	assert.NoError(t, err)

	fmt.Print("User UID : " + results.UID)
	assert.Equal(t, expected.UID, results.UID)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Name, results.Name)
}

func TestUpdateUser(t *testing.T) {
	db, err := setupEnv()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	userRepo := repository.NewRepository[entity.User](db)
	userService := services.NewUserRepository(userRepo)

	uid := "bd9d5b57-6167-473a-9a91-91ae63a0cea1"

	dummy := model.UpdateUserRequest{
		Name:            "John Doex",
		Password:        "123",
		Secret:          "secret",
		ContractAddress: "contract",
		Status:          enum.Approved.String(),
	}

	expected := model.UserQuery{
		NIK:             "123456",
		Name:            "John Doex",
		Secret:          "secret",
		ContractAddress: "contract",
		Status:          enum.Approved.String(),
	}

	results, err := userService.UpdateUser(uid, dummy)
	assert.NoError(t, err)

	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Name, results.Name)
	assert.Equal(t, expected.Secret, results.Secret)
	assert.Equal(t, expected.Status, results.Status)
	assert.Equal(t, expected.ContractAddress, results.ContractAddress)

}
