package model

import "go-crawler/crawler/backend/pkg/errno"

type AdminModel struct {
	Id       uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	Username string `json:"username" gorm:"column:username;unique;not null" binding:"required" validate:"min=5"`
	Password string `json:"password" gorm:"column:password;not null" binding:"required" validate:"min=6,max=128"`
	RoleId   int64  `json:"role_id" gorm:"column:role_id"`
	Name	 string `json:"name" gorm:"column:name"`
}

func (c *AdminModel) TableName() string {
	return "admin"
}

// Create creates a new user account.
func (a *AdminModel) Create() error {
	return DB.Self.Create(&a).Error
}

// Update updates an user account information.
func (a *AdminModel) Update() error {
	return DB.Self.Save(a).Error
}

// GetUser gets an user by the user identifier.
func GetAdmin(username string) (*AdminModel, error) {
	a := &AdminModel{}
	d := DB.Self.Where("username = ?", username).First(&a)
	return a, d.Error
}

func GetAdminById(id uint64) (*AdminModel, error) {
	a := &AdminModel{}
	d := DB.Self.Where("id = ?", id).First(&a)
	return a, d.Error
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (a *AdminModel) Compare(pwd string) (err error) {
	if a.Password != pwd {
		return errno.ErrPasswordIncorrect
	}
	return
}
