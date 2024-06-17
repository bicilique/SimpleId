package devenv

import (
	"SimpleId/internal/config"
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/repository"
	"SimpleId/internal/services"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupEnvRequest() (*gorm.DB, error) {
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
		return nil, err
	}

	db.AutoMigrate(&entity.Request{})
	return db, nil
}

func TestSetupRequestRepositort(t *testing.T) {
	db, err := setupEnvRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)
	_ = requestService
}

func TestCreateNewRequest(t *testing.T) {

	db, err := setupEnvRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	dummyRequest := model.SharingRequest{
		NIK:         "123",
		Stakeholder: "PT. XYZ",
	}

	expected := model.RequestResponse{
		NIK:         "123",
		Stakeholder: "PT. XYZ",
	}

	results, err := requestService.CreateNewRequest(dummyRequest)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Stakeholder, results.Stakeholder)

}

func TestCreateNewRequest2(t *testing.T) {

	db, err := setupEnvRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	dummyRequest := model.SharingRequest{
		NIK:         "456",
		Stakeholder: "PT. JDH",
	}

	expected := model.RequestResponse{
		NIK:         "456",
		Stakeholder: "PT. JDH",
	}

	results, err := requestService.CreateNewRequest(dummyRequest)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Stakeholder, results.Stakeholder)

}

func TestGetRequestById(t *testing.T) {
	db, err := setupEnvRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	request := model.GetSharingRequest{
		RequestID: "12763b53-c873-4b26-a382-472a2790e9e3",
		NIK:       "123",
	}

	expected := model.RequestResponse{
		NIK:         "123",
		Stakeholder: "PT. XYZ",
	}

	results, err := requestService.GetRequestById(request)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Stakeholder, results.Stakeholder)

}

func TestGetRequestByNIK(t *testing.T) {
	db, err := setupEnvRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	request := model.GetSharingRequest{
		RequestID: "12763b53-c873-4b26-a382-472a2790e9e3",
		NIK:       "123",
	}

	expected := model.RequestResponse{
		NIK:         "123",
		Stakeholder: "PT. XYZ",
	}

	results, err := requestService.GetRequestByNIK(request)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Stakeholder, results.Stakeholder)
}

func TestShowAllRequest(t *testing.T) {
	db, err := setupEnvRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	expectedUserResponses := []model.RequestResponse{
		{NIK: "123",
			Stakeholder: "PT. XYZ"},
		{NIK: "456",
			Stakeholder: "PT. JDH"},
	}

	results, err := requestService.ShowAllRequest()
	assert.NoError(t, err)

	for i, expected := range expectedUserResponses {
		assert.Equal(t, expected.NIK, results[i].NIK)
		assert.Equal(t, expected.Stakeholder, results[i].Stakeholder)
	}
}

func TestUpdateSharingRequest(t *testing.T) {
	db, err := setupEnvRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	dummyRequest := model.UpdateSharingRequest{
		RequestID:   "12763b53-c873-4b26-a382-472a2790e9e3",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		Status:      enum.Waiting.String(),
	}

	expected := model.UpdateSharingRequest{
		RequestID:   "12763b53-c873-4b26-a382-472a2790e9e3",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		Status:      enum.Waiting.String(),
	}
	results, err := requestService.UpdateSharingRequest(dummyRequest)
	assert.NoError(t, err)
	assert.Equal(t, expected.NIK, results.NIK)
	assert.Equal(t, expected.Stakeholder, results.Stakeholder)

}

func TestShowFilteredRequest(t *testing.T) {
	db, err := setupEnvRequest()
	if err != nil {
		log.Fatal(err.Error())
	}
	assert.NoError(t, err)
	requestRepo := repository.NewRepository[entity.Request](db)
	requestService := services.NewRequestRepository(requestRepo)

	expectedUserResponses := []model.RequestResponse{
		{NIK: "123",
			Stakeholder: "PT. XYZ"},
		{NIK: "456",
			Stakeholder: "PT. JDH"},
	}

	results, err := requestService.ShowFilteredRequest(enum.Waiting)
	assert.NoError(t, err)

	for i, expected := range expectedUserResponses {
		assert.Equal(t, expected.NIK, results[i].NIK)
		assert.Equal(t, expected.Stakeholder, results[i].Stakeholder)
	}
}
