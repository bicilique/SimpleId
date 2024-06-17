package testing

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/services"
	"SimpleId/test/mock_repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

var clockie_request = time.Now()

func setupEnv(ctrl *gomock.Controller) (*mock_repository.MockRepository[entity.Request], services.RequestRepository) {
	mock_repo := mock_repository.NewMockRepository[entity.Request](ctrl)
	request_services := services.NewRequestRepository(mock_repo)
	return mock_repo, request_services
}

func TestShowAllRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo, request_services := setupEnv(ctrl)

	dummyRequest := []entity.Request{
		{RequestID: "1", NIK: "123", Stakeholder: "PT. XYZ", CreatedAt: clockie_request, UpdatedAt: clockie_request},
		{RequestID: "2", NIK: "456", Stakeholder: "PT. SDF", CreatedAt: clockie_request, UpdatedAt: clockie_request},
	}

	expectedRequestResponse := []model.RequestResponse{
		{RequestID: "1", NIK: "123", Stakeholder: "PT. XYZ", CreatedAt: clockie_request, UpdatedAt: clockie_request},
		{RequestID: "2", NIK: "456", Stakeholder: "PT. SDF", CreatedAt: clockie_request, UpdatedAt: clockie_request},
	}

	mock_repo.EXPECT().SelectAll().Return(dummyRequest, nil)
	requestResponse, err := request_services.ShowAllRequest()
	assert.NoError(t, err)
	for i, expected := range expectedRequestResponse {
		assert.Equal(t, expected.RequestID, requestResponse[i].RequestID)
		assert.Equal(t, expected.NIK, requestResponse[i].NIK)
		assert.Equal(t, expected.Stakeholder, requestResponse[i].Stakeholder)
		assert.Equal(t, expected.CreatedAt, requestResponse[i].CreatedAt)
		assert.Equal(t, expected.UpdatedAt, requestResponse[i].UpdatedAt)
	}
}

func TestShowRequestById(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo, request_services := setupEnv(ctrl)

	dummyRequest := entity.Request{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	expected := model.RequestResponse{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	mock_repo.EXPECT().SelectById("1").Return(dummyRequest, nil)
	requestResponse, err := request_services.GetRequestById(model.GetSharingRequest{RequestID: "1"})
	assert.NoError(t, err)
	assert.Equal(t, expected.RequestID, requestResponse.RequestID)
	assert.Equal(t, expected.NIK, requestResponse.NIK)
	assert.Equal(t, expected.Stakeholder, requestResponse.Stakeholder)
	assert.Equal(t, expected.CreatedAt, requestResponse.CreatedAt)
	assert.Equal(t, expected.UpdatedAt, requestResponse.UpdatedAt)

}

func TestShowRequestByNIK(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo, request_services := setupEnv(ctrl)

	dummyRequest := entity.Request{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	expected := model.RequestResponse{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	mock_repo.EXPECT().SelectByField("nik", "123").Return(dummyRequest, nil)
	requestResponse, err := request_services.GetRequestByNIK(model.GetSharingRequest{NIK: "123"})
	assert.NoError(t, err)
	assert.Equal(t, expected.RequestID, requestResponse.RequestID)
	assert.Equal(t, expected.NIK, requestResponse.NIK)
	assert.Equal(t, expected.Stakeholder, requestResponse.Stakeholder)
	assert.Equal(t, expected.CreatedAt, requestResponse.CreatedAt)
	assert.Equal(t, expected.UpdatedAt, requestResponse.UpdatedAt)
}

func TestCreateNewRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo, request_services := setupEnv(ctrl)

	sharingRequest := model.SharingRequest{
		NIK:         "123",
		Stakeholder: "PT. XYZ",
	}

	dummyRequest := entity.Request{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	expected := model.RequestResponse{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	mock_repo.EXPECT().Insert(gomock.Any()).Return(dummyRequest, nil)
	requestResponse, err := request_services.CreateNewRequest(sharingRequest)
	assert.NoError(t, err)
	assert.Equal(t, expected.RequestID, requestResponse.RequestID)
	assert.Equal(t, expected.NIK, requestResponse.NIK)
	assert.Equal(t, expected.Stakeholder, requestResponse.Stakeholder)
	assert.Equal(t, expected.CreatedAt, requestResponse.CreatedAt)
	assert.Equal(t, expected.UpdatedAt, requestResponse.UpdatedAt)

}

func TestUpdateSharingRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo, request_services := setupEnv(ctrl)

	dummyRequest := entity.Request{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		Status:      enum.Waiting.String(),
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	updatedRequest := entity.Request{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		Status:      enum.Approved.String(),
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	expected := model.RequestResponse{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		Status:      enum.Approved.String(),
		CreatedAt:   clockie_request,
		UpdatedAt:   clockie_request,
	}

	mock_repo.EXPECT().SelectById("1").Return(dummyRequest, nil)
	mock_repo.EXPECT().UpdateById("1", gomock.Any()).Return(updatedRequest, nil)

	userInput := model.UpdateSharingRequest{
		RequestID:   "1",
		NIK:         "123",
		Stakeholder: "PT. XYZ",
		Status:      enum.Approved.String(),
	}

	requestResponse, err := request_services.UpdateSharingRequest(userInput)
	assert.NoError(t, err)
	assert.Equal(t, expected.RequestID, requestResponse.RequestID)
	assert.Equal(t, expected.NIK, requestResponse.NIK)
	assert.Equal(t, expected.Stakeholder, requestResponse.Stakeholder)
	assert.Equal(t, expected.Status, requestResponse.Status)
	assert.Equal(t, expected.CreatedAt, requestResponse.CreatedAt)
	assert.Equal(t, expected.UpdatedAt, requestResponse.UpdatedAt)

}

func TestShowRequestByField(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock_repo, request_services := setupEnv(ctrl)

	dummyRequest := []entity.Request{
		{RequestID: "1", NIK: "123", Stakeholder: "PT. XYZ", Status: enum.Approved.String(),
			CreatedAt: clockie_request, UpdatedAt: clockie_request},
		{RequestID: "2", NIK: "456", Stakeholder: "PT. SDF", Status: enum.Approved.String(),
			CreatedAt: clockie_request, UpdatedAt: clockie_request},
	}

	expectedRequestResponse := []model.RequestResponse{
		{RequestID: "1", NIK: "123", Stakeholder: "PT. XYZ", Status: enum.Approved.String(),
			CreatedAt: clockie_request, UpdatedAt: clockie_request},
		{RequestID: "2", NIK: "456", Stakeholder: "PT. SDF", Status: enum.Approved.String(),
			CreatedAt: clockie_request, UpdatedAt: clockie_request},
	}

	mock_repo.EXPECT().ShowByField("status", enum.Approved.String()).Return(dummyRequest, nil)

	requestResponse, err := request_services.ShowFilteredRequest(enum.Approved)
	assert.NoError(t, err)
	for i, expected := range expectedRequestResponse {
		assert.Equal(t, expected.RequestID, requestResponse[i].RequestID)
		assert.Equal(t, expected.NIK, requestResponse[i].NIK)
		assert.Equal(t, expected.Status, requestResponse[i].Status)
		assert.Equal(t, expected.Stakeholder, requestResponse[i].Stakeholder)
		assert.Equal(t, expected.CreatedAt, requestResponse[i].CreatedAt)
		assert.Equal(t, expected.UpdatedAt, requestResponse[i].UpdatedAt)
	}
}
