package model

import (
	"SimpleId/internal/model/enum"
	"time"
)

type RequestResponse struct {
	RequestID       string    `json:"requestId,omitempty"`
	NIK             string    `json:"nik,omitempty"`
	Status          string    `json:"status,omitempty"`
	Stakeholder     string    `json:"stakeholder,omitempty"`
	ContractAddress string    `json:"contractaddress"  binding:"required" validate:"max=100"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
}

type VerifyRequest struct {
	Token string `validate:"required,max=100"`
}

type SharingRequest struct {
	NIK         string `json:"nik"  binding:"required" validate:"max=100"`
	Stakeholder string `json:"stakeholder"  binding:"required" validate:"max=100"`
}

type UpdateSharingRequest struct {
	RequestID       string `json:"requestId" validate:"required,max=100"`
	NIK             string `json:"nik"  binding:"required" validate:"max=100"`
	Stakeholder     string `json:"stakeholder"  binding:"required" validate:"max=100"`
	ContractAddress string `json:"contract_address"  binding:"required" validate:"max=100"`
	Status          string `json:"status" binding:"required" validate:"max=100"`
}

type GetSharingRequest struct {
	RequestID string `json:"requestId,omitempty" validate:"max=100"`
	NIK       string `json:"nik,omitempty" validate:"max=100"`
}

type UserRequest struct {
	UID      string `json:"uid" validate:"required,max=100"`
	Username string `json:"username" binding:"required"`
}

type StatusFilterRequest struct {
	Status enum.Status `json:"status" validate:"required,max=100"`
}

type ApproveSharingRequest struct {
	RequestID       string `json:"requestId,omitempty" validate:"max=100"`
	NIK             string `json:"nik,omitempty" validate:"max=100"`
	ContractAddress string `json:"contractAddress,omitempty"  binding:"required" validate:"max=100"`
}

type ReceiveSharingRequest struct {
	Secret          string `json:"secret,omitempty" validate:"max=100"`
	ContractAddress string `json:"contractAddress,omitempty"  binding:"required" validate:"max=100"`
}
