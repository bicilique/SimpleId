package converter

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
)

func UserToResponse(user *entity.User) *model.UserQuery {
	return &model.UserQuery{
		UID:             user.UID,
		Name:            user.Name,
		NIK:             user.NIK,
		ContractAddress: user.ContractAddress,
		Secret:          user.Secret,
		Status:          user.Status,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}

func UserArrayToResponses(users []entity.User) []*model.UserQuery {
	responses := make([]*model.UserQuery, len(users))
	for i, user := range users {
		responses[i] = &model.UserQuery{
			UID:             user.UID,
			Name:            user.Name,
			NIK:             user.NIK,
			ContractAddress: user.ContractAddress,
			Secret:          user.Secret,
			Status:          user.Status,
			CreatedAt:       user.CreatedAt,
			UpdatedAt:       user.UpdatedAt,
		}
	}
	return responses
}

func EntityToUserResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		UID:             user.UID,
		Name:            user.Name,
		NIK:             user.NIK,
		ContractAddress: user.ContractAddress,
		Status:          user.Status,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
	}
}
