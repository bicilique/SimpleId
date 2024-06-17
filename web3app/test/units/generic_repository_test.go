package testing

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/repository"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mockRepo = &repository.GenericRepositoryMock[entity.User]{}

func TestInsert_Success(t *testing.T) {
	// Prepare test data
	entity := entity.User{
		NIK:             "123456789",
		Name:            "Test User",
		Username:        "testuser",
		Password:        "password123",
		ContractAddress: "0x123abc",
		Secret:          "secret123",
		Status:          "active",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	mockRepo.Mock.On("Insert", entity).Return(entity, nil)
	// Call the Insert method
	result, err := mockRepo.Insert(entity)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, entity, result)
}

func TestInsert_Error(t *testing.T) {
	// Create a mock repository
	mockRepo := &repository.GenericRepositoryMock[entity.User]{} // Initialize the mock

	// Prepare test data
	preparedData := entity.User{
		NIK:             "123456789",
		Name:            "Test User",
		Username:        "testuser",
		Password:        "password123",
		ContractAddress: "0x123abc",
		Secret:          "secret123",
		Status:          "active",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Set expectations on the mock to return an error
	mockRepo.Mock.On("Insert", preparedData).Return(entity.User{}, errors.New("insert error"))

	// Call the Insert method
	result, err := mockRepo.Insert(preparedData)

	// Assert the error
	assert.Error(t, err)
	assert.Equal(t, entity.User{}, result)
}

func TestSelcetAll(t *testing.T) {
	users := []entity.User{
		{UID: "1", NIK: "123", Name: "User 1", Username: "user1", Password: "pass1"},
		{UID: "2", NIK: "456", Name: "User 2", Username: "user2", Password: "pass2"},
	}

	// Set expectations on the mock
	mockRepo.Mock.On("SelectAll").Return(users, nil)

	// Call the SelectAll method
	result, err := mockRepo.SelectAll()

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, users, result)
}

func TestSelectAll_Error(t *testing.T) {
	// Set expectations on the mock to return an error
	mockRepo.Mock.On("SelectAll").Return([]entity.User{}, errors.New("select error"))

	// Call the SelectAll method
	result, err := mockRepo.SelectAll()

	// Assert the error
	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestSelectById_Success(t *testing.T) {
	// Prepare test data
	user := entity.User{UID: "1", NIK: "123", Name: "Test User", Username: "testuser", Password: "password"}

	// Set expectations on the mock
	mockRepo.Mock.On("SelectById", "1").Return(user, nil)

	// Call the SelectById method
	result, err := mockRepo.SelectById("1")

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, user, result)
}

func TestSelectById_Error(t *testing.T) {
	mockRepo.Mock.On("SelectById", "1").Return(entity.User{}, errors.New("select error"))

	// Call the SelectById method
	result, err := mockRepo.SelectById("1")

	// Assert the error
	assert.Error(t, err)
	assert.Empty(t, result)
}
