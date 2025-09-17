package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Signup(user User) (User, error)
	Login(user User, emailFromBody string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Signup(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *userRepository) Login(user User, emailFromBody string) error {
	return r.db.First(&user, "email = ?", emailFromBody).Error
}
