package services

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
	"SimpleId/internal/model/converter"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/repository"
	"time"

	"github.com/google/uuid"
)

type RequestRepository interface {
	ShowAllRequest() ([]*model.RequestResponse, error)
	GetRequestById(userInput model.GetSharingRequest) (*model.RequestResponse, error)
	GetRequestByNIK(userInput model.GetSharingRequest) (*model.RequestResponse, error)
	CreateNewRequest(userInput model.SharingRequest) (*model.RequestResponse, error)
	UpdateSharingRequest(userInput model.UpdateSharingRequest) (*model.RequestResponse, error)
	ShowFilteredRequest(status enum.Status) ([]*model.RequestResponse, error)
}

type requestRepository struct {
	repo repository.Repository[entity.Request]
}

func NewRequestRepository(repo repository.Repository[entity.Request]) *requestRepository {
	return &requestRepository{repo}
}

func (r *requestRepository) GetRequestById(userInput model.GetSharingRequest) (*model.RequestResponse, error) {
	request, err := r.repo.SelectByRequestId(userInput.RequestID)
	if err != nil {
		return nil, err
	}
	results := converter.RequestToResponse(&request)
	return results, nil
}

func (r *requestRepository) GetRequestByNIK(userInput model.GetSharingRequest) (*model.RequestResponse, error) {
	request, err := r.repo.SelectByField("nik", userInput.NIK)
	if err != nil {
		return nil, err
	}
	results := converter.RequestToResponse(&request)
	return results, nil
}

func (r *requestRepository) ShowAllRequest() ([]*model.RequestResponse, error) {
	requests, err := r.repo.SelectAll()
	if err != nil {
		return nil, err
	}
	results := converter.RequestArrayToResponses(requests)
	return results, nil
}

func (r *requestRepository) CreateNewRequest(userInput model.SharingRequest) (*model.RequestResponse, error) {
	request := entity.Request{}
	request.RequestID = uuid.New().String()
	request.NIK = userInput.NIK
	request.Stakeholder = userInput.Stakeholder
	request.Status = string(enum.Waiting)

	newRequest, err := r.repo.Insert(request)
	if err != nil {
		return nil, err
	}

	reults := converter.RequestToResponse(&newRequest)
	return reults, nil
}

func (r *requestRepository) UpdateSharingRequest(userInput model.UpdateSharingRequest) (*model.RequestResponse, error) {
	exsistingRequest, err := r.GetRequestById(model.GetSharingRequest{RequestID: userInput.RequestID})
	if err != nil {
		return exsistingRequest, err
	}

	updates := make(map[string]interface{})
	if userInput.NIK != "" {
		updates["nik"] = userInput.NIK
	}

	if userInput.Stakeholder != "" {
		updates["stakeholder"] = userInput.Stakeholder
	}

	if userInput.Status != "" {
		updates["status"] = userInput.Status
	}

	if userInput.ContractAddress != "" {
		updates["contract_address"] = userInput.ContractAddress
	}

	updates["updated_at"] = time.Now()

	updatedRequest, err := r.repo.UpdateByRequestId(userInput.RequestID, updates)
	if err != nil {
		return exsistingRequest, err
	}

	results := converter.RequestToResponse(&updatedRequest)
	return results, nil
}

func (r *requestRepository) ShowFilteredRequest(status enum.Status) ([]*model.RequestResponse, error) {
	requests, err := r.repo.ShowByField("status", status.String())
	if err != nil {
		return nil, err
	}
	results := converter.RequestArrayToResponses(requests)
	return results, nil
}
