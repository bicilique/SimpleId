package model

import "time"

type UserQuery struct {
	UID             string    `json:"uid,omitempty"`
	Name            string    `json:"name,omitempty"`
	NIK             string    `json:"nik,omitempty"`
	ContractAddress string    `json:"contractAddress,omitempty"`
	Secret          string    `json:"secret,omitempty"`
	Status          string    `json:"status,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

type VerifyUserRequest struct {
	Token string `validate:"required,max=100"`
}

type RegisterUserRequest struct {
	NIK      string `json:"nik" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	UID             string `json:"-" validate:"required,max=100"`
	Password        string `json:"password,omitempty" validate:"max=100"`
	Name            string `json:"name,omitempty" validate:"max=100"`
	ContractAddress string `json:"contractAddress,omitempty"`
	Secret          string `json:"secret,omitempty" validate:"max=100"`
	Status          string `json:"status,omitempty" validate:"max=100"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}

type LoginResponse struct {
	Username string `json:"username" validate:"required,max=100"`
	Token    string `json:"token,omitempty"`
	Message  string `json:"message,omitempty"`
}

type UserResponse struct {
	UID             string    `json:"uid,omitempty"`
	Name            string    `json:"name,omitempty"`
	NIK             string    `json:"nik,omitempty"`
	ContractAddress string    `json:"contractAddress,omitempty"`
	Status          string    `json:"status,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

type LogoutUserRequest struct {
	UID string `json:"uid" validate:"required,max=100"`
}

type UserInformation struct {
	UID       string `json:"uid,omitempty" validate:"max=160"`
	Name      string `json:"name,omitempty" validate:"max=160"`
	NIK       string `json:"nik,omitempty" validate:"max=160"`
	Address   string `json:"address,omitempty" validate:"max=160"`
	Country   string `json:"country,omitempty" validate:"max=160"`
	Email     string `json:"email,omitempty" validate:"max=160"`
	BirthDate string `json:"birthdate,omitempty" validate:"max=160"`
	Issuer    string `json:"issuer,omitempty" validate:"max=160"`
	Status    string `json:"status,omitempty" validate:"max=160"`
	Active    string `json:"active,omitempty" validate:"max=160"`
}

type UserDTO struct {
	UID      string `json:"uid" validate:"required,max=100"`
	Username string `json:"username" binding:"required"`
}
