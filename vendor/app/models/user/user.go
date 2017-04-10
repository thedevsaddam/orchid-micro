package user

import (
	"app/database"
	"github.com/jinzhu/gorm"
)

type Users struct {
	gorm.Model
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type UsersWithProfile struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Bio      string `json:"bio"`
}

func FetchUsers() []UsersWithProfile {
	db := database.Connect()
	defer db.Close()
	users_with_profiles := []UsersWithProfile{}
	db.Table("users").
		Select("users.username, profiles.name, profiles.bio").
		Joins("left join profiles on profiles.user_id = users.id").
		Scan(&users_with_profiles)
	return users_with_profiles

}
