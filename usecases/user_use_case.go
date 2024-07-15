package usecases

import "github.com/pphaiaiai/orange-farm-fiber/entities"

type UserUseCase interface {
	CreateUser(user *entities.User) error
}

type UserService struct {
	UserRepo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		UserRepo: repo,
	}
}

func (s *UserService) CreateUser(user *entities.User) error {
	return s.UserRepo.CreateUser(user)
}
