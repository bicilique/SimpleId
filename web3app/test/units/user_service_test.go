package testing

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
	"SimpleId/internal/services"
	"SimpleId/test/mock_repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var clockie = time.Now()

func TestShowAllUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo := mock_repository.NewMockRepository[entity.User](ctrl)
	user_services := services.NewUserRepository(mock_repo)

	// Define dummy users
	dummyUsers := []entity.User{
		{UID: "1", NIK: "123", Name: "John Doe", ContractAddress: "0x123", CreatedAt: clockie, Status: "active", UpdatedAt: clockie},
		{UID: "2", NIK: "456", Name: "Jane Doe", ContractAddress: "0x456", CreatedAt: clockie, Status: "inactive", UpdatedAt: clockie},
	}

	// Define expected user responses
	expectedUserResponses := []model.UserQuery{
		{UID: "1", NIK: "123", Name: "John Doe", ContractAddress: "0x123", CreatedAt: clockie, Status: "active", UpdatedAt: clockie},
		{UID: "2", NIK: "456", Name: "Jane Doe", ContractAddress: "0x456", CreatedAt: clockie, Status: "inactive", UpdatedAt: clockie},
	}

	mock_repo.EXPECT().SelectAll().Return(dummyUsers, nil)
	userResponses, err := user_services.ShowAllUser()
	assert.NoError(t, err)
	for i, expected := range expectedUserResponses {
		assert.Equal(t, expected.UID, userResponses[i].UID)
		assert.Equal(t, expected.NIK, userResponses[i].NIK)
		assert.Equal(t, expected.Name, userResponses[i].Name)
		assert.Equal(t, expected.ContractAddress, userResponses[i].ContractAddress)
		assert.Equal(t, expected.CreatedAt, userResponses[i].CreatedAt)
		assert.Equal(t, expected.Status, userResponses[i].Status)
		assert.Equal(t, expected.UpdatedAt, userResponses[i].UpdatedAt)
	}

}

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo := mock_repository.NewMockRepository[entity.User](ctrl)
	user_services := services.NewUserRepository(mock_repo)

	userRequest := model.RegisterUserRequest{
		NIK:      "123456",
		Name:     "John Doe",
		Username: "johndoe",
		Password: "hashed_password123",
	}

	expectedUser := model.UserQuery{
		UID:             "1",
		NIK:             "123",
		Name:            "John Doe",
		ContractAddress: "0x123",
		CreatedAt:       clockie,
		Status:          "active",
		UpdatedAt:       clockie,
	}

	dummyUser := entity.User{
		UID:             "1",
		NIK:             "123",
		Name:            "John Doe",
		ContractAddress: "0x123",
		CreatedAt:       clockie,
		Status:          "active",
		UpdatedAt:       clockie,
	}

	mock_repo.EXPECT().Insert(gomock.Any()).Return(dummyUser, nil)
	userResponses, err := user_services.CreateUser(userRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.UID, userResponses.UID)
	assert.Equal(t, expectedUser.NIK, userResponses.NIK)
	assert.Equal(t, expectedUser.Name, userResponses.Name)
	assert.Equal(t, expectedUser.ContractAddress, userResponses.ContractAddress)
	assert.Equal(t, expectedUser.CreatedAt, userResponses.CreatedAt)
	assert.Equal(t, expectedUser.Status, userResponses.Status)
	assert.Equal(t, expectedUser.UpdatedAt, userResponses.UpdatedAt)
}

func TestGetUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo := mock_repository.NewMockRepository[entity.User](ctrl)
	user_services := services.NewUserRepository(mock_repo)

	dummyUser := entity.User{
		UID:             "1",
		NIK:             "123",
		Name:            "John Doe",
		ContractAddress: "0x123",
		CreatedAt:       clockie,
		Status:          "active",
		UpdatedAt:       clockie,
	}

	expectedUser := model.UserQuery{
		UID:             "1",
		NIK:             "123",
		Name:            "John Doe",
		ContractAddress: "0x123",
		CreatedAt:       clockie,
		Status:          "active",
		UpdatedAt:       clockie,
	}

	mock_repo.EXPECT().SelectById("1").Return(dummyUser, nil)
	userResponses, err := user_services.GetUserById("1")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.UID, userResponses.UID)
	assert.Equal(t, expectedUser.NIK, userResponses.NIK)
	assert.Equal(t, expectedUser.Name, userResponses.Name)
	assert.Equal(t, expectedUser.ContractAddress, userResponses.ContractAddress)
	assert.Equal(t, expectedUser.CreatedAt, userResponses.CreatedAt)
	assert.Equal(t, expectedUser.Status, userResponses.Status)
	assert.Equal(t, expectedUser.UpdatedAt, userResponses.UpdatedAt)
}

func TestGetUserByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo := mock_repository.NewMockRepository[entity.User](ctrl)
	user_services := services.NewUserRepository(mock_repo)

	dummyUser := entity.User{
		UID:             "1",
		NIK:             "123",
		Name:            "John Doe",
		ContractAddress: "0x123",
		CreatedAt:       clockie,
		Status:          "active",
		UpdatedAt:       clockie,
	}

	expectedUser := model.UserQuery{
		UID:             "1",
		NIK:             "123",
		Name:            "John Doe",
		ContractAddress: "0x123",
		CreatedAt:       clockie,
		Status:          "active",
		UpdatedAt:       clockie,
	}

	mock_repo.EXPECT().SelectByField("username", "johndoe").Return(dummyUser, nil)
	userResponses, err := user_services.GetUserByUsername("johndoe")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.UID, userResponses.UID)
	assert.Equal(t, expectedUser.NIK, userResponses.NIK)
	assert.Equal(t, expectedUser.Name, userResponses.Name)
	assert.Equal(t, expectedUser.ContractAddress, userResponses.ContractAddress)
	assert.Equal(t, expectedUser.CreatedAt, userResponses.CreatedAt)
	assert.Equal(t, expectedUser.Status, userResponses.Status)
	assert.Equal(t, expectedUser.UpdatedAt, userResponses.UpdatedAt)
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo := mock_repository.NewMockRepository[entity.User](ctrl)
	user_services := services.NewUserRepository(mock_repo)

	dummyUser := entity.User{
		UID:             "1",
		NIK:             "123",
		Name:            "John Doe",
		ContractAddress: "0x123",
		CreatedAt:       clockie,
		Status:          "DISABLE",
		UpdatedAt:       clockie,
	}

	updatedUser := entity.User{
		UID:             "1",
		NIK:             "123456",
		Name:            "John Smith",
		Username:        "johnsmith",
		Password:        "hashed_newpassword123",
		ContractAddress: "0x123",
		Secret:          "newsecret",
		Status:          "ACTIVE",
		UpdatedAt:       clockie,
	}

	expectedUser := model.UserQuery{
		UID:             "1",
		NIK:             "123456",
		Name:            "John Smith",
		ContractAddress: "0x123",
		Status:          "ACTIVE",
		UpdatedAt:       clockie,
	}

	mock_repo.EXPECT().SelectById("1").Return(dummyUser, nil)
	mock_repo.EXPECT().UpdateById("1", gomock.Any()).Return(updatedUser, nil)

	userInput := model.UpdateUserRequest{
		UID:             "1",
		Name:            "John Smith",
		Password:        "newpassword123",
		ContractAddress: "0x123",
		Secret:          "newsecret",
		Status:          "ACTIVE",
	}

	userResponses, err := user_services.UpdateUser("1", userInput)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser.UID, userResponses.UID)
	assert.Equal(t, expectedUser.NIK, userResponses.NIK)
	assert.Equal(t, expectedUser.Name, userResponses.Name)
	assert.Equal(t, expectedUser.ContractAddress, userResponses.ContractAddress)
	assert.Equal(t, expectedUser.Status, userResponses.Status)
	assert.Equal(t, expectedUser.UpdatedAt, userResponses.UpdatedAt)
}
