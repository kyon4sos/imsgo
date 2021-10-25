package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ChatUser struct {
	UUID     string `gorm:"type: binary(36);"`
	Username string `gorm:"type:varchar(255)"`
	Password string `gorm:"type:varchar(255)"`
}
func (u *ChatUser) BeforeCreate(tx *gorm.DB) (err error) {
	//newUUID, err := uuid.NewUUID()
	u.UUID = uuid.NewString()
	return
}