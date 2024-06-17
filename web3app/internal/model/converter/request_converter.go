package converter

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
)

func RequestToResponse(request *entity.Request) *model.RequestResponse {
	return &model.RequestResponse{
		RequestID:   request.RequestID,
		NIK:         request.NIK,
		Status:      request.Status,
		Stakeholder: request.Stakeholder,
		CreatedAt:   request.CreatedAt,
		UpdatedAt:   request.UpdatedAt,
	}
}

func RequestArrayToResponses(requests []entity.Request) []*model.RequestResponse {
	responses := make([]*model.RequestResponse, len(requests))
	for i, request := range requests {
		responses[i] = &model.RequestResponse{
			RequestID:   request.RequestID,
			NIK:         request.NIK,
			Status:      request.Status,
			Stakeholder: request.Stakeholder,
			CreatedAt:   request.CreatedAt,
			UpdatedAt:   request.UpdatedAt,
		}
	}
	return responses
}

func ApproveSharingRequestToString(input model.ApproveSharingRequest) (string, model.GetSharingRequest) {
	return input.ContractAddress, model.GetSharingRequest{
		RequestID: input.RequestID,
		NIK:       input.NIK,
	}
}
