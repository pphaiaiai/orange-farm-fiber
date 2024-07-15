package usecases

import "github.com/pphaiaiai/orange-farm-fiber/entities"

type UserRepository interface {
	CreateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
}
