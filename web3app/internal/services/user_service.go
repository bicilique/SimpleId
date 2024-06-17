package services

import (
	"SimpleId/internal/entity"
	"SimpleId/internal/model"
	"SimpleId/internal/model/converter"
	"SimpleId/internal/model/enum"
	"SimpleId/internal/repository"
	"time"

	"SimpleId/internal/utils"

	"github.com/google/uuid"
)

type UserRepository interface {
	ShowAllUser() ([]*model.UserQuery, error)
	GetUserById(id string) (*model.UserQuery, error)
	GetUserByUsername(username string) (*model.UserQuery, error)
	CreateUser(userRequest model.RegisterUserRequest) (*model.UserResponse, error)
	UpdateUser(id string, userInput model.UpdateUserRequest) (*model.UserQuery, error)
	GetUserByNIK(nik string) (*model.UserQuery, error)
	LoginUser(userInput model.LoginUserRequest) (*model.LoginResponse, error)
}

type userRepository struct {
	repo repository.Repository[entity.User]
}

func NewUserRepository(repo repository.Repository[entity.User]) *userRepository {
	return &userRepository{repo}
}

func (u *userRepository) ShowAllUser() ([]*model.UserQuery, error) {
	users, err := u.repo.SelectAll()
	if err != nil {
		return nil, err
	}
	results := converter.UserArrayToResponses(users)
	return results, nil
}

func (u *userRepository) GetUserById(id string) (*model.UserQuery, error) {
	user, err := u.repo.SelectByUID(id)
	if err != nil {
		return nil, err
	}
	results := converter.UserToResponse(&user)
	return results, nil
}

func (u *userRepository) GetUserByUsername(username string) (*model.UserQuery, error) {
	user, err := u.repo.SelectByField("username", username)
	if err != nil {
		return nil, err
	}
	results := converter.UserToResponse(&user)
	return results, nil
}

func (u *userRepository) GetUserByNIK(nik string) (*model.UserQuery, error) {
	user, err := u.repo.SelectByField("nik", nik)
	if err != nil {
		return nil, err
	}
	results := converter.UserToResponse(&user)
	return results, nil
}

func (u *userRepository) CreateUser(userRequest model.RegisterUserRequest) (*model.UserResponse, error) {

	// checkResults, _ := u.repo.SelectByField("username", userRequest.Username)
	// if checkResults.Username != "" {
	// 	return nil, err
	// }

	user := entity.User{}
	user.UID = uuid.New().String()
	user.NIK = userRequest.NIK
	user.Name = userRequest.Name
	user.Username = userRequest.Username
	user.Status = enum.Waiting.String()
	user.Password, _ = utils.HashAndEncodePassword(userRequest.Password)
	user.RoleID = 2
	newUser, err := u.repo.Insert(user)
	if err != nil {
		return nil, err
	}

	results := converter.EntityToUserResponse(&newUser)
	return results, nil
}

func (u *userRepository) UpdateUser(id string, userInput model.UpdateUserRequest) (*model.UserQuery, error) {
	// Retrieve the existing user by ID
	existingUser, err := u.GetUserById(id)
	if err != nil {
		return existingUser, err
	}

	// Prepare updates
	updates := make(map[string]interface{})

	// if userInput.NIK != "" {
	// 	updates["nik"] = userInput.NIK
	// }

	// if userInput.Username != "" {
	// 	updates["username"] = userInput.Username
	// }

	if userInput.Name != "" {
		updates["name"] = userInput.Name
	}

	if userInput.Password != "" {
		hashedPassword, err := utils.HashAndEncodePassword(userInput.Password)
		if err != nil {
			return existingUser, err
		}
		updates["password"] = hashedPassword
	}
	if userInput.ContractAddress != "" {
		updates["contract_address"] = userInput.ContractAddress
	}
	if userInput.Secret != "" {
		updates["secret"] = userInput.Secret
	}
	if userInput.Status != "" {
		updates["status"] = userInput.Status
	}
	updates["updated_at"] = time.Now()

	updatedUser, err := u.repo.UpdateByUID(id, updates)
	if err != nil {
		return existingUser, err
	}
	// Return the updated user
	results := converter.UserToResponse(&updatedUser)
	return results, nil
}

func (u *userRepository) LoginUser(userInput model.LoginUserRequest) (*model.LoginResponse, error) {
	user, err := u.repo.SelectByField("username", userInput.Username)
	if err != nil {
		return nil, err
	}
	result, err := utils.CheckPasswordHash(userInput.Password, user.Password)
	if !result {
		return nil, err
	}

	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Username: user.Username,
		Token:    jwt,
		Message:  "Successfully logged in",
	}, nil

	// return converter.EntityToUserResponse(&user), nil
}
