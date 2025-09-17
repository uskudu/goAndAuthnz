package userService

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserServiceIface interface {
	Register(email, password string) error
	Authenticate(email, password string) (*User, error)
}

type UserService struct {
	repo UserRepositoryIface
}

func NewUserService(r UserRepositoryIface) UserServiceIface {
	return &UserService{repo: r}
}

func (s *UserService) Register(email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	user := User{
		Email:    email,
		Password: string(hash),
	}
	return s.repo.Create(user)
}

func (s *UserService) Authenticate(email, password string) (*User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return nil, errors.New("invalid password")
	}
	return user, nil
}
