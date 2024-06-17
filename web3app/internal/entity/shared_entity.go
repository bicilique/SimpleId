package entity

import "time"

type Shared struct {
	SharedID        int       `gorm:"primaryKey" json:"sharedId"`
	NIK             string    `json:"nik"`
	ContractAddress string    `json:"contractAddress"`
	SharedSecret    string    `json:"sharedSecret"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (u *Shared) TableName() string {
	return "shared"
}
