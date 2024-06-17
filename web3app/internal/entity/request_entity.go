package entity

import "time"

type Request struct {
	RequestID       string    `gorm:"primaryKey" json:"requestId"`
	NIK             string    `json:"nik"`
	Status          string    `json:"status"`
	Stakeholder     string    `json:"stakeholder"`
	ContractAddress string    `json:"contractaddress"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (u *Request) TableName() string {
	return "requests"
}
