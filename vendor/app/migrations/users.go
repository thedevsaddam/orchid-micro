package migrations

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Username string `json:"username" gorm:"size:40"`
	Password string `json:"password" gorm:"size:80"`
}
