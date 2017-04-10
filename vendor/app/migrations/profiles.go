package migrations

import "github.com/jinzhu/gorm"

type Profiles struct {
	gorm.Model
	Bio string `json:"bio,omitempty" gorm:"size:255"`
	Age int    `json:"age,omitempty" gorm:"size:5"`
}
