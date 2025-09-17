package userService

import (
	"gorm.io/gorm"
)

type UserRepositoryIface interface {
	Create(user User) error
	FindByEmail(emailFromBody string) (*User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryIface {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user User) error {
	return r.db.Create(&user).Error
}

func (r *UserRepository) FindByEmail(emailFromBody string) (*User, error) {
	var user User
	if err := r.db.First(&user, "email = ?", emailFromBody).Error; err != nil {
		return &User{}, err
	}
	return &user, nil
}
