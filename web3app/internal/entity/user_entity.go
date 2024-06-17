package entity

import "time"

type User struct {
	UID             string    `gorm:"primaryKey" json:"uid"`
	NIK             string    `gorm:"unique" json:"nik"`
	Name            string    `json:"name"`
	Username        string    `gorm:"unique" json:"username"`
	Password        string    `json:"password"`
	ContractAddress string    `json:"contractAddress"`
	Secret          string    `json:"secret"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	RoleID          uint      `gorm:"not null;DEFAULT:3" json:"role_id"`
	Role            Role      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (u *User) TableName() string {
	return "users"
}
