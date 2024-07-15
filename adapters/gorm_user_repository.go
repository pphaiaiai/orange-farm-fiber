package adapters

import (
	"github.com/pphaiaiai/orange-farm-fiber/entities"
	"github.com/pphaiaiai/orange-farm-fiber/usecases"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) usecases.UserRepository {
	return &GormUserRepository{
		db: db,
	}
}

// CreateUser implements usecases.UserRepository.
func (g *GormUserRepository) CreateUser(user *entities.User) error {
	return g.db.Create(user).Error
}

// FindUserByEmail implements usecases.UserRepository.
func (g *GormUserRepository) FindUserByEmail(email string) (*entities.User, error) {
	return nil, nil
}
