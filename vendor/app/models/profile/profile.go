package profile

import "github.com/jinzhu/gorm"

type Profiles struct {
	gorm.Model
	Name   string `json:"name,omitempty"`
	Bio    string `json:"bio,omitempty"`
	Age    int    `json:"age,omitempty"`
	UserId int    `json:"user_id,omitempty"`
}
