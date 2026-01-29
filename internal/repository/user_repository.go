package repository

import (
	"backend-boking-ticket/internal/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user entity.User) error
	FindByEmail(email string) (entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user entity.User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) FindByEmail(email string) (entity.User, error) {
	var user entity.User
	// GORM akan mengisi user jika ditemukan, error jika terjadi masalah query.
	// Jika record not found, GORM v2 mengembalikan error ErrRecordNotFound.
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
