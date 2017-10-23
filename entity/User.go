package entity

import (
	"github.com/jinzhu/gorm"
)

//User entity
type User struct {
	gorm.Model
	Age     int    `json:"age"`
	Name    string `gorm:"size:255" json:"name"`
	Address string `gorm:"type:varchar(100) " json:"address"`
}
